/*
------------------------------------------------------------------------------------------------------------------------
####### pgsql ####### (c) 2020-2021 mls-361 ######################################################## MIT License #######
------------------------------------------------------------------------------------------------------------------------
*/

package pgsql

import "time"

// AcquireLock AFAIRE.
func (cb *Backend) AcquireLock(name, owner string, duration time.Duration) (bool, error) {
	return false, nil
}

// ReleaseLock AFAIRE.
func (cb *Backend) ReleaseLock(name, owner string) error {
	return nil
}

/*
######################################################################################################## @(°_°)@ #######
*/
