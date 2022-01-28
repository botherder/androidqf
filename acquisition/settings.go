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
	"strings"
)

func (a *Acquisition) Settings() error {
	fmt.Println("Extracting device settings...")

	results := map[string]map[string]string{}
	namespaces := []string{"system", "secure", "global"}

	for _, namespace := range namespaces {
		results[namespace] = map[string]string{}
		out, err := a.ADB.Shell(fmt.Sprintf("cmd settings list %s", namespace))
		if err != nil {
			return fmt.Errorf("failed to run `cmd settings %s`: %v",
				namespace, err)
		}

		for _, line := range strings.Split(out, "\n") {
			line = strings.TrimSpace(line)
			if line == "" {
				continue
			}

			fields := strings.SplitN(line, "=", 2)
			if len(fields) != 2 {
				continue
			}
			results[namespace][fields[0]] = fields[1]
		}
	}

	settingsJSONPath := filepath.Join(a.StoragePath, "settings.json")
	settingsJSON, err := os.Create(settingsJSONPath)
	if err != nil {
		return fmt.Errorf("failed to save settings to file: %v",
			err)
	}
	defer settingsJSON.Close()

	buf, _ := json.MarshalIndent(results, "", "    ")

	settingsJSON.WriteString(string(buf[:]))
	settingsJSON.Sync()

	return nil
}
