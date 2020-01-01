//          FILE:  platform.go
// 
//         USAGE:  platform.go
// 
//   DESCRIPTION:  $description
// 
//       OPTIONS:  ---
//  REQUIREMENTS:  ---
//          BUGS:  ---
//         NOTES:  ---
//        AUTHOR:  John van Zantvoort (jvzantvoort), john@vanzantvoort.org
//       COMPANY:  JDC
//       CREATED:  01-Jan-2020
//
// Copyright (C) 2020 JDC
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
// vim: noexpandtab filetype=go
