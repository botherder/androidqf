// Copyright (c) 2021-2023 Claudio Guarnieri.
// Use of this source code is governed by the MVT License 1.1
// which can be found in the LICENSE file.

package modules

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/botherder/androidqf/adb"
	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
)

const (
	apkAll       = "All"
	apkNotSystem = "Only non-system packages"
	apkNone      = "Do not download any"
)

type Packages struct {
	StoragePath string
	ApksPath    string
}

func NewPackages() *Packages {
	return &Packages{}
}

func (p *Packages) Name() string {
	return "packages"
}

func (p *Packages) InitStorage(storagePath string) error {
	p.StoragePath = storagePath
	p.ApksPath = filepath.Join(storagePath, "apks")
	err := os.Mkdir(p.ApksPath, 0o755)
	if err != nil {
		return fmt.Errorf("failed to create apks folder: %v", err)
	}

	return nil
}

func (p *Packages) getPathToLocalCopy(packageName, filePath string) string {
	fileName := ""
	if strings.Contains(filePath, "==/") {
		fileName = fmt.Sprintf(
			"_%s",
			strings.Replace(strings.Split(filePath, "==/")[1], ".apk", "", 1),
		)
	}

	localPath := filepath.Join(p.ApksPath, fmt.Sprintf("%s%s.apk", packageName, fileName))
	counter := 0
	for {
		if _, err := os.Stat(localPath); os.IsNotExist(err) {
			break
		}

		counter++
		localPath = filepath.Join(
			p.ApksPath,
			fmt.Sprintf("%s%s_%d.apk", packageName, fileName, counter),
		)
	}

	return localPath
}

func (p *Packages) Run() error {
	fmt.Println("Collecting information on installed apps. This might take a while...")

	packages, err := adb.Client.GetPackages()
	if err != nil {
		return fmt.Errorf("failed to retrieve list of installed packages: %v", err)
	}

	fmt.Printf(
		"Found a total of %s installed packages\n",
		color.CyanString(strconv.Itoa(len(packages))),
	)

	fmt.Println("Would you like to download copies of all apps or only non-system ones?")
	downloadPrompt := promptui.Select{
		Label: "Download",
		Items: []string{apkAll, apkNotSystem, apkNone},
	}
	_, download, err := downloadPrompt.Run()
	if err != nil {
		return fmt.Errorf("failed to make selection for download option: %v", err)
	}

	// If the user decides to not download any APK, then we skip this.
	// Otherwise we walk through the list of package, pull the files, and hash them.
	if download != apkNone {
		for _, pack := range packages {
			// If we the user did not request to download all packages and if
			// the package is marked as system, we skip it.
			if download != apkAll && pack.System {
				continue
			}

			fmt.Printf("Found Android package: %s\n", color.CyanString(pack.Name))

			for _, packageFile := range pack.Files {
				localPath := p.getPathToLocalCopy(pack.Name, packageFile.Path)

				out, err := adb.Client.Pull(packageFile.Path, localPath)
				if err != nil {
					color.Red(
						fmt.Sprintf("ERROR: failed to download %s: %s", packageFile.Path, out),
					)
					continue
				}

				fmt.Printf(
					"Downloaded %s to %s\n",
					color.CyanString(packageFile.Path),
					color.MagentaString(localPath),
				)
			}
		}
	}

	resultsPath := filepath.Join(p.StoragePath, "packages.json")
	results, err := os.Create(resultsPath)
	if err != nil {
		return fmt.Errorf("failed to save list of installed packages to file: %v", err)
	}
	defer results.Close()

	buf, _ := json.MarshalIndent(packages, "", "    ")

	results.WriteString(string(buf[:]))
	results.Sync()

	return nil
}
