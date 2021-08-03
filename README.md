# Snoopdroid 2

[![Go Report Card](https://goreportcard.com/badge/github.com/botherder/snoopdroid2)](https://goreportcard.com/report/github.com/botherder/snoopdroid2)

Snoopdroid 2 is a portable tool to simplify the acquisition of relevant forensic data from Android devices. It is the successor of [Snoopdroid](https://github.com/botherder/snoopdroid), re-written in Go and leveraging official adb binaries.

## Build

In order to build Snoopdroid 2 you will need Go 1.15+ installed. You will also need to install `make`. When ready you can clone the repository and run any of the following commands, depending on which platform you would like to run Snoopdroid 2 on:

    make linux
    make darwin
    make windows

These commands will generate binaries in a *build/* folder.

## Use

Before launching Snoopdroid you need to have the target Android device connected to your computer via USB, and you will need to have enabled USB debugging. Please refer to the [official documentation](https://developer.android.com/studio/debug/dev-options#enable) on how to do this, but also be mindful that Android phones from different manufacturers might require different navigation steps than the defaults.

Once USB debugging is enabled, you can proceed launching Snoopdroid 2. It will first attempt to connect to the device over the USB bridge, which should result in the Android phone to prompt you to manually authorize the host keys. Make sure to authorize them, ideally permanently so that the prompt wouldn't appear again.

Now Snoopdroid 2 should be executing and creating an acquisition folder at the same path you have placed your Snoopdroid 2 binary. At some point in the execution, Snoopdroid 2 will prompt you some choices: these prompts will pause the acquisition until you provide a selection, so pay attention.

The following data can be extracted:

1. A list of all packages installed and related distribution files.
2. (Optional) Copy of all installed APKs or of only those not safelisted by Snoopdroid 2.
3. The output of the `dumpsys` shell command, providing diagnostic information about the device.
4. The output of the `getprop` shell command, providing build information and configuration parameters.
5. The output of the `ps` shell command, providing a list of all running processes.
6. (Optional) A backup of SMS and MMS messages.
