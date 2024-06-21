package main

import (
	"fmt"
	"os"
	"runtime"

	"golang.org/x/sys/cpu"
)

func main() {
	fmt.Printf("cpu.IsBigEndian: %t\n", cpu.IsBigEndian)
	fmt.Printf("os.Getpagesize(): %d\n", os.Getpagesize())
	fmt.Printf("runtime.GOOS: %s\n", runtime.GOOS)
	fmt.Printf("runtime.GOARCH: %s\n", runtime.GOARCH)
}
