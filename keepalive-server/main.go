package main

import (
	"fmt"
	"os"

	"github.com/enbeec/keepalive/keepalive-server/db"
)

var d *db.Connection = nil

func init() {
	home := os.Getenv("HOME")
	// providing a tempPath to diskv allows for atomic operations
	//	all changes are written to to the temp dir and then renamed
	//	for this reason:
	//
	// **the base and temp directories must be on the same filesystem**
	tempPath := home + "/keepalive-data__temp"
	dbPath := home + "/keepalive-data"

	//d = db.Connect(dbPath)
	d = db.ConnectDiskv(dbPath, db.TempDir(tempPath))
}

func main() {
	if d != nil {
		fmt.Println("Database is connected.")
	}
}
