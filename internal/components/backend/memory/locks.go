/*
------------------------------------------------------------------------------------------------------------------------
####### memory ####### (c) 2020-2021 mls-361 ####################################################### MIT License #######
------------------------------------------------------------------------------------------------------------------------
*/

package memory

import "time"

// AcquireLock AFAIRE.
func (cb *Backend) AcquireLock(name, owner string, duration time.Duration) (bool, error) {
	cb.mutex.Lock()
	defer cb.mutex.Unlock()

	l, ok := cb.locks[name]
	if !ok {
		cb.locks[name] = &lock{
			owner:  owner,
			expiry: time.Now().Add(duration),
		}

		return true, nil
	}

	if l.owner == owner {
		l.expiry = time.Now().Add(duration)
		return true, nil
	}

	if l.expiry.Before(time.Now()) {
		l.owner = owner
		l.expiry = time.Now().Add(duration)

		return true, nil
	}

	return false, nil
}

// ReleaseLock AFAIRE.
func (cb *Backend) ReleaseLock(name, owner string) error {
	cb.mutex.Lock()
	defer cb.mutex.Unlock()

	l, ok := cb.locks[name]
	if !ok {
		return nil
	}

	if l.owner == owner {
		l.owner = ""
	}

	return nil
}

/*
######################################################################################################## @(°_°)@ #######
*/
