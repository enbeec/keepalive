package main

import (
	"fmt"
	"os"

	"github.com/enbeec/keepalive/keepalive-api/db"
	"github.com/peterbourgon/diskv/v3"
)

var d *diskv.Diskv = nil

func init() {
	var dbPath string

	home := os.Getenv("HOME")
	dbPath = home + "/keepalive-data"

	d = db.Connect(dbPath)
}

func main() {
	if d != nil {
		fmt.Println("Database is connected.")
	}
}
