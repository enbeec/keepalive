package db

import (
	"github.com/peterbourgon/diskv/v3"
)

var dbBasePath string

const ERR_DBNOINIT = "Database is not initialized!"

func fullPath(key string) string {
	if dbBasePath == "" {
		panic(ERR_DBNOINIT)
	}
	return dbBasePath + "/" + key
}

func Connect(basePath string) *diskv.Diskv {
	dbBasePath = basePath
	d := diskv.New(diskv.Options{
		BasePath:          basePath,
		AdvancedTransform: keepaliveTransform,
		InverseTransform:  keepaliveTransformInverse,
		CacheSizeMax:      1024 * 1024,
	})

	return d
}
