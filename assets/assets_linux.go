// Copyright (c) 2021-2023 Claudio Guarnieri.
// Use of this source code is governed by the MVT License 1.1
// which can be found in the LICENSE file.

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
