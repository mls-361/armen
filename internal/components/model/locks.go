/*
------------------------------------------------------------------------------------------------------------------------
####### model ####### (c) 2020-2021 mls-361 ######################################################## MIT License #######
------------------------------------------------------------------------------------------------------------------------
*/

package model

import "time"

// AcquireLock AFAIRE.
func (cm *Model) AcquireLock(name, owner string, duration time.Duration) (bool, error) {
	return cm.components.Backend.AcquireLock(name, owner, duration)
}

// ReleaseLock AFAIRE.
func (cm *Model) ReleaseLock(name, owner string) error {
	return cm.components.Backend.ReleaseLock(name, owner)
}

/*
######################################################################################################## @(°_°)@ #######
*/
