// androidqf - Android Quick Forensics
// Copyright (c) 2021 Claudio Guarnieri.
// Use of this software is governed by the MVT License 1.1 that can be found at
//   https://license.mvt.re/1.1/

package acquisition

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/botherder/androidqf/adb"
	"github.com/satori/go.uuid"
)

type Acquisition struct {
	UUID     string
	ADB      *adb.ADB
	BasePath string
	APKSPath string
	Datetime time.Time
}

// New returns a new Acquisition instance.
func New() (*Acquisition, error) {
	acq := Acquisition{}
	uuidBytes := uuid.NewV4()
	acq.UUID = uuidBytes.String()
	acq.Datetime = time.Now().UTC()

	err := acq.initADB()
	if err != nil {
		return nil, fmt.Errorf("Failed to initialize adb: %s", err)
	}

	err = acq.createFolder()
	if err != nil {
		return nil, fmt.Errorf("Failed to create acquisition folder: %s", err)
	}

	return &acq, nil
}

func (a *Acquisition) initADB() error {
	var err error
	a.ADB, err = adb.New()
	if err != nil {
		return fmt.Errorf("Failed to initialize adb: %s", err)
	}

	_, err = a.ADB.GetState()
	if err != nil {
		return fmt.Errorf("Unable to get adb state, are you sure a device is connected?")
	}

	return nil
}

func (a *Acquisition) createFolder() error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	a.BasePath = filepath.Join(cwd, a.UUID)
	err = os.Mkdir(a.BasePath, 0755)
	if err != nil {
		return err
	}

	a.APKSPath = filepath.Join(a.BasePath, "apks")
	err = os.Mkdir(a.APKSPath, 0755)
	if err != nil {
		return err
	}

	return nil
}
