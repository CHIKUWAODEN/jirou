package main

import (
	"fmt"
	"os"
)

import "./jirou"

func main() {
	// Parse comand line option
	option := jirou.ParseOption()

	if option.Help {
		jirou.Help()
	} else if option.Setup {
		fmt.Println("setup Jirou API server...")
		jirou.Setup()
	} else {
		server := jirou.NewServer()
		server.Run(option)
	}
	os.Exit(0)
}
