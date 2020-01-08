// Package go_platform allows for simple identification of platforms.
//
// Source code and other details for the project are available at GitHub:
//
//   https://github.com/jvzantvoort/go_platform
//

package go_platform

import (
	"bufio"
	"log"
	"os"
	"runtime"
)

// Struct to contain the relevant variables
type OSRelease struct {
	Kernel       string
	MajorVersion string
	Version      string
	Name         string
}

// LSBInfo contains parameters and ...
type LSBInfo struct {
	Filename   string
	Name       string
	AllowEmpty bool
}

// lsbInfoSets a list of structs with options
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

// getInfoSets gets the info sets relevant to us
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

// parseFile example
func (o *OSRelease) parseFile(filename string) bool {
	retv := false
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		instring := scanner.Bytes()
		if o.isCentOS(instring) {
			retv = true
			break
		}
		if o.isRedHat(instring) {
			retv = true
			break
		}

	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return retv
}

// GetPlatform set the Kernel, Version and Name:
//  import (
//          "fmt"
//          "github.com/jvzantvoort/go_platform"
//  )
//  platf := go_platform.OSRelease{}
//  platf.GetPlatform()
//  fmt.Printf("name=%s\n", platf.Name)
//  fmt.Printf("version=%s\n", platf.Version)
//  fmt.Printf("major_version=%s\n", platf.MajorVersion)
//
func (o *OSRelease) GetPlatform() {
	o.Kernel = runtime.GOOS
	if o.Kernel != "linux" {
		return
	}
	o.Version = "undef"
	o.Name = "undef"
	infosets := getInfoSets()
	for _, y := range infosets {
		o.parseFile(y.Filename)
	}
}

// vim: noexpandtab filetype=go
