package main

import (
	"fmt"
	"os"

	"github.com/enbeec/keepalive/keepalive-api/db"
)

var d *db.Connection = nil

func init() {
	home := os.Getenv("HOME")
	dbPath := home + "/keepalive-data"
	tempPath := home + "/keepalive-data__temp"

	//d = db.Connect(dbPath)
	d = db.Connect(dbPath, db.TempDir(tempPath))
}

func main() {
	if d != nil {
		fmt.Println("Database is connected.")
	}
}
