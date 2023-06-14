package modules

import (
	"fmt"
	"path/filepath"

	"github.com/botherder/androidqf/adb"
)

type Services struct {
	StoragePath string
}

func NewServices() *Services {
	return &Services{}
}

func (s *Services) Name() string {
	return "services"
}

func (s *Services) InitStorage(storagePath string) error {
	s.StoragePath = storagePath
	return nil
}

func (s *Services) Run() error {
	fmt.Println("Collecting list of services...")

	out, err := adb.Client.Shell("service list")
	if err != nil {
		return fmt.Errorf("failed to run `adb shell service list`: %v", err)
	}

	return saveCommandOutput(filepath.Join(s.StoragePath, "services.txt"), out)
}
