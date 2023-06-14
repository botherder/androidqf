//go:build linux

package adb

import (
	"os/exec"
	"path/filepath"

	"github.com/botherder/androidqf/assets"
	rt "github.com/botherder/go-savetime/runtime"
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

	a.ExePath = filepath.Join(rt.GetExecutableDirectory(), "adb")
	return nil
}
