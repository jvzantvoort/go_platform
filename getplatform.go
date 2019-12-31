//          FILE:  getplatform.go
//
//         USAGE:  getplatform.go
//
//   DESCRIPTION:  $description
//
//       OPTIONS:  ---
//  REQUIREMENTS:  ---
//          BUGS:  ---
//         NOTES:  ---
//        AUTHOR:  John van Zantvoort (jvzantvoort), john@vanzantvoort.org
//       COMPANY:  JDC
//       CREATED:  31-Dec-2019
//
// Copyright (C) 2019 JDC
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to
// deal in the Software without restriction, including without limitation the
// rights to use, copy, modify, merge, publish, distribute, sublicense, and/or
// sell copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS
// IN THE SOFTWARE.
//
package go_platform

import (
	"io/ioutil"
	"os"
	"fmt"
	"runtime"
)

type OSRelease struct {
	Version string
	Name    string
}

type LSBInfo struct {
	Filename string
	Name     string
	AllowEmpty bool
}

var lsbInfoSets = []LSBInfo{
	LSBInfo{Filename: "/etc/alpine-release", Name: "Alpine"},
	LSBInfo{Filename: "/etc/altlinux-release", Name: "Altlinux"},
	LSBInfo{Filename: "/etc/arch-release", Name: "Archlinux", AllowEmpty: true},
	LSBInfo{Filename: "/etc/centos-release", Name: "CentOS"},
	LSBInfo{Filename: "/etc/coreos/update.conf", Name: "Coreos"},
	LSBInfo{Filename: "/etc/debian_version", Name: "Debian"},
	LSBInfo{Filename: "/etc/fedora-release", Name: "Fedora"},
	LSBInfo{Filename: "/etc/gentoo-release", Name: "Gentoo"},
	LSBInfo{Filename: "/etc/lsb-release", Name: "Debian"},
	LSBInfo{Filename: "/etc/lsb-release", Name: "Mandriva"},
	LSBInfo{Filename: "/etc/mandrake-release", Name: "Mandrake"},
	LSBInfo{Filename: "/etc/openwrt_release", Name: "OpenWrt"},
	LSBInfo{Filename: "/etc/oracle-release", Name: "OracleLinux"},
	LSBInfo{Filename: "/etc/os-release", Name: "Archlinux"},
	LSBInfo{Filename: "/etc/os-release", Name: "Debian"},
	LSBInfo{Filename: "/etc/os-release", Name: "NA"},
	LSBInfo{Filename: "/etc/os-release", Name: "SUSE"},
	LSBInfo{Filename: "/etc/redhat-release", Name: "RedHat"},
	LSBInfo{Filename: "/etc/slackware-version", Name: "Slackware"},
	LSBInfo{Filename: "/etc/snow-release", Name: "Snow"},
	LSBInfo{Filename: "/etc/sourcemage-release", Name: "SMGL"},
	LSBInfo{Filename: "/etc/SuSE-release", Name: "SUSE"},
	LSBInfo{Filename: "/etc/system-release", Name: "Amazon"},
	LSBInfo{Filename: "/etc/vmware-release", Name: "VMwareESX", AllowEmpty: true},
	LSBInfo{Filename: "/etc/wrs-release", Name: "WindRiver"},
	LSBInfo{Filename: "/usr/lib/os-release", Name: "ClearLinux"},
}

// get the info sets relevant to us
func getInfoSets() []LSBInfo {
	var retv = []LSBInfo{}
	for _, i := range lsbInfoSets {
		if _, err := os.Stat(i.Filename); os.IsNotExist(err) {
			continue
		}
		retv = append(retv, i)
	}
	return retv
}

/*
    SEARCH_STRING = {
        'OracleLinux': 'Oracle Linux',
        'RedHat': 'Red Hat',
        'Altlinux': 'ALT',
        'SMGL': 'Source Mage GNU/Linux',
    }
    */
func (o *OSRelease) GetPlatform() {
	if runtime.GOOS != "linux" {
		return
	}
	infosets := getInfoSets()
	for _, y := range infosets {
		content, err := ioutil.ReadFile(y.Filename)
		if err != nil {
			panic(err)
		}
		fmt.Print(string(content))
	}
}

// vim: noexpandtab filetype=go
