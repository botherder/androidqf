// androidqf - Android Quick Forensics
// Copyright (c) 2021-2022 Claudio Guarnieri.
// Use of this software is governed by the MVT License 1.1 that can be found at
//   https://license.mvt.re/1.1/

package acquisition

import (
	"fmt"
)

func (a *Acquisition) Services() error {
	fmt.Println("Collecting list of services...")

	out, err := a.ADB.Shell("service list")
	if err != nil {
		return fmt.Errorf("failed to run `adb shell service list`: %v", err)
	}

	return a.saveOutput("services.txt", out)
}
