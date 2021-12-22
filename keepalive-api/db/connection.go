package db

import (
	"github.com/peterbourgon/diskv/v3"
)

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
