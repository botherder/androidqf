// Copyright (c) 2021-2023 Claudio Guarnieri.
// Use of this source code is governed by the MVT License 1.1
// which can be found in the LICENSE file.

package acquisition

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	rt "github.com/botherder/go-savetime/runtime"
	"github.com/google/uuid"
)

type Acquisition struct {
	UUID        string    `json:"uuid"`
	StoragePath string    `json:"storage_path"`
	Started     time.Time `json:"started"`
	Completed   time.Time `json:"completed"`
}

func New() (*Acquisition, error) {
	acq := Acquisition{
		UUID:    uuid.New().String(),
		Started: time.Now().UTC(),
	}

	acq.StoragePath = filepath.Join(rt.GetExecutableDirectory(), acq.UUID)
	err := os.Mkdir(acq.StoragePath, 0o755)
	if err != nil {
		return nil, fmt.Errorf("failed to create acquisition folder: %v", err)
	}

	return &acq, nil
}

func (a *Acquisition) Complete() {
	a.Completed = time.Now().UTC()
}

func (a *Acquisition) StoreInfo() error {
	fmt.Println("Saving details about acquisition and device...")

	info, err := json.MarshalIndent(a, "", " ")
	if err != nil {
		return fmt.Errorf("failed to json marshal the acquisition details: %v",
			err)
	}

	infoPath := filepath.Join(a.StoragePath, "acquisition.json")

	err = ioutil.WriteFile(infoPath, info, 0o644)
	if err != nil {
		return fmt.Errorf("failed to write acquisition details to file: %v",
			err)
	}

	return nil
}
