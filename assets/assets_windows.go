// Copyright (c) 2021-2023 Claudio Guarnieri.
// Use of this source code is governed by the MVT License 1.1
// which can be found in the LICENSE file.

package assets

import (
	_ "embed"
)

//go:embed "adb.exe"
var adbData []byte

//go:embed "AdbWinApi.dll"
var adbWinApiData []byte

//go:embed "AdbWinUsbApi.dll"
var adbWinUsbApiData []byte

func getAssets() []Asset {
	return []Asset{
		{Name: "adb.exe", Data: adbData},
		{Name: "AdbWinApi.dll", Data: adbWinApiData},
		{Name: "AdbWinUsbApi.dll", Data: adbWinUsbApiData},
	}
}
