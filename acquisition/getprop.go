// androidqf - Android Quick Forensics
// Copyright (c) 2021 Claudio Guarnieri.
// Use of this software is governed by the MVT License 1.1 that can be found at
//   https://license.mvt.re/1.1/

package acquisition

import (
	"fmt"
	"os"
	"path/filepath"
)

func (a *Acquisition) GetProp() error {
	fmt.Println("Extracting device properties...")

	out, err := a.ADB.Shell("getprop")
	if err != nil {
		return fmt.Errorf("failed to run `adb shell getprop`: %v", err)
	}

	file, err := os.Create(filepath.Join(a.BasePath, "getprop.txt"))
	if err != nil {
		return fmt.Errorf("failed to create getprop.txt file: %v", err)
	}
	defer file.Close()

	file.WriteString(out)

	return nil
}
