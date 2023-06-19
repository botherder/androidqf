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
