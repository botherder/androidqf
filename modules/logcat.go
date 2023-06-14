package modules

import (
	"fmt"
	"path/filepath"

	"github.com/botherder/androidqf/adb"
)

type Logcat struct {
	StoragePath string
}

func NewLogcat() *Logcat {
	return &Logcat{}
}

func (l *Logcat) Name() string {
	return "logcat"
}

func (l *Logcat) InitStorage(storagePath string) error {
	l.StoragePath = storagePath
	return nil
}

func (l *Logcat) Run() error {
	fmt.Println("Collecting logcat...")

	out, err := adb.Client.Shell("logcat", "-d", "-b", "all", "\"*:V\"")
	if err != nil {
		return fmt.Errorf("Failed to run `adb shell logcat`: %v\n", err)
	}

	return saveCommandOutput(filepath.Join(l.StoragePath, "logcat.txt"), out)
}
