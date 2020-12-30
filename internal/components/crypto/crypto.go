/*
------------------------------------------------------------------------------------------------------------------------
####### crypto ####### (c) 2020-2021 mls-361 ####################################################### MIT License #######
------------------------------------------------------------------------------------------------------------------------
*/

package crypto

import (
	"github.com/mls-361/armen/internal/components"
	"github.com/mls-361/component"
	"github.com/mls-361/crypto"
)

type (
	// Crypto AFAIRE.
	Crypto struct {
		*component.Base
		crypto *crypto.Crypto
	}
)

// New AFAIRE.
func New(components *components.Components) *Crypto {
	crypto := crypto.New()
	components.Crypto = crypto

	return &Crypto{
		Base:   component.NewBase("crypto", "crypto"),
		crypto: crypto,
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
	cc.Built()
	return nil
}

/*
######################################################################################################## @(°_°)@ #######
*/
