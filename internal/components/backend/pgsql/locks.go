/*
------------------------------------------------------------------------------------------------------------------------
####### pgsql ####### (c) 2020-2021 mls-361 ######################################################## MIT License #######
------------------------------------------------------------------------------------------------------------------------
*/

package pgsql

import (
	"time"

	"github.com/mls-361/pgsql"
)

// AcquireLock AFAIRE.
func (cb *Backend) AcquireLock(name, owner string, duration time.Duration) (bool, error) {
	client, err := cb.primary()
	if err != nil {
		return false, err
	}

	ctx, cancel := pgsql.Context(5 * time.Second)
	defer cancel()

	now := time.Now()
	expiry := now.Add(duration)

	count, err := client.Execute(
		ctx,
		`
		UPDATE locks SET owner = $1, expiry = $2
		WHERE name = $3 AND (owner = $4 OR expiry IS NULL OR expiry <= $5)`,
		owner,
		expiry,
		name,
		owner,
		now,
	)
	if err != nil {
		return false, err
	}

	return count == 1, nil
}

// ReleaseLock AFAIRE.
func (cb *Backend) ReleaseLock(name, owner string) error {
	client, err := cb.primary()
	if err != nil {
		return err
	}

	ctx, cancel := pgsql.Context(5 * time.Second)
	defer cancel()

	_, err = client.Execute(
		ctx,
		"UPDATE locks SET owner = NULL, expiry = NULL WHERE name = $1 AND owner = $2",
		name,
		owner,
	)

	return err
}

/*
######################################################################################################## @(°_°)@ #######
*/
