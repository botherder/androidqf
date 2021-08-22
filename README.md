# androidqf

[![Go Report Card](https://goreportcard.com/badge/github.com/botherder/androidqf)](https://goreportcard.com/report/github.com/botherder/androidqf)

androidqf (Android Quick Forensics) is a portable tool to simplify the acquisition of relevant forensic data from Android devices. It is the successor of [Snoopdroid](https://github.com/botherder/snoopdroid), re-written in Go and leveraging official adb binaries.

androidqf is intended to provide a simple and portable cross-platform utility to quickly acquire data from Android devices. It is similar in functionality to [mvt-android](https://github.com/mvt-project/mvt). However, contrary to MVT, androidqf is designed to be easily run by non-tech savvy users as well.

[Download androidqf](https://github.com/botherder/androidqf/releases/latest)

## Build

Executable binaries for Linux, Windows and Mac should be available in the [latest release](https://github.com/botherder/androidqf/releases/latest). In case you have issues running the binary you might want to build it by yourself.

In order to build androidqf you will need Go 1.15+ installed. You will also need to install `make`. When ready you can clone the repository and run any of the following commands, depending on which platform you would like to run androidqf on:

    make linux
    make darwin
    make windows

These commands will generate binaries in a *build/* folder.

## How to use

Before launching androidqf you need to have the target Android device connected to your computer via USB, and you will need to have enabled USB debugging. Please refer to the [official documentation](https://developer.android.com/studio/debug/dev-options#enable) on how to do this, but also be mindful that Android phones from different manufacturers might require different navigation steps than the defaults.

Once USB debugging is enabled, you can proceed launching androidqf. It will first attempt to connect to the device over the USB bridge, which should result in the Android phone to prompt you to manually authorize the host keys. Make sure to authorize them, ideally permanently so that the prompt wouldn't appear again.

Now androidqf should be executing and creating an acquisition folder at the same path you have placed your androidqf binary. At some point in the execution, androidqf will prompt you some choices: these prompts will pause the acquisition until you provide a selection, so pay attention.

The following data can be extracted:

1. A list of all packages installed and related distribution files.
2. (Optional) Copy of all installed APKs or of only those not safelisted by androidqf.
3. The output of the `dumpsys` shell command, providing diagnostic information about the device.
4. The output of the `getprop` shell command, providing build information and configuration parameters.
5. The output of the `ps` shell command, providing a list of all running processes.
6. (Optional) A backup of SMS and MMS messages.

## License

The purpose of androidqf is to facilitate the ***consensual forensic analysis*** of devices of those who might be targets of sophisticated mobile spyware attacks, especially members of civil society and marginalized communities. We do not want androidqf to enable privacy violations of non-consenting individuals. Therefore, the goal of this license is to prohibit the use of androidqf (and any other software licensed the same) for the purpose of *adversarial forensics*.

In order to achieve this androidqf is released under [MVT License 1.1](https://license.mvt.re/1.1/), an adaptation of [Mozilla Public License v2.0](https://www.mozilla.org/MPL). This modified license includes a new clause 3.0, "Consensual Use Restriction" which permits the use of the licensed software (and any *"Larger Work"* derived from it) exclusively with the explicit consent of the person/s whose data is being extracted and/or analysed (*"Data Owner"*).
