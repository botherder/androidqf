// androidqf - Android Quick Forensics
// Copyright (c) 2021-2022 Claudio Guarnieri.
// Use of this software is governed by the MVT License 1.1 that can be found at
//   https://license.mvt.re/1.1/

package assets

import (
	"io/ioutil"
	"path/filepath"

	"github.com/botherder/androidqf/utils"
)

type Asset struct {
	Name string
	Data []byte
}

// DeployAssets is used to retrieve the embedded adb binaries and store them.
func DeployAssets() error {
	cwd := utils.GetBinFolder()

	for _, asset := range getAssets() {
		assetPath := filepath.Join(cwd, asset.Name)
		err := ioutil.WriteFile(assetPath, asset.Data, 0755)
		if err != nil {
			return err
		}
	}

	return nil
}
