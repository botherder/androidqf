// androidqf - Android Quick Forensics
// Copyright (c) 2021-2022 Claudio Guarnieri.
// Use of this software is governed by the MVT License 1.1 that can be found at
//   https://license.mvt.re/1.1/

package adb

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type Package struct {
	Name       string        `json:"name"`
	Files      []interface{} `json:"files"`
	Installer  string        `json:"installer"`
	UID        int           `json:"uid"`
	Disabled   bool          `json:"disabled"`
	System     bool          `json:"system"`
	ThirdParty bool          `json:"third_party"`
}

// GetPackages returns the list of installed package names.
func (a *ADB) GetPackages() ([]Package, error) {
	out, err := a.Shell("pm", "list", "packages", "-U", "-u", "-i")
	if err != nil {
		return []Package{}, fmt.Errorf("failed to launch `pm list packages` command: %v",
			err)
	}

	packages := []Package{}
	for _, line := range strings.Split(out, "\n") {
		fields := strings.Fields(line)
		packageName := strings.TrimPrefix(strings.TrimSpace(fields[0]), "package:")
		installer := strings.TrimPrefix(strings.TrimSpace(fields[1]), "installer=")
		uid, _ := strconv.Atoi(strings.TrimPrefix(strings.TrimSpace(fields[2]), "uid:"))

		if packageName == "" {
			continue
		}

		newPackage := Package{
			Name:       packageName,
			Installer:  installer,
			UID:        uid,
			Disabled:   false,
			System:     false,
			ThirdParty: false,
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
		if err != nil {
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
						reflect.ValueOf(&packages[pIndex]).Elem().FieldByName(fieldName).SetBool(true)
					}
				}
			}
		}
	}

	return packages, nil
}

// GetPackagePaths returns a list of file paths associated with the provided
// package name.
func (a *ADB) GetPackagePaths(packageName string) ([]string, error) {
	out, err := a.Shell("pm", "path", packageName)
	if err != nil {
		return []string{}, fmt.Errorf("failed to launch `pm path` command: %v",
			err)
	}

	packagePaths := []string{}
	for _, line := range strings.Split(out, "\n") {
		packagePath := strings.TrimPrefix(strings.TrimSpace(line), "package:")
		if packagePath == "" {
			continue
		}

		packagePaths = append(packagePaths, packagePath)
	}

	return packagePaths, nil
}
