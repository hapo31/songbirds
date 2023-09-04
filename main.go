package main

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/hapo31/songbirds/songbirds"
)

func main() {

	wlanInterface, _ := songbirds.LookUpWlanInterface(runtime.GOOS)

	ctx, _, err := songbirds.HTTPServer(8080, func() (bool, string, error) {
		accessPoints, err := songbirds.ScanAccessPoint(wlanInterface, runtime.GOOS)
		if err != nil {
			fmt.Println(err)
			return false, "", err
		}

		for _, o := range accessPoints {
			fmt.Printf("%s\n", o.ESSID)
			if strings.Contains(o.ESSID, "switch") {
				return true, o.ESSID, nil
			}
		}

		return false, "", nil
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("wait...")

	<-ctx.Done()
}
