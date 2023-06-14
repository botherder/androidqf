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
		err := ioutil.WriteFile(assetPath, asset.Data, 0755)
		if err != nil {
			return err
		}
	}

	return nil
}
