package db

import (
	"github.com/peterbourgon/diskv/v3"
)

var dbBasePath string

func fullPath(key string) string {
	var prefix string
	if dbBasePath == "" {
		prefix = "UNKNOWN_ROOT"
	} else {
		prefix = dbBasePath
	}
	return prefix + "/" + key
}

type Connection struct {
	diskv *diskv.Diskv
}

func Connect(basePath string) *Connection {
	dbBasePath = basePath
	dv := diskv.New(diskv.Options{
		BasePath:          basePath,
		AdvancedTransform: KeepaliveTransform,
		InverseTransform:  KeepaliveTransformInverse,
		CacheSizeMax:      1024 * 1024,
	})

	return &Connection{diskv: dv}
}
