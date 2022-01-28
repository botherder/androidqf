// androidqf - Android Quick Forensics
// Copyright (c) 2021-2022 Claudio Guarnieri.
// Use of this software is governed by the MVT License 1.1 that can be found at
//   https://license.mvt.re/1.1/

package acquisition

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func (a *Acquisition) GetProp() error {
	fmt.Println("Extracting device properties...")

	results := map[string]string{}

	out, err := a.ADB.Shell("getprop")
	if err != nil {
		return fmt.Errorf("failed to run `adb shellgetprop`: %v", err)
	}

	re := regexp.MustCompile(`\[(.+?)\]: \[(.+?)\]`)

	for _, line := range strings.Split(out, "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		matches := re.FindStringSubmatch(line)
		if len(matches) != 3 {
			continue
		}

		results[matches[1]] = matches[2]
	}

	getpropJSONPath := filepath.Join(a.StoragePath, "getprop.json")
	getpropJSON, err := os.Create(getpropJSONPath)
	if err != nil {
		return fmt.Errorf("failed to save getprop to file: %v",
			err)
	}
	defer getpropJSON.Close()

	buf, _ := json.MarshalIndent(results, "", "    ")

	getpropJSON.WriteString(string(buf[:]))
	getpropJSON.Sync()

	return nil
}
