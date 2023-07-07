// Copyright (c) 2021-2023 Claudio Guarnieri.
// Use of this source code is governed by the MVT License 1.1
// which can be found in the LICENSE file.

package adb

import (
	"fmt"
	"os/exec"
	"strings"
)

type ADB struct {
	ExePath string
}

var Client *ADB

// New returns a new ADB instance.
func New() (*ADB, error) {
	adb := ADB{}
	err := adb.findExe()
	if err != nil {
		return nil, fmt.Errorf("failed to find a usable adb executable: %v", err)
	}

	return &adb, nil
}

// GetState returns the output of `adb get-state`.
// It is used to check whether a device is connected. If it is not, adb
// will exit with status 1.
func (a *ADB) GetState() (string, error) {
	out, err := exec.Command(a.ExePath, "get-state").Output()
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(out)), nil
}

// Shell executes a shell command through adb.
func (a *ADB) Shell(cmd ...string) (string, error) {
	fullCmd := append([]string{"shell"}, cmd...)
	out, err := exec.Command(a.ExePath, fullCmd...).Output()
	if err != nil {
		if out == nil {
			return "", err
		}
		// Still return a value because some commands returns 1 but still works.
		return strings.TrimSpace(string(out)), err
	}

	return strings.TrimSpace(string(out)), nil
}

// Pull downloads a file from the device to a local path.
func (a *ADB) Pull(remotePath, localPath string) (string, error) {
	out, err := exec.Command(a.ExePath, "pull", remotePath, localPath).Output()
	if err != nil {
		return string(out), err
	}

	return string(out), nil
}

// Backup generates a backup of the specified app, or of all.
func (a *ADB) Backup(arg string) error {
	cmd := exec.Command(a.ExePath, "backup", arg)
	return cmd.Run()
}

// List files in a folder using ls, returns array of strings.
func (a *ADB) ListFiles(remotePath string) []string {
	var remoteFiles []string

	out, _ := a.Shell("find", remotePath, "-type", "f", "2>", "/dev/null")
	if out == "" {
		return []string{}
	}

	for _, file := range strings.Split(out, "\n") {
		if strings.HasPrefix(file, "find:") {
			continue
		}
		remoteFiles = append(remoteFiles, file)
	}

	return remoteFiles
}
