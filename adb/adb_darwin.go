// androidqf - Android Quick Forensics
// Copyright (c) 2021-2022 Claudio Guarnieri.
// Use of this software is governed by the MVT License 1.1 that can be found at
//   https://license.mvt.re/1.1/

package adb

import (
	"os/exec"
	"path/filepath"

	"github.com/botherder/androidqf/assets"
	"github.com/botherder/androidqf/utils"
)

func (a *ADB) findExe() error {
	adbPath, err := exec.LookPath("adb")
	if err == nil {
		a.ExePath = adbPath
		return nil
	}

	err = assets.DeployAssets()
	if err != nil {
		return err
	}

	a.ExePath = filepath.Join(utils.GetBinFolder(), "adb")
	return nil
}
