/*
------------------------------------------------------------------------------------------------------------------------
####### util ####### (c) 2020-2021 mls-361 ######################################################### MIT License #######
------------------------------------------------------------------------------------------------------------------------
*/

package util

import (
	"os"
	"path/filepath"

	"github.com/mitchellh/mapstructure"
)

// BinaryDir AFAIRE.
func BinaryDir() (string, error) {
	path, err := os.Executable()
	if err != nil {
		return "", err
	}

	return filepath.Dir(path), nil
}

// DecodeData AFAIRE.
func DecodeData(input, output interface{}) error {
	decoderConfig := &mapstructure.DecoderConfig{
		DecodeHook:       mapstructure.StringToTimeDurationHookFunc(),
		Result:           output,
		WeaklyTypedInput: true,
	}

	decoder, err := mapstructure.NewDecoder(decoderConfig)
	if err != nil {
		return err
	}

	return decoder.Decode(input)
}

/*
######################################################################################################## @(°_°)@ #######
*/
