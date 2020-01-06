package go_platform

import (
	"regexp"
)

// isCentOS returns true if system OS is centos:
//
//  if osr.isCentOS() {
//      fmt.Println("Centos")
//  } else {
//      fmt.Println("Not centos")
//  }
//
func (o *OSRelease) isCentOS(instring []byte) bool {
	re := regexp.MustCompile(`CentOS Linux release ((\d+)\.(\d+)\.(\d+))`)
	matches := re.FindSubmatch(instring)
	if len(matches) ==  0 {
		return false
	}
	o.Name = "CentOS"
	o.MajorVersion = string(matches[2])
	o.Version = string(matches[2]) + "." + string(matches[3])
	return true
}
