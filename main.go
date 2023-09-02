package main

import (
	"fmt"
	"runtime"

	"github.com/hapo31/songbirds/songbirds"
)

func main() {
	interfaces, _ := songbirds.EnumInterfaces(runtime.GOOS)

	for _, i := range interfaces {
		fmt.Println(i)
	}
	fmt.Println("finish.")
}
