// androidqf - Android Quick Forensics
// Copyright (c) 2021-2022 Claudio Guarnieri.
// Use of this software is governed by the MVT License 1.1 that can be found at
//   https://license.mvt.re/1.1/

package acquisition

import (
	"fmt"
)

func (a *Acquisition) DumpSys() error {
	fmt.Println("Collecting device diagnostic information. This might take a while...")

	out, err := a.ADB.Shell("dumpsys")
	if err != nil {
		return fmt.Errorf("failed to run `adb shell dumpsys`: %v", err)
	}

	return a.saveOutput("dumpsys.txt", out)
}
