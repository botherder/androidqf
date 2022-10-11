// androidqf - Android Quick Forensics
// Copyright (c) 2021-2022 Claudio Guarnieri.
// Use of this software is governed by the MVT License 1.1 that can be found at
//   https://license.mvt.re/1.1/

package acquisition

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/botherder/go-savetime/text"
)

func (a *Acquisition) Logs() error {
	fmt.Println("Collecting system logs...")

	logFiles := []string{
		"/data/system/uiderrors.txt",
		"/proc/kmsg",
		"/proc/last_kmsg",
		"/sys/fs/pstore/console-ramoops",
	}

	for _, logFolder := range []string{"/data/anr/", "/data/log/"} {
		files := a.ADB.ListFiles(logFolder)
		if len(files) == 0 {
			continue
		}

		logFiles = append(logFiles, files...)
	}

	for _, logFile := range logFiles {
		localPath := filepath.Join(a.LogsPath, logFile)
		localDir, _ := filepath.Split(localPath)

		err := os.MkdirAll(localDir, 0755)
		if err != nil {
			fmt.Printf("Failed to create folders for logs %s: %v", localDir, err)
			continue
		}

		out, err := a.ADB.Pull(logFile, localPath)
		if err != nil {
			if !text.ContainsNoCase(out, "Permission denied") {
				fmt.Printf("Failed to pull log file %s: %s\n", logFile, strings.TrimSpace(out))
			}
			continue
		}
	}

	return nil
}
