package go_platform

import (
	"regexp"
)

// isRedHat returns true if system OS is centos:
//
//  if osr.isRedHat() {
//      fmt.Println("Red Hat")
//  } else {
//      fmt.Println("Not Red Hat")
//  }
//
func (o *OSRelease) isRedHat(instring []byte) bool {
	re := regexp.MustCompile(`Red Hat Enterprise Linux Server release ((\d+)\.(\d+))`)
	matches := re.FindSubmatch(instring)
	if len(matches) == 0 {
		return false
	}
	o.Name = "RedHat"
	o.MajorVersion = string(matches[2])
	o.Version = string(matches[2]) + "." + string(matches[3])
	return true
}
