package main

import (
	"fmt"
	"mcsim/machine"
)

func main() {
	mach := machine.New()
	fmt.Println(mach.Analyze(10))
	fmt.Println(mach.Analyze(50))
	fmt.Println(mach.Analyze(100))
	fmt.Println(mach.Analyze(1000))
}
