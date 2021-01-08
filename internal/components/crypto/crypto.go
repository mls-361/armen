/*
------------------------------------------------------------------------------------------------------------------------
####### crypto ####### (c) 2020-2021 mls-361 ####################################################### MIT License #######
------------------------------------------------------------------------------------------------------------------------
*/

package crypto

import (
	"io/ioutil"

	"github.com/mls-361/armen-sdk/components"
	"github.com/mls-361/crypto"
	"github.com/mls-361/minikit"
)

type (
	// Crypto AFAIRE.
	Crypto struct {
		*minikit.Base
		components *components.Components
		crypto     *crypto.Crypto
	}
)

// New AFAIRE.
func New(components *components.Components) *Crypto {
	cc := crypto.New()
	components.Crypto = cc

	return &Crypto{
		Base:       minikit.NewBase("crypto", "crypto"),
		components: components,
		crypto:     cc,
	}
}

// Dependencies AFAIRE.
func (cc *Crypto) Dependencies() []string {
	return []string{
		"application",
	}
}

// Build AFAIRE.
func (cc *Crypto) Build(_ *minikit.Manager) error {
	keyFile, ok := cc.components.Application.LookupEnv("KEY_FILE")
	if !ok {
		return nil
	}

	key, err := ioutil.ReadFile(keyFile)
	if err != nil {
		return err
	}

	return cc.crypto.SetKey(string(key))
}

/*
######################################################################################################## @(°_°)@ #######
*/
