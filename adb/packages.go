// Copyright (c) 2021-2023 Claudio Guarnieri.
// Use of this source code is governed by the MVT License 1.1
// which can be found in the LICENSE file.

package adb

import (
	"fmt"
	"reflect"
	"strings"
)

type PackageFile struct {
	Path      string `json:"path"`
	LocalName string `json:"local_name"`
	MD5       string `json:"md5"`
	SHA1      string `json:"sha1"`
	SHA256    string `json:"sha256"`
	SHA512    string `json:"sha512"`
}

type Package struct {
	Name       string        `json:"name"`
	Files      []PackageFile `json:"files"`
	Installer  string        `json:"installer"`
	UID        int           `json:"uid"`
	Disabled   bool          `json:"disabled"`
	System     bool          `json:"system"`
	ThirdParty bool          `json:"third_party"`
}

func (a *ADB) getPackageFiles(packageName string) []PackageFile {
	out, err := a.Shell("pm", "path", packageName)
	if err != nil {
		fmt.Printf("Failed to get file paths for package %s: %v: %s\n", packageName, err, out)
		return []PackageFile{}
	}

	packageFiles := []PackageFile{}
	for _, line := range strings.Split(out, "\n") {
		packagePath := strings.TrimPrefix(strings.TrimSpace(line), "package:")
		if packagePath == "" {
			continue
		}

		packageFile := PackageFile{
			Path: packagePath,
		}

		md5Out, err := a.Shell("md5sum", packagePath)
		if err == nil {
			packageFile.MD5 = strings.SplitN(md5Out, " ", 2)[0]
		}
		sha1Out, err := a.Shell("sha1sum", packagePath)
		if err == nil {
			packageFile.SHA1 = strings.SplitN(sha1Out, " ", 2)[0]
		}
		sha256Out, err := a.Shell("sha256sum", packagePath)
		if err == nil {
			packageFile.SHA256 = strings.SplitN(sha256Out, " ", 2)[0]
		}
		sha512Out, err := a.Shell("sha512sum", packagePath)
		if err == nil {
			packageFile.SHA512 = strings.SplitN(sha512Out, " ", 2)[0]
		}

		packageFiles = append(packageFiles, packageFile)
	}

	return packageFiles
}

// GetPackages returns the list of installed package names.
func (a *ADB) GetPackages() ([]Package, error) {
	out, err := a.Shell("pm", "list", "packages", "-u", "-i")
	if err != nil && out == "" {
		return []Package{}, fmt.Errorf("failed to launch `pm list packages` command: %v",
			err)
	}

	packages := []Package{}
	for _, line := range strings.Split(out, "\n") {
		fields := strings.Fields(line)
		packageName := strings.TrimPrefix(strings.TrimSpace(fields[0]), "package:")
		installer := strings.TrimPrefix(strings.TrimSpace(fields[1]), "installer=")

		if packageName == "" {
			continue
		}

		newPackage := Package{
			Name:       packageName,
			Installer:  installer,
			Disabled:   false,
			System:     false,
			ThirdParty: false,
			Files:      a.getPackageFiles(packageName),
		}

		packages = append(packages, newPackage)
	}

	cmds := []map[string]string{
		{"field": "Disabled", "arg": "-d"},
		{"field": "System", "arg": "-s"},
		{"field": "ThirdParty", "arg": "-3"},
	}
	for _, cmd := range cmds {
		out, err = a.Shell("pm", "list", "packages", cmd["arg"])
		if err != nil && out == "" {
			fmt.Printf("Failed to get packages filtered by `%s`: %v: %s\n",
				cmd["arg"], err, out)
			continue
		}

		for _, line := range strings.Split(out, "\n") {
			packageName := strings.TrimPrefix(strings.TrimSpace(line), "package:")
			if packageName == "" {
				continue
			}

			for pIndex, p := range packages {
				if p.Name != packageName {
					continue
				}

				elems := reflect.ValueOf(&p).Elem()
				for i := 0; i < elems.NumField(); i++ {
					fieldName := elems.Type().Field(i).Name
					if fieldName == cmd["field"] {
						reflect.ValueOf(&packages[pIndex]).
							Elem().
							FieldByName(fieldName).
							SetBool(true)
					}
				}
			}
		}
	}

	return packages, nil
}
