package jirou

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

import (
	"github.com/julienschmidt/httprouter"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/opt"
)

type Server struct {
	Router *httprouter.Router
	Db     *leveldb.DB
}

func NewServer() *Server {
	var s = new(Server)
	s.Router = httprouter.New()
	s.Db = nil
	return s
}

func (self *Server) Root(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
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

func (self *Server) V1Root(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	// [todo] - 文字列を []byte にキャストするとメモリコピーが走っちゃうらしい
	response := []byte(`
	{
		"link" : {
			"root"   : { "method" : "GET",  "uri" : "/" },
			"index"  : { "method" : "GET",  "uri" : "/v1" }
		}
	}
	`)
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(response)
}

func (self *Server) Search(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {

	values := make([]string, 0, 5)
	iter := self.Db.NewIterator(nil, nil)
	for iter.Next() {
		values = append(values, string(iter.Value()))
	}
	iter.Release()
	err := iter.Error()
	if nil != err {
		// Error
	}

	content := strings.Join(values, ",")
	response := []byte(fmt.Sprintf(`
		{
			"link" : {
				"root"   : { "method" : "GET",  "uri" : "/" },
				"index"  : { "method" : "GET",  "uri" : "/v1" }
			},
			"content" : [
				%v
			]
		}`,
		content,
	))

	writer.Header().Set("Content-Type", "application/json")
	writer.Write(response)
}

// func (self *Server) Create(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
// 	writer.Header().Set("Content-Type", "application/json")
// 	http.Error(
// 		writer, "API Under construction", http.StatusNotImplemented)
// }

func (self *Server) Read(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	key := []byte(params.ByName("id"))
	content, _ := self.Db.Get(key, nil)

	response := []byte(fmt.Sprintf(`
		{
			"link" : {
				"root"   : { "method" : "GET",  "uri" : "/" },
				"index"  : { "method" : "GET",  "uri" : "/v1" }
			},
			"content" : %v
		}`,
		string(content),
	))

	writer.Header().Set("Content-Type", "application/json")
	writer.Write(response)
}

// func (self *Server) Update(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
// 	writer.Header().Set("Content-Type", "application/json")
// 	http.Error(
// 		writer, "API Under construction", http.StatusNotImplemented)
// }

// func (self *Server) Delete(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
// 	writer.Header().Set("Content-Type", "application/json")
// 	http.Error(
// 		writer, "API Under construction", http.StatusNotImplemented)
// }

func (self *Server) Run(serverOption *Option) {
	// Build Routing
	router := httprouter.New()
	router.GET("/", self.Root)
	router.GET("/v1", self.V1Root)
	router.GET("/v1/jirou", self.Search)
	router.GET("/v1/jirou/:id", self.Read)
	// router.POST("/v1/jirou", self.Create)
	// router.PUT("/v1/jirou/:id", self.Update)
	// router.DELETE("/v1/jirou/:id", self.Delete)

	// Open db
	dbOption := opt.Options{
		ErrorIfMissing: false,
	}
	db, err := leveldb.OpenFile("./jirou.db", &dbOption)
	for {
		if err != nil {
			fmt.Println(err)
			break
		} else {
			defer db.Close()
			self.Db = db
		}

		// Start serving
		fmt.Println("Starting API server, ", serverOption)
		var port string = fmt.Sprintf(":%d", serverOption.Port)
		log.Fatal(
			http.ListenAndServe(port, router))
	}
}
