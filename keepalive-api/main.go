package main

import (
	"fmt"

	"github.com/peterbourgon/diskv/v3"
)

var db *diskv.Diskv = nil

func init() {
	db = CreateDB()
}

func main() {
	if db != nil {
		fmt.Println("Database is connected.")
	}
}
