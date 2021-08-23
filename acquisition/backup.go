// androidqf - Android Quick Forensics
// Copyright (c) 2021 Claudio Guarnieri.
// Use of this software is governed by the MVT License 1.1 that can be found at
//   https://license.mvt.re/1.1/

package acquisition

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/i582/cfmt"
	"github.com/manifoldco/promptui"
)

const (
	backupOnlySMS    = "Only SMS"
	backupEverything = "Everything"
	backupNothing    = "No backup"
)

func (a *Acquisition) Backup() error {
	fmt.Println("Would you like to take a backup of the device?")
	promptBackup := promptui.Select{
		Label: "Backup",
		Items: []string{backupOnlySMS, backupEverything, backupNothing},
	}
	_, backupOption, err := promptBackup.Run()
	if err != nil {
		return fmt.Errorf("failed to make selection for backup option: %v",
			err)
	}

	var arg string
	switch backupOption {
	case backupOnlySMS:
		arg = "com.android.providers.telephony"
		break
	case backupEverything:
		arg = "-all"
		break
	case backupNothing:
		return nil
	}

	cfmt.Printf("Generating a backup with argument {{%s}}::bold. Please check the device to authorize the backup...\n",
		arg)

	err = a.ADB.Backup(arg)
	if err != nil {
		return err
	}

	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	curBackupPath := filepath.Join(cwd, "backup.ab")
	backupPath := filepath.Join(a.StoragePath, "backup.ab")

	err = os.Rename(curBackupPath, backupPath)
	if err != nil {
		return err
	}

	cfmt.Printf("Backup completed and stored at {{%s}}::magenta|underline\n",
		backupPath)

	return nil
}
