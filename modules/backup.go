package modules

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/botherder/androidqf/adb"
	"github.com/manifoldco/promptui"
)

const (
	backupOnlySMS    = "Only SMS"
	backupEverything = "Everything"
	backupNothing    = "No backup"
)

type Backup struct {
	StoragePath string
}

func NewBackup() *Backup {
	return &Backup{}
}

func (b *Backup) Name() string {
	return "logcat"
}

func (b *Backup) InitStorage(storagePath string) error {
	b.StoragePath = storagePath
	return nil
}

func (b *Backup) Run() error {
	fmt.Println("Would you like to take a backup of the device?")
	promptBackup := promptui.Select{
		Label: "Backup",
		Items: []string{backupOnlySMS, backupEverything, backupNothing},
	}
	_, backupOption, err := promptBackup.Run()
	if err != nil {
		return fmt.Errorf("failed to make selection for backup option: %v", err)
	}

	var arg string
	switch backupOption {
	case backupOnlySMS:
		arg = "com.android.providers.telephony"
	case backupEverything:
		arg = "-all"
	case backupNothing:
		return nil
	}

	fmt.Printf(
		"Generating a backup with argument %s. Please check the device to authorize the backup...\n",
		arg,
	)

	err = adb.Client.Backup(arg)
	if err != nil {
		return err
	}

	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	origBackupPath := filepath.Join(cwd, "backup.ab")
	backupPath := filepath.Join(b.StoragePath, "backup.ab")

	err = os.Rename(origBackupPath, backupPath)
	if err != nil {
		return err
	}

	fmt.Println("Backup completed!")

	return nil
}
