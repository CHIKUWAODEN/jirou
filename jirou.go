package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

import (
	"github.com/julienschmidt/httprouter"
)

// Represents command line option
type option struct {
	help bool
	port int
}

func Help() {
	fmt.Println("JIROU API")
}

func Root(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	// [todo] - 文字列を []byte にキャストするとメモリコピーが走っちゃうらしい
	response := []byte(`
	{
		"link" : {
			"root"  : { "method" : "GET", "uri" : "/"   },
			"index" : { "method" : "GET", "uri" : "/v1" }
		}
	}
	`)
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(response)
}

func V1Root(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	// [todo] - 文字列を []byte にキャストするとメモリコピーが走っちゃうらしい
	response := []byte(`
	{
		"link" : {
			"root"   : { "method" : "GET",  "uri" : "/" },
			"index"  : { "method" : "GET",  "uri" : "/v1" },
			"create" : { "method" : "POST", "uri" : "/v1/jirou" } 
		}
	}
	`)
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(response)
}

func Index(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json")
	http.Error(
		writer, "API Under construction", http.StatusNotImplemented)
}

func Create(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json")
	http.Error(
		writer, "API Under construction", http.StatusNotImplemented)
}

func Read(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json")
	http.Error(
		writer, "API Under construction", http.StatusNotImplemented)
}

func Update(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json")
	http.Error(
		writer, "API Under construction", http.StatusNotImplemented)
}

func Delete(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json")
	http.Error(
		writer, "API Under construction", http.StatusNotImplemented)
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

	// Build Routing
	router := httprouter.New()
	router.GET("/", Root)
	router.GET("/v1", V1Root)
	router.GET("/v1/jirou", Index)
	router.POST("/v1/jirou", Create)
	router.GET("/v1/jirou/:id", Read)
	router.PUT("/v1/jirou/:id", Update)
	router.DELETE("/v1/jirou/:id", Delete)

	// Start serving
	fmt.Println("Starting API server, ", option)
	var port string = fmt.Sprintf(":%d", option.port)
	log.Fatal(http.ListenAndServe(port, router))
}
