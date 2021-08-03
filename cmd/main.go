// Snoopdroid 2
// Copyright (c) 2021 Claudio Guarnieri.
// Use of this software is governed by the MVT License 1.1 that can be found at
//   https://license.mvt.re/1.1/

package main

import (
	"os"
	"time"

	"github.com/botherder/snoopdroid2/acquisition"
	"github.com/i582/cfmt"
)

func systemPause() {
	cfmt.Println("Press {{Enter}}::bold|green to finish ...")
	os.Stdin.Read(make([]byte, 1))
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

	cfmt.Printf("Started new acquisition at: {{%s}}::magenta|underline\n", acq.BasePath)

	err = acq.GetProp()
	if err != nil {
		cfmt.Printf("{{ERROR:}}::red|bold {{%s}}::italic\n", err)
	}
	err = acq.Processes()
	if err != nil {
		cfmt.Printf("{{ERROR:}}::red|bold {{%s}}::italic\n", err)
	}
	err = acq.DumpSys()
	if err != nil {
		cfmt.Printf("{{ERROR:}}::red|bold {{%s}}::italic\n", err)
	}
	err = acq.DownloadAPKs()
	if err != nil {
		cfmt.Printf("{{ERROR:}}::red|bold {{%s}}::italic\n", err)
	}
	err = acq.Backup()
	if err != nil {
		cfmt.Printf("{{ERROR:}}::red|bold {{%s}}::italic\n", err)
	}

	cfmt.Printf("Acquisition completed. The results are stored at: {{%s}}::magenta|underline\n", acq.BasePath)

	systemPause()
}
