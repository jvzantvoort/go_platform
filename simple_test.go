package go_platform_test

import (
	"testing"
        "github.com/jvzantvoort/go_platform"
	"gotest.tools/assert"
)

const (
	KERNEL = "linux"
	NAME = "CentOS"
	MAJORVERSION = "7"
	VERSION = "7.7"
)

func TestPlatformName(t *testing.T) {
	platform_obj := go_platform.OSRelease{}
	platform_obj.GetPlatform()
	assert.Equal(t, platform_obj.Name, NAME)
}

func TestPlatformVersion(t *testing.T) {
	platform_obj := go_platform.OSRelease{}
	platform_obj.GetPlatform()
	assert.Equal(t, platform_obj.Version, VERSION)
}

func TestPlatformMajorVersion(t *testing.T) {
	platform_obj := go_platform.OSRelease{}
	platform_obj.GetPlatform()
	assert.Equal(t, platform_obj.MajorVersion, MAJORVERSION)
}

