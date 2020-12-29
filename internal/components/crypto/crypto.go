/*
------------------------------------------------------------------------------------------------------------------------
####### crypto ####### (c) 2020-2021 mls-361 ####################################################### MIT License #######
------------------------------------------------------------------------------------------------------------------------
*/

package crypto

import (
	"github.com/mls-361/component"
)

type (
	// Crypto AFAIRE.
	Crypto struct {
		*component.Base
	}
)

// New AFAIRE.
func New() *Crypto {
	return &Crypto{
		Base: component.NewBase("crypto", "crypto"),
	}
}

// Dependencies AFAIRE.
func (cc *Crypto) Dependencies() []string {
	return []string{
		"logger",
	}
}

// Build AFAIRE.
func (cc *Crypto) Build(_ *component.Manager) error {
	return nil
}

/*
######################################################################################################## @(°_°)@ #######
*/
