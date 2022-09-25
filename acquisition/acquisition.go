// androidqf - Android Quick Forensics
// Copyright (c) 2021-2022 Claudio Guarnieri.
// Use of this software is governed by the MVT License 1.1 that can be found at
//   https://license.mvt.re/1.1/

package acquisition

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"github.com/botherder/androidqf/adb"
	"github.com/botherder/androidqf/utils"
	"github.com/satori/go.uuid"
)

type Acquisition struct {
	UUID        string    `json:"uuid"`
	ADB         *adb.ADB  `json:"-"`
	StoragePath string    `json:"storage_path"`
	APKSPath    string    `json:"apks_path"`
	Started     time.Time `json:"started"`
	Completed   time.Time `json:"completed"`
}

// New returns a new Acquisition instance.
func New() (*Acquisition, error) {
	acq := Acquisition{}
	uuidBytes := uuid.NewV4()
	acq.UUID = uuidBytes.String()
	acq.Started = time.Now().UTC()

	err := acq.initADB()
	if err != nil {
		return nil, fmt.Errorf("failed to initialize adb: %v", err)
	}

	err = acq.createFolder()
	if err != nil {
		return nil, fmt.Errorf("failed to create acquisition folder: %v", err)
	}

	return &acq, nil
}

func (a *Acquisition) Complete() {
	a.Completed = time.Now().UTC()
}

func (a *Acquisition) initADB() error {
	var err error
	a.ADB, err = adb.New()
	if err != nil {
		return fmt.Errorf("failed to initialize adb: %v", err)
	}

	_, err = a.ADB.GetState()
	if err != nil {
		return fmt.Errorf("failed to get adb state (are you sure a device is connected?): %v",
			err)
	}

	return nil
}

func (a *Acquisition) createFolder() error {
	a.StoragePath = filepath.Join(utils.GetBinFolder(), a.UUID)
	err := os.Mkdir(a.StoragePath, 0755)
	if err != nil {
		return err
	}

	a.APKSPath = filepath.Join(a.StoragePath, "apks")
	err = os.Mkdir(a.APKSPath, 0755)
	if err != nil {
		return err
	}

	return nil
}

func (a *Acquisition) StoreInfo() error {
	fmt.Println("Saving details about acquisition and device...")

	info, err := json.MarshalIndent(a, "", " ")
	if err != nil {
		return fmt.Errorf("failed to json marshal the acquisition details: %v",
			err)
	}

	infoPath := filepath.Join(a.StoragePath, "acquisition.json")

	err = ioutil.WriteFile(infoPath, info, 0644)
	if err != nil {
		return fmt.Errorf("failed to write acquisition details to file: %v",
			err)
	}

	return nil
}
