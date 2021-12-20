package main

import (
	"fmt"
	"os"

	"github.com/peterbourgon/diskv/v3"
)

var db *diskv.Diskv = nil

func init() {
	var dbPath string

	home = os.Getenv("HOME")
	dbPath = home + "/keepalive-data"

	db = InitDB(dbPath)
}

func main() {
	if db != nil {
		fmt.Println("Database is connected.")
	}
}
