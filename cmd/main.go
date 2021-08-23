// androidqf - Android Quick Forensics
// Copyright (c) 2021 Claudio Guarnieri.
// Use of this software is governed by the MVT License 1.1 that can be found at
//   https://license.mvt.re/1.1/

package main

import (
	"fmt"
	"os"
	"time"

	"github.com/botherder/androidqf/acquisition"
	"github.com/i582/cfmt"
)

func init() {
	cfmt.Print(`
	{{                    __           _     __      ____ }}::green
	{{   ____  ____  ____/ /________  (_)___/ /___  / __/ }}::yellow
	{{  / __ '/ __ \/ __  / ___/ __ \/ / __  / __ '/ /_   }}::red
	{{ / /_/ / / / / /_/ / /  / /_/ / / /_/ / /_/ / __/   }}::magenta
	{{ \__,_/_/ /_/\__,_/_/   \____/_/\__,_/\__, /_/      }}::blue
	{{                                        /_/         }}::cyan
	`)
	cfmt.Println("\tandroidqf - Android Quick Forensics")
	cfmt.Println()
}

func systemPause() {
	cfmt.Println("Press {{Enter}}::bold|green to finish ...")
	os.Stdin.Read(make([]byte, 1))
}

func printError(desc string, err error) {
	cfmt.Printf("{{ERROR:}}::red|bold %s: {{%s}}::italic\n",
		desc, err.Error())
}

func main() {
	var acq *acquisition.Acquisition
	var err error

	for true {
		acq, err = acquisition.New()
		if err == nil {
			break
		}

		cfmt.Println("{{ERROR:}}::red|bold Unable to get device state. Please make sure it is connected and authorized. Trying again in 5 seconds...")
		time.Sleep(5 * time.Second)
	}

	cfmt.Printf("Started new acquisition {{%s}}::magenta|underline\n",
		acq.UUID)

	err = acq.GetProp()
	if err != nil {
		printError("Failed to get device properties", err)
	}
	err = acq.Processes()
	if err != nil {
		printError("Failed to get list of running processes", err)
	}
	err = acq.DumpSys()
	if err != nil {
		printError("Failed to get output of dumpsys", err)
	}
	err = acq.DownloadAPKs()
	if err != nil {
		printError("Failed to download APKs", err)
	}
	err = acq.Backup()
	if err != nil {
		printError("Failed to create backup", err)
	}

	err = acq.StoreSecurely()
	if err != nil {
		printError("Something failed while encrypting the acquisition", err)
		cfmt.Println("{{WARNING: The secure storage of the acquisition folder failed! The data is unencrypted!}}::red|bold")
	}

	fmt.Println("Acquisition completed.")

	systemPause()
}
