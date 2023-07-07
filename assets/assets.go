// Copyright (c) 2021-2023 Claudio Guarnieri.
// Use of this source code is governed by the MVT License 1.1
// which can be found in the LICENSE file.

package assets

import (
	"io/ioutil"
	"path/filepath"

	rt "github.com/botherder/go-savetime/runtime"
)

type Asset struct {
	Name string
	Data []byte
}

// DeployAssets is used to retrieve the embedded adb binaries and store them.
func DeployAssets() error {
	cwd := rt.GetExecutableDirectory()

	for _, asset := range getAssets() {
		assetPath := filepath.Join(cwd, asset.Name)
		err := ioutil.WriteFile(assetPath, asset.Data, 0o755)
		if err != nil {
			return err
		}
	}

	return nil
}
