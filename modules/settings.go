package modules

import (
	"fmt"
	"path/filepath"

	"github.com/botherder/androidqf/adb"
)

type Settings struct {
	StoragePath string
}

func NewSettings() *Settings {
	return &Settings{}
}

func (s *Settings) Name() string {
	return "logcat"
}

func (s *Settings) InitStorage(storagePath string) error {
	s.StoragePath = storagePath
	return nil
}

func (s *Settings) Run() error {
	fmt.Println("Collecting device settings...")

	for _, namespace := range []string{"system", "secure", "global"} {
		out, err := adb.Client.Shell(fmt.Sprintf("cmd settings list %s", namespace))
		if err != nil {
			return fmt.Errorf("failed to run `cmd settings %s`: %v", namespace, err)
		}

		err = saveCommandOutput(
			filepath.Join(s.StoragePath, fmt.Sprintf("settings_%s.txt", namespace)),
			out,
		)
		if err != nil {
			fmt.Println(err)
		}
	}

	return nil
}
