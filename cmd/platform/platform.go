package main

import (
    "fmt"
    "github.com/jvzantvoort/go_platform"
)

func main() {
	platf := go_platform.OSRelease{}
	platf.GetPlatform()
	fmt.Printf("name=%s\n", platf.Name)
	fmt.Printf("version=%s\n", platf.Version)
	fmt.Printf("major_version=%s\n", platf.MajorVersion)
}
