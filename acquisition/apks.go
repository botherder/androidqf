// androidqf - Android Quick Forensics
// Copyright (c) 2021-2022 Claudio Guarnieri.
// Use of this software is governed by the MVT License 1.1 that can be found at
//   https://license.mvt.re/1.1/

package acquisition

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/i582/cfmt/cmd/cfmt"
	"github.com/manifoldco/promptui"
)

const (
	apkAll       = "All"
	apkNotSystem = "Only non-system packages"
	apkNone      = "Do not download any"
)

type File struct {
	Path      string `json:"path"`
	LocalName string `json:"local_name"`
	SHA256    string `json:"sha256"`
}

func (a *Acquisition) getPathToLocalCopy(packageName, filePath string) string {
	fileName := ""
	if strings.Contains(filePath, "==/") {
		fileName = fmt.Sprintf("_%s", strings.Replace(strings.Split(filePath, "==/")[1], ".apk", "", 1))
	}

	localPath := filepath.Join(a.APKSPath, fmt.Sprintf("%s%s.apk", packageName, fileName))
	counter := 0
	for true {
		if _, err := os.Stat(localPath); os.IsNotExist(err) {
			break
		}

		counter++
		localPath = filepath.Join(a.APKSPath, fmt.Sprintf("%s%s_%s.apk", packageName, fileName, counter))
	}

	return localPath
}

func (a *Acquisition) DownloadAPKs() error {
	fmt.Println("Downloading copies of installed apps...")

	packages, err := a.ADB.GetPackages()
	if err != nil {
		return fmt.Errorf("failed to retrieve list of installed packages: %v", err)
	}

	cfmt.Printf("Found a total of {{%d}}::cyan|bold installed packages\n",
		len(packages))

	fmt.Println("Would you like to download copies of all apps or only non-system ones?")
	promptAll := promptui.Select{
		Label: "Download",
		Items: []string{apkAll, apkNotSystem, apkNone},
	}
	_, downloadOption, err := promptAll.Run()
	if err != nil {
		return fmt.Errorf("failed to make selection for download option: %v",
			err)
	}

	// If the user decides to not download any APK, then we skip this.
	// Otherwise we walk through the list of package, pull the files, and hash them.
	if downloadOption != apkNone {
		for _, p := range packages {
			// If we the user did not request to download all packages and if
			// the package is marked as system, we skip it.
			if downloadOption != apkAll && p.System == true {
				continue
			}

			cfmt.Printf("Found Android package: {{%s}}::cyan|bold\n", p.Name)

			pFilePaths, err := a.ADB.GetPackagePaths(p.Name)
			if err == nil {
				for _, pFilePath := range pFilePaths {
					localPath := a.getPathToLocalCopy(p.Name, pFilePath)

					out, err := a.ADB.Pull(pFilePath, localPath)
					if err != nil {
						cfmt.Printf("{{ERROR:}}::red|bold Failed to download {{%s}}::cyan|underline: {{%s}}::italic\n",
							pFilePath, out)

						continue
					}

					cfmt.Printf("Downloaded {{%s}}::cyan|underline to {{%s}}::magenta|underline\n",
						pFilePath, localPath)
				}
			}
		}
	}

	// Store the results into a JSON file.
	packagesJSONPath := filepath.Join(a.StoragePath, "packages.json")
	packagesJSON, err := os.Create(packagesJSONPath)
	if err != nil {
		return fmt.Errorf("failed to save list of installed packages to file: %v",
			err)
	}
	defer packagesJSON.Close()

	buf, _ := json.MarshalIndent(packages, "", "    ")

	packagesJSON.WriteString(string(buf[:]))
	packagesJSON.Sync()

	return nil
}
