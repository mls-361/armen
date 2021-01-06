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
	crypto := crypto.New()
	components.Crypto = crypto

	return &Crypto{
		Base:       minikit.NewBase("crypto", "crypto"),
		components: components,
		crypto:     crypto,
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

	if err := cc.crypto.SetKey(string(key)); err != nil {
		return err
	}

	cc.Built()

	return nil
}

/*
######################################################################################################## @(°_°)@ #######
*/
