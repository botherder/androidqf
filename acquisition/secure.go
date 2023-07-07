// Copyright (c) 2021-2023 Claudio Guarnieri.
// Use of this source code is governed by the MVT License 1.1
// which can be found in the LICENSE file.

package acquisition

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"filippo.io/age"
	"github.com/botherder/go-savetime/files"
	rt "github.com/botherder/go-savetime/runtime"
)

func (a *Acquisition) StoreSecurely() error {
	cwd := rt.GetExecutableDirectory()

	keyFilePath := filepath.Join(cwd, "key.txt")
	if _, err := os.Stat(keyFilePath); os.IsNotExist(err) {
		return nil
	}

	fmt.Println("You provided an age public key, storing the acquisition securely.")

	zipFileName := fmt.Sprintf("%s.zip", a.UUID)
	zipFilePath := filepath.Join(cwd, zipFileName)

	fmt.Println("Compressing the acquisition folder. This might take a while...")

	err := files.Zip(a.StoragePath, zipFilePath)
	if err != nil {
		return err
	}

	fmt.Println("Encrypting the compressed archive. This might take a while...")

	publicKey, err := ioutil.ReadFile(keyFilePath)
	if err != nil {
		return err
	}
	publicKeyStr := strings.TrimSpace(string(publicKey))

	recipient, err := age.ParseX25519Recipient(publicKeyStr)
	if err != nil {
		return fmt.Errorf("failed to parse public key %q: %v", publicKeyStr, err)
	}

	zipFile, err := os.Open(zipFilePath)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	encFileName := fmt.Sprintf("%s.age", zipFileName)
	encFilePath := filepath.Join(cwd, encFileName)
	encFile, err := os.OpenFile(encFilePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0o600)
	if err != nil {
		return fmt.Errorf("unable to create encrypted file: %v", err)
	}
	defer encFile.Close()

	w, err := age.Encrypt(encFile, recipient)
	if err != nil {
		return fmt.Errorf("failed to create encrypted file: %v", err)
	}

	_, err = io.Copy(w, zipFile)
	if err != nil {
		return fmt.Errorf("failed to write to encrypted file: %v", err)
	}

	if err := w.Close(); err != nil {
		return fmt.Errorf("failed to close encrypted file: %v", err)
	}

	fmt.Println("Acquisition successfully encrypted at ", encFilePath)

	// TODO: we should securely wipe the files.
	zipFile.Close()
	err = os.Remove(zipFilePath)
	if err != nil {
		return fmt.Errorf("failed to delete the unencrypted compressed archive: %v", err)
	}
	err = os.RemoveAll(a.StoragePath)
	if err != nil {
		return fmt.Errorf("failed to delete the original unencrypted acquisition folder: %v", err)
	}

	return nil
}
