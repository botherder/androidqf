// Copyright (c) 2021-2023 Claudio Guarnieri.
// Use of this source code is governed by the MVT License 1.1
// which can be found in the LICENSE file.

package modules

import (
	"fmt"
	"os"
)

type Module interface {
	Name() string
	InitStorage(storagePath string) error
	Run() error
}

func List() []Module {
	return []Module{
		NewGetProp(),
		NewLogcat(),
		NewProcesses(),
		NewServices(),
		NewSettings(),
		NewDumpsys(),
		NewBackup(),
		NewLogs(),
		NewPackages(),
	}
}

func saveCommandOutput(filePath, output string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create %s file: %v", filePath, err)
	}
	defer file.Close()

	_, err = file.WriteString(output)
	if err != nil {
		return fmt.Errorf("failed to write command output to %s: %v", filePath, err)
	}

	file.Sync()

	return nil
}
