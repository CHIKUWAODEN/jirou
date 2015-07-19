/*
Package jirou implements a server and utilities.
Includes:

	- sever.go  The Server
	- setup.go  Setup command
	- help.go   Print help
*/
package jirou

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

import (
	"github.com/julienschmidt/httprouter"
	"github.com/pborman/uuid"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/opt"
)

// Server class
type Server struct {
	Router   *httprouter.Router
	Db       *leveldb.DB
	DbReport *leveldb.DB
}

// Generate new Server instance
func NewServer() *Server {
	var s = new(Server)
	s.Router = httprouter.New()
	s.Db = nil
	s.DbReport = nil
	return s
}

// API function : "/v1/jirou"
// Get index of all shops.
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
			"content" : [
				%v
			]
		}`,
		content,
	))

	writer.Header().Set("Content-Type", "application/json")
	writer.Write(response)
}

// API function : "/v1/jirou"
// Get a shop data by ID
func (self *Server) Read(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	key := []byte(params.ByName("id"))
	content, _ := self.Db.Get(key, nil)

	response := []byte(fmt.Sprintf(`
		{
			"content" : %v
		}`,
		string(content),
	))

	writer.Header().Set("Content-Type", "application/json")
	writer.Write(response)
}

// API function : GET "/v1/jirou/:id/report"
func (self *Server) SearchReport(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {

	values := make([]string, 0, 5)
	iter := self.DbReport.NewIterator(nil, nil)
	for iter.Next() {
		values = append(values, string(iter.Value()))
	}
	fmt.Printf("%v\n", values)
	iter.Release()
	err := iter.Error()
	if nil != err {
		// Error
		http.Error(writer, "DB iteration error", http.StatusInternalServerError)
		return
	}

	content := strings.Join(values, ",")
	response := []byte(fmt.Sprintf(`
		{
			"content" : [
				%v
			]
		}`,
		content,
	))

	writer.Header().Set("Content-Type", "application/json")
	writer.Write(response)
}

/*
API function : POST "/v1/jirou/:id/report"

.reporter (string)
  Reporter name

.comment (string)
	Free word comment

.rating.noodle (float)
	麺の評価レートで、0.0 以上 5.0 未満の数値を入力する.
	0.0 を下回る値は 0.0 に丸められます
	5.0 を上回る値は 5.0 に丸められます

.rating.soup (float)
	スープの評価レートで、0.0 以上 5.0 未満の数値を入力する.
	0.0 を下回る値は 0.0 に丸められます
	5.0 を上回る値は 5.0 に丸められます

Example:

	{
		"reporter": "Jhon Smith",
		"comment": "lorem ipsum dolor sit amet ... (Free comment)",
		"rating": {
			"noodle": 5.0,
			"soup": 5.0
		}
	}
*/

// レポの投稿データ構造に含まれるレーティングの構造
type Rating struct {
	Noodle float32 `json:"noodle"`
	Soup   float32 `json:"soup"`
}

// レポの投稿データ構造
type Report struct {
	Reporter string `json:"reporter"`
	Comment  string `json:"comment"`
	Rating   Rating `json:"rating"`
}

// データベース保存用の Report
type ReportDb struct {
	Report
	Date   string `json:"date"`
	Uuid   string `json:"uuid"`
	ShopID string `json:"shop_id"`
}

// 以上、レポっす
func (self *Server) PostReport(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	// Parse POST data
	var report ReportDb
	decorder := json.NewDecoder(request.Body)
	decodeErr := decorder.Decode(&report)
	if decodeErr != nil {
		// Error
		http.Error(writer, "Decode Error", http.StatusInternalServerError)
		return
	}

	// Report の情報を付け足す
	report.Date = time.Now().Format(time.RFC1123Z)
	report.Uuid = uuid.New()
	report.ShopID = params.ByName("id")

	// marshal
	marshaled, marshalErr := json.Marshal(report)
	if marshalErr != nil {
		http.Error(writer, "Marshal Error", http.StatusInternalServerError)
		return
	}

	// データベースに書き込む
	dbPutErr := self.DbReport.Put([]byte(report.Uuid), marshaled, nil)
	if dbPutErr != nil {
		http.Error(writer, "DB Error", http.StatusInternalServerError)
		return
	}

	response := []byte(fmt.Sprintf(`
		{
			"content" : %s
		}
		`,
		string(marshaled),
	))
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(response)
}

// Run up The server
func (self *Server) Run(serverOption *Option) {

	// Build Routing
	router := httprouter.New()
	router.GET("/v1/jirou", self.Search)
	router.GET("/v1/jirou/:id", self.Read)
	router.GET("/v1/jirou/:id/report", self.SearchReport)
	router.POST("/v1/jirou/:id/report", self.PostReport)

	// Open db
	dbOption := opt.Options{
		ErrorIfMissing: false,
	}
	// TODO Not DRY
	db, err := leveldb.OpenFile("./jirou.db", &dbOption)
	dbReport, errDbReport := leveldb.OpenFile("./report.db", &dbOption)
	for {
		if err != nil {
			fmt.Println(err)
			break
		} else {
			defer db.Close()
			self.Db = db
		}

		if errDbReport != nil {
			fmt.Println(errDbReport)
			break
		} else {
			defer dbReport.Close()
			self.DbReport = dbReport
		}

		// Start serving
		fmt.Println("Starting API server, ", serverOption)
		var port string = fmt.Sprintf(":%d", serverOption.Port)
		log.Fatal(
			http.ListenAndServe(port, router))
	}
}
