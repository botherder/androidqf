package assets

import (
	_ "embed"
)

//go:embed "adb"
var adbData []byte

func getAssets() []Asset {
	return []Asset{
		{Name: "adb", Data: adbData},
	}
}
