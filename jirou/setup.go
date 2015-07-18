package jirou

import (
	"encoding/json"
	"fmt"
)

import (
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/opt"
)

func Setup() error {
	option := opt.Options{
		ErrorIfMissing: false,
	}

	// 店舗データを直接記述する
	// データは限定しているが、ほぼ同人誌のためだけのプログラムということで、
	// 簡単のためにこのようにしている
	var jsonBlob = []byte(`[
	  {
			"shop_name" : "ラーメン二郎 三田本店",
			"business_hours" : [{
				"open" : "08:30",
				"last" : "15:30"
			},{
			  "open" : "17:00",
			  "last" : "20:00"
			}],
			"regular_holiday" : [ "sun", "national" ],
			"address" : {
				"prefecture"  : "東京都",
				"city"        : "港区",
				"address1"    : "三田",
				"address2"    : "2-16-4",
				"postal_code" : "108-0073"
			},
			"menues" : [
				{ "name" : "ラーメン"             , "price" : 600  },
				{ "name" : "ぶたラーメン"         , "price" : 700  },
				{ "name" : "ぶたダブルラーメン"   , "price" : 800  },
				{ "name" : "大ラーメン"           , "price" : 650  },
				{ "name" : "ぶた入り大ラーメン"   , "price" : 750  },
				{ "name" : "ぶたダブル大ラーメン" , "price" : 1000 }
			],
			"created"   : "Fri Jun 12 23:03:50 JST 2015",
			"modified"  : "Fri Jun 12 23:03:50 JST 2015"
		},
		{
			"shop_name" : "ラーメン二郎 目黒店",
			"business_hours" : [{
				"open" : "12:00",
				"last" : "16:00"
			},{
			  "open" : "18:00",
			  "last" : "24:00"
			}],
			"regular_holiday" : [ "wed" ],
			"address" : {
				"prefecture"  : "東京都",
				"city"        : "目黒区",
				"address1"    : "目黒",
				"address2"    : "3-7-2",
				"postal_code" : "153-0063"
			},
			"menues" : [
				{ "name" : "小ラーメン"          , "price" : 500  },
				{ "name" : "小ラーメン豚入り"    , "price" : 600  },
				{ "name" : "小ラーメンW豚入り"   , "price" : 700  },
				{ "name" : "大ラーメン"          , "price" : 600  },
				{ "name" : "大ラーメン豚入り"    , "price" : 700  },
				{ "name" : "大ラーメンW豚入り"   , "price" : 800  }
			],
			"created"   : "Fri Jun 12 23:03:50 JST 2015",
			"modified"  : "Fri Jun 12 23:03:50 JST 2015"
		},
		{
			"shop_name" : "ラーメン二郎 仙川店",
			"business_hours" : [{
				"open" : "17:30",
				"last" : "23:30"
			}],
			"regular_holiday" : [ "sun" ],
			"address" : {
				"prefecture"  : "東京都",
				"city"        : "調布市",
				"address1"    : "仙川町",
				"address2"    : "1-10-17",
				"postal_code" : "182-0002"
			},
			"menues" : [
				{ "name" : "ラーメン"                   , "price" :  700 },
				{ "name" : "豚入りラーメン"             , "price" :  800 },
				{ "name" : "豚入りラーメン"             , "price" :  900 },
				{ "name" : "大盛ラーメン"               , "price" :  800 },
				{ "name" : "大盛豚入りラーメン"         , "price" :  900 },
				{ "name" : "大ダブルラーメン（鍋二郎）" , "price" : 1000 },
				{ "name" : "麺少なめ"                   , "price" :  700 },
				{ "name" : "生麺"                       , "price" :  100 }
			],
			"created"   : "Fri Jun 12 23:03:50 JST 2015",
			"modified"  : "Fri Jun 12 23:03:50 JST 2015"
		},
	  {
			"shop_name" : "ラーメン二郎 歌舞伎町店",
			"business_hours" : [{
				"open" : "11:30",
				"last" : "28:00"
			}],
			"regular_holiday" : [],
			"address" : {
				"prefecture"  : "東京都",
				"city"        : "新宿区",
				"address1"    : "歌舞伎町",
				"address2"    : "1-19-3",
				"postal_code" : "160-0021"
			},
			"menues" : [
				{ "name" : "普通盛"                 , "price" :  700 },
				{ "name" : "チャーシュー"           , "price" :  800 },
				{ "name" : "チャーシューダブル"     , "price" :  900 },
				{ "name" : "大盛"                   , "price" :  800 },
				{ "name" : "大盛チャーシュー"       , "price" :  900 },
				{ "name" : "大盛チャーシューダブル" , "price" : 1000 },
				{ "name" : "つけ麺普通"             , "price" :  800 },
				{ "name" : "つけ麺大盛"             , "price" :  900 },
				{ "name" : "煮玉子"                 , "price" :  100 },
				{ "name" : "メンマ"                 , "price" :  100 },
				{ "name" : "生ビール"               , "price" :  500 }
			],
			"created"   : "Fri Jun 12 23:03:50 JST 2015",
			"modified"  : "Fri Jun 12 23:03:50 JST 2015"
		},
	  {
			"shop_name" : "ラーメン二郎 品川店",
			"business_hours" : [{
				"open" : "10:00",
				"last" : "14:30"
			},{
				"open" : "17:00",
				"last" : "22:30"
			},{
				"open" : "11:00",
				"last" : "14:00"
			}],
			"regular_holiday" : [ "sun", "national" ],
			"address" : {
				"prefecture"  : "東京都",
				"city"        : "品川区",
				"address1"    : "北品川",
				"address2"    : "1-18-5",
				"postal_code" : "140-0001"
			},
			"menues" : [
				{ "name" : "普通盛 小"             , "price" :  700 },
				{ "name" : "普通盛焼豚 小ブタ"     , "price" :  800 },
				{ "name" : "普通盛焼豚大 小ダブル" , "price" :  900 },
				{ "name" : "大盛 大"               , "price" :  800 },
				{ "name" : "大盛焼豚 大ブタ"       , "price" :  900 },
				{ "name" : "大盛焼豚 大ダブル"     , "price" : 1000 },
				{ "name" : "煮玉子"                , "price" :  100 }
			],
			"created"   : "Fri Jun 12 23:03:50 JST 2015",
			"modified"  : "Fri Jun 12 23:03:50 JST 2015"
		}
	]`)

	var blob []interface{}
	err := json.Unmarshal(jsonBlob, &blob)
	if err != nil {
		fmt.Println(err)
		return err
	}

	db, err := leveldb.OpenFile("./jirou.db", &option)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer db.Close()

	// データベースにデータを保存する
	num := 0
	for _, c := range blob {
		v, _ := json.Marshal(c)
		key := []byte(fmt.Sprintf("%d", num))
		_ = db.Put(key, v, nil)
		num++
	}

	// 格納できているかどうかテスト
	//*
	for i := 0; i < num; i++ {
		key := []byte(fmt.Sprintf("%d", i))
		d, _ := db.Get(key, nil)
		fmt.Println("###", string(d), "\n")
	}
	//*/

	/*
		_ = db.Put([]byte("key1"), []byte("value1"), nil)
		_ = db.Put([]byte("key2"), []byte("value2"), nil)
		_ = db.Put([]byte("key3"), []byte("value3"), nil)

		d1, _ := db.Get([]byte("key1"), nil)
		d2, _ := db.Get([]byte("key2"), nil)
		d3, _ := db.Get([]byte("key3"), nil)

		fmt.Println(string(d1))
		fmt.Println(string(d2))
		fmt.Println(string(d3))

		db.Delete([]byte("key1"), nil)
		db.Delete([]byte("key2"), nil)
		db.Delete([]byte("key3"), nil)
	//*/

	return nil
}
