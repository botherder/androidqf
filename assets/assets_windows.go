// androidqf - Android Quick Forensics
// Copyright (c) 2021-2022 Claudio Guarnieri.
// Use of this software is governed by the MVT License 1.1 that can be found at
//   https://license.mvt.re/1.1/

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
