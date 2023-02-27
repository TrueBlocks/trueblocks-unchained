package main

import (
	"fmt"
)

func main() {
	cmd := getCommand()
	switch cmd {
	case "publish":
		Publish()
	case "get":
		Get()
	case "add-name":
		AddName()
	default:
		fmt.Println("Unknown command:", cmd)
	}
}
