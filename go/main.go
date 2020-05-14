package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	path := flag.String("p", "/tmp/rocksdb_data", "rocksdb data path")
	flag.Parse()
	log.Println("Data path:", *path)

	//NEW RocksDB
	rocksdbCl := newRocksdb(*path)
	if rocksdbCl == nil {
		panic("Failed to newRocksdbCache.")
	}

	//INIT test data
	key := "test_1"
	value := "test_1_value"

	//SET
	err := rocksdbCl.Set(key, []byte(value))
	if err != nil {
		panic(fmt.Sprintf("Failed to RocksDB_Set; Err: %v", err))
	}

	//GET
	getValue, err := rocksdbCl.Get(key)
	if err != nil {
		panic(fmt.Sprintf("Failed to RocksDB_GET; Err: %v", err))
	}
	log.Println("get value:", string(getValue))
	if string(getValue) != value {
		panic(fmt.Sprintf("get value != set value: %v != %v", string(getValue), value))
	}

	//get stat
	stat := rocksdbCl.GetStat()
	log.Printf("stat: %v\n", stat)

	//DEL
	err = rocksdbCl.Del(key)
	if err != nil {
		panic(fmt.Sprintf("Failed to RocksDB_Del; Err: %v", err))
	}

	//Close
	rocksdbCl.CloseRocksdb()
}
