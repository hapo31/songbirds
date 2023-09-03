package main

import (
	"fmt"
	"runtime"

	"github.com/hapo31/songbirds/songbirds"
)

func main() {
	wlanInterface, _ := songbirds.LookUpWlanInterface(runtime.GOOS)

	fmt.Printf("%s\n", wlanInterface)
	fmt.Println("finish.")
}
