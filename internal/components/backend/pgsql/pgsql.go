/*
------------------------------------------------------------------------------------------------------------------------
####### pgsql ####### (c) 2020-2021 mls-361 ######################################################## MIT License #######
------------------------------------------------------------------------------------------------------------------------
*/

package pgsql

import (
	"fmt"
	"time"

	"github.com/mls-361/failure"
	"github.com/mls-361/hapgsql"
	"github.com/mls-361/logger"
	"github.com/mls-361/pgsql"
	"github.com/mls-361/uuid"

	"github.com/mls-361/armen/internal/components"
)

const (
	_poolMaxConns   = 10
	_connectTimeout = 5 * time.Second
	_updateInterval = 5 * time.Second
	_lockInsertJob  = 1
)

type (
	config struct {
		Host     string
		Port     int
		Username string
		Password string
		Database string
	}

	// Backend AFAIRE.
	Backend struct {
		components *components.Components
		logger     *logger.Logger
		cluster    *hapgsql.Cluster
	}
)

func New(components *components.Components) *Backend {
	return &Backend{
		components: components,
	}
}

func (cb *Backend) newClient(c *config) (*pgsql.Client, error) {
	password, err := cb.components.CCrypto.DecryptString(c.Password)
	if err != nil {
		return nil, err
	}

	uri := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?pool_max_conns=%d",
		c.Username,
		password,
		c.Host,
		c.Port,
		c.Database,
		_poolMaxConns,
	)

	client := pgsql.NewClient(cb.logger)

	ctx, cancel := client.ContextWT(_connectTimeout)
	defer cancel()

	return client, client.Connect(ctx, uri)
}

// Build AFAIRE.
func (cb *Backend) Build() error {
	logger := cb.components.CLogger.CreateLogger(uuid.New(), "backend")

	cb.logger = logger

	cluster := hapgsql.NewCluster(
		hapgsql.WithLogger(logger),
		hapgsql.WithUpdateInterval(_updateInterval),
	)

	var cfg []*config

	if err := cb.components.CConfig.Decode(&cfg, true, "components", "backend", "pgsql"); err != nil {
		return err
	}

	for _, c := range cfg {
		client, err := cb.newClient(c)
		if err != nil {
			cluster.Close()
			return err
		}

		cluster.AddNode(hapgsql.NewNode(c.Host, client))
	}

	cb.cluster = cluster

	cluster.Update()

	return nil
}

func (cb *Backend) primary() (*pgsql.Client, error) {
	node := cb.cluster.Primary()
	if node == nil {
		return nil, failure.New(nil).Msg("there is no primary node") ///////////////////////////////////////////////////
	}

	return node.Client(), nil
}

/*
func (cb *Backend) primaryPreferred() (*pgsql.Client, error) {
	node := cb.cluster.PrimaryPreferred()
	if node == nil {
		return nil, failure.New(nil).Msg("there is no alive node") /////////////////////////////////////////////////////
	}

	return node.Client(), nil
}
*/

func (cb *Backend) advisoryLock(t *pgsql.Transaction, id int) error {
	_, err := t.Execute("SELECT pg_advisory_xact_lock($1)", id)
	return err
}

// Close AFAIRE.
func (cb *Backend) Close() {
	if cb.cluster != nil {
		cb.cluster.Close()
	}
}

/*
######################################################################################################## @(°_°)@ #######
*/
