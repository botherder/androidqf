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
	adbPath, err := exec.LookPath("adb")
	if err == nil {
		a.ExePath = adbPath
		return nil
	}

	adbPath = filepath.Join(utils.GetBinFolder(), "adb")
	adbData, err := Asset("adb")
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
