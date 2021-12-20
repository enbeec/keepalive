package main

// cribbed from https://github.com/peterbourgon/diskv/blob/master/examples/advanced-transform/advanced-transform.go

import (
	"strings"

	"github.com/peterbourgon/diskv/v3"
)

var dbBasePath string

const ERR_DBNOINIT = "Database is not initialized!"

func fullPath(key string) string {
	if dbBasePath == nil {
		panic(ERR_DBNOINIT)
	}
	return dbBasePath + "/" + key
}

func InitDB(basePath string) *diskv.Diskv {
	dbBasePath = basePath
	d := diskv.New(diskv.Options{
		BasePath:          basePath,
		AdvancedTransform: KeepaliveTransform,
		InverseTransform:  KeepaliveTransformInverse,
		CacheSizeMax:      1024 * 1024,
	})

	return d
}

func KeepaliveTransform(key string) *diskv.PathKey {
	path := strings.Split(key, "/")

	// CONDITION 1: keys must start with "keepalive/"
	if len(path) <= 0 {
		panic("can't write to empty path")
	} else if path[0] != "keepalive" {
		panic("Can't use unrelated path: " + fullPath(key))
	}

	last := len(path) - 1
	var extension string

	/*
	 * CONDITION 2: (for working with a user)
	 *  - last == 2 (keepalive/users/${username})
	 *  - second to last path member is "users"
	 *  - extension will be set to ".user"
	 */
	if last == 2 && path[last-1] == "users" {
		extension == ".user"
	}

	/*
	 * CONDITION 3: (for working with a todo)
	 *  - last == 3 (keepalive/todos/${username}/todo
	 *  - third to last path member is "todos"
	 *  - extension will be set to ".txt"
	 *  ? could check and see if second to last path member is a valid username
	 */
	if last == 3 && path[last-2] == "todos" {
		extension = ".txt"
	}

	return &diskv.PathKey{
		Path:     path[:last],
		FileName: path[last] + extension,
	}
}

func KeepaliveTransformInverse(pathKey *diskv.PathKey) (key string) {
	// condition 1
	if pathKey.Path[0] != "keepalive" {
		panic("Can't use unrelated path: " + fullPath(key))
	}

	last := len(pathKey.Path)
	var extension string

	// condition 2
	if last == 2 && pathKey.Path[last-1] == "users" {
		extension = ".user"
	}

	// condition 3
	if last == 3 && pathKey.Path[last-2] == "todos" {
		extension = ".txt"
	}

	return strings.Join(pathKey.Path, "/") + pathKey.FileName[:len(pathKey.FileName)-len(extension)]
}
