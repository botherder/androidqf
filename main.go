// Copyright (c) 2021-2023 Claudio Guarnieri.
// Use of this source code is governed by the MVT License 1.1
// which can be found in the LICENSE file.

package main

import (
	"fmt"
	"os"
	"time"

	"github.com/botherder/androidqf/acquisition"
	"github.com/botherder/androidqf/adb"
	"github.com/botherder/androidqf/modules"
	"github.com/fatih/color"
)

func fatal(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	color.Red("FATAL: %s", msg)
	os.Exit(1)
}

func fatalIfError(err error, format string, args ...interface{}) {
	if err != nil {
		fatal(format, args)
	}
}

func init() {
	fmt.Println()
	color.Green("   █████  ███    ██ ██████  ██████   ██████  ██ ██████   ██████  ███████")
	color.Yellow("  ██   ██ ████   ██ ██   ██ ██   ██ ██    ██ ██ ██   ██ ██    ██ ██")
	color.Red("  ███████ ██ ██  ██ ██   ██ ██████  ██    ██ ██ ██   ██ ██    ██ █████")
	color.Magenta("  ██   ██ ██  ██ ██ ██   ██ ██   ██ ██    ██ ██ ██   ██ ██ ▄▄ ██ ██")
	color.Blue("  ██   ██ ██   ████ ██████  ██   ██  ██████  ██ ██████   ██████  ██")
	color.Cyan("                                                            ▀▀")
	fmt.Println("     androidqf - Android Quick Forensics")
	fmt.Println("     https://github.com/botherder/androidqf")
	fmt.Println()
	fmt.Println(
		"In order to use androidqf, the device needs to be authorized and have USB debugging enabled.",
	)
	fmt.Println("Please follow the these instructions if you haven't configured the device yet:")
	fmt.Println("  https://developer.android.com/studio/debug/dev-options#enable")
	fmt.Println()
}

func main() {
	var err error
	adb.Client, err = adb.New()
	fatalIfError(err, "%v", err)

	for {
		_, err = adb.Client.GetState()
		if err == nil {
			break
		}

		color.Red(
			"ERROR: Unable to get device state. Please make sure it is connected and authorized. Trying again in 5 seconds...",
		)
		time.Sleep(5 * time.Second)
	}

	acq, err := acquisition.New()
	fatalIfError(err, "%v", err)

	fmt.Printf("Started new acquisition %s\n", color.MagentaString(acq.UUID))

	mods := modules.List()
	for _, mod := range mods {
		err = mod.InitStorage(acq.StoragePath)
		if err != nil {
			color.Red(
				fmt.Sprintf(
					"ERROR: failed to initialize storage for module %s: %v",
					mod.Name(),
					err,
				),
			)
			continue
		}

		err = mod.Run()
		if err != nil {
			color.Red("ERROR: failed to run module %s: %v", mod.Name(), err)
		}
	}

	acq.Complete()

	err = acq.StoreSecurely()
	if err != nil {
		color.Red("Something failed while encrypting the acquisition: %v", err)
		color.Red(
			"WARNING: The secure storage of the acquisition folder failed! The data is unencrypted!",
		)
	}

	fmt.Println("Acquisition completed!")

	fmt.Printf("Press %s to finish ...\n", color.GreenString("Enter"))
	os.Stdin.Read(make([]byte, 1))
}
