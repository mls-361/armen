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

	"github.com/mls-361/armen/internal/components"
)

const (
	_poolMaxConns     = 10
	_connectTimeout   = 5 * time.Second
	_updateInterval   = 5 * time.Second
	_lockInsertJob    = 1
	_defaultHistoryRT = 7
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
		historyRT  time.Duration
		cluster    *hapgsql.Cluster
	}
)

func New(components *components.Components) *Backend {
	return &Backend{
		components: components,
	}
}

func (cb *Backend) newClient(logger logger.Logger, cfg *config) (*pgsql.Client, error) {
	password, err := cb.components.CCrypto.DecryptString(cfg.Password)
	if err != nil {
		return nil, err
	}

	uri := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?pool_max_conns=%d",
		cfg.Username,
		password,
		cfg.Host,
		cfg.Port,
		cfg.Database,
		_poolMaxConns,
	)

	client := pgsql.NewClient(logger)

	ctx, cancel := client.ContextWT(_connectTimeout)
	defer cancel()

	return client, client.Connect(ctx, uri)
}

// Build AFAIRE.
func (cb *Backend) Build() error {
	logger := cb.components.CLogger
	cconfig := cb.components.CConfig

	hrt, err := cconfig.Data().IntWD(_defaultHistoryRT, "components", "backend", "history", "retention_time")
	if err != nil {
		return err
	}

	logger.Info("History retention time", "hours", hrt) //::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::

	cb.historyRT = time.Duration(-1*hrt) * time.Hour

	cluster := hapgsql.NewCluster(
		hapgsql.WithLogger(logger),
		hapgsql.WithUpdateInterval(_updateInterval),
	)

	var cfg []*config

	if err := cconfig.Decode(&cfg, true, "components", "backend", "pgsql"); err != nil {
		return err
	}

	for _, c := range cfg {
		client, err := cb.newClient(logger, c)
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

func (cb *Backend) primaryPreferred() (*pgsql.Client, error) {
	node := cb.cluster.PrimaryPreferred()
	if node == nil {
		return nil, failure.New(nil).Msg("there is no alive node") /////////////////////////////////////////////////////
	}

	return node.Client(), nil
}

func (cb *Backend) advisoryLock(t *pgsql.Transaction, id int) error {
	_, err := t.Execute("SELECT pg_advisory_xact_lock($1)", id)
	return err
}

// Clean AFAIRE.
func (cb *Backend) Clean() (int, int, error) {
	client, err := cb.primary()
	if err != nil {
		return 0, 0, err
	}

	ctx, cancel := client.ContextWT(10 * time.Second)
	defer cancel()

	cj, err := cb.deleteFinishedJobs(ctx, client)
	if err != nil {
		return 0, 0, err
	}

	cw, err := cb.deleteFinishedWorkflows(ctx, client)
	if err != nil {
		return 0, 0, err
	}

	return int(cj), int(cw), cb.deleteOldestHistory(ctx, client)
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
