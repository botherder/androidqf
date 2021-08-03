// Snoopdroid 2
// Copyright (c) 2021 Claudio Guarnieri.
// Use of this software is governed by the MVT License 1.1 that can be found at
//   https://license.mvt.re/1.1/

package acquisition

import (
	"fmt"
	"os"
	"path/filepath"
)

func (a *Acquisition) Processes() error {
	fmt.Println("Extracting list of running processes...")

	out, err := a.ADB.Shell("ps")
	if err != nil {
		return fmt.Errorf("Unable to run `adb shell ps`: %s", err)
	}

	fileName := "ps.txt"
	file, err := os.Create(filepath.Join(a.BasePath, fileName))
	if err != nil {
		return fmt.Errorf("Unable to create %s file: %s", fileName, err)
	}
	defer file.Close()

	file.WriteString(out)

	return nil
}
