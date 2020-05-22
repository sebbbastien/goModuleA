package gomodulea

import "fmt"

func init() {
	fmt.Println("Loading module A")
}

func FctFromModuleA() {
	fmt.Println("Calling FctFromModuleA() v2")
}
