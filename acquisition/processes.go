// androidqf - Android Quick Forensics
// Copyright (c) 2021-2022 Claudio Guarnieri.
// Use of this software is governed by the MVT License 1.1 that can be found at
//   https://license.mvt.re/1.1/

package acquisition

import (
	"fmt"
)

func (a *Acquisition) Processes() error {
	fmt.Println("Collecting list of running processes...")

	out, err := a.ADB.Shell("ps -A")
	if err != nil {
		return fmt.Errorf("failed to run `adb shell ps -A`: %v", err)
	}

	return a.saveOutput("ps.txt", out)
}
