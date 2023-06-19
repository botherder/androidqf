// Copyright (c) 2021-2023 Claudio Guarnieri.
// Use of this source code is governed by the MVT License 1.1
// which can be found in the LICENSE file.

package modules

import (
	"fmt"
	"path/filepath"

	"github.com/botherder/androidqf/adb"
)

type Dumpsys struct {
	StoragePath string
}

func NewDumpsys() *Dumpsys {
	return &Dumpsys{}
}

func (d *Dumpsys) Name() string {
	return "dumpsys"
}

func (d *Dumpsys) InitStorage(storagePath string) error {
	d.StoragePath = storagePath
	return nil
}

func (d *Dumpsys) Run() error {
	fmt.Println("Collecting device diagnostic information. This might take a while...")

	out, err := adb.Client.Shell("dumpsys")
	if err != nil {
		return fmt.Errorf("failed to run `adb shell dumpsys`: %v", err)
	}

	return saveCommandOutput(filepath.Join(d.StoragePath, "dumpsys.txt"), out)
}
