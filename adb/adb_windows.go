// androidqf - Android Quick Forensics
// Copyright (c) 2021-2022 Claudio Guarnieri.
// Use of this software is governed by the MVT License 1.1 that can be found at
//   https://license.mvt.re/1.1/

package adb

import (
	"io/ioutil"
	"os/exec"
	"path/filepath"

	"github.com/botherder/androidqf/utils"
)

func (a *ADB) findExe() error {
	adbPath, err := exec.LookPath("adb.exe")
	if err == nil {
		a.ExePath = adbPath
		return nil
	}

	cwd := utils.GetBinFolder()

	dll1Path := filepath.Join(cwd, "AdbWinApi.dll")
	dll1Data, err := Asset("AdbWinApi.dll")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(dll1Path, dll1Data, 0755)
	if err != nil {
		return err
	}

	dll2Path := filepath.Join(cwd, "AdbWinUsbApi.dll")
	dll2Data, err := Asset("AdbWinUsbApi.dll")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(dll2Path, dll2Data, 0755)
	if err != nil {
		return err
	}

	adbPath = filepath.Join(cwd, "adb.exe")
	adbData, err := Asset("adb.exe")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(adbPath, adbData, 0755)
	if err != nil {
		return err
	}

	a.ExePath = adbPath
	return nil
}
