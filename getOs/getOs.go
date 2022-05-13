package main

import (
	"fmt"
	"runtime"
)

func main() {
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "windows":
		fmt.Println("windows")
	case "linux":
		fmt.Println("linux")
	default:
		fmt.Println("os \r\n" + os)
	}
}
