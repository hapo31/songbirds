package main

import (
	"fmt"
	"runtime"

	"github.com/hapo31/songbirds/songbirds"
)

func main() {
	wlanInterface, _ := songbirds.LookUpWlanInterface(runtime.GOOS)

	fmt.Printf("%s\n", wlanInterface)

	accessPoints, err := songbirds.ScanAccessPoint(wlanInterface, runtime.GOOS)

	if err != nil {
		fmt.Println(err)
	}

	if len(accessPoints) <= 0 {
		fmt.Println("No entiry scan")
	}

	for _, o := range accessPoints {
		fmt.Printf("%v\n", o)
	}

	fmt.Println("finish.")
}
