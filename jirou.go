package main

import (
	"flag"
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
)

// Represents command line option
type option struct {
	help bool
	port int
}

func Help() {
	fmt.Println("JIROU API")
}

func main() {
	// Parse comand line option
	var option = option{}
	flag.BoolVar(&option.help, "help", false, "print help message")
	flag.IntVar(&option.port, "port", 8080, "set port number")
	flag.Parse()

	if option.help {
		Help()
		os.Exit(0)
	}

	// HTTP server
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		var response string = fmt.Sprintf("Hello, %q", html.EscapeString(request.URL.Path))
		fmt.Fprintf(writer, response)
	})

	var p string = fmt.Sprintf(":%d", option.port)
	log.Fatal(http.ListenAndServe(p, nil))
}
