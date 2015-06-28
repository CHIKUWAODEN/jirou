package jirou

import (
	"fmt"
)

import (
	"github.com/syndtr/goleveldb/leveldb"
)

func BulkInsert(filepath string) {
	fmt.Println("BulkInsert()")
	db, err := leveldb.OpenFile(filepath)
	defer db.Close()
}
