package main

import (
	"flag"
	"fmt"
	"os"
)

import "./jirou"

func Help() {
	fmt.Println(`
JIROU API
	`)
}

func main() {
	// Parse comand line option
	var option = jirou.Option{}
	flag.BoolVar(&option.Help, "help", false, "print help message")
	flag.BoolVar(&option.Setup, "setup", false, "insert records to database from file")
	flag.IntVar(&option.Port, "port", 8080, "set port number")
	flag.Parse()

	if option.Help { // Print help
		Help()
	} else if option.Setup { // Setup
		fmt.Println("setup Jirou API server...")
		jirou.Setup()
	} else { // Execute Server
		server := jirou.NewServer()
		server.Run(&option)
	}
	os.Exit(0)
}
