package db

import (
	"strings"

	"github.com/peterbourgon/diskv/v3"
)

const USER_EXT = ".user"
const USER_EXT_LEN = len(USER_EXT)
const TODO_EXT = ".txt"
const TODO_EXT_LEN = len(TODO_EXT)

func KeepaliveTransform(key string) *diskv.PathKey {
	path := strings.Split(key, "/")

	// initial checks
	if len(path) == 0 {
		panic("can't write to empty path")
	} else if path[0] != "keepalive" {
		panic("Can't use unrelated path: " + key)
	}

	last := len(path) - 1
	var fileName string

	// set user extension
	if last == 2 && path[last-1] == "users" {
		fileName = path[last] + USER_EXT
	}

	// set todo extension
	if last == 3 && path[last-2] == "todos" {
		fileName = path[last] + TODO_EXT
	}

	return &diskv.PathKey{
		Path:     path[:last],
		FileName: fileName,
	}
}

func KeepaliveTransformInverse(pathKey *diskv.PathKey) (key string) {
	// initial checks
	if pathKey.Path[0] != "keepalive" {
		panic("Can't use unrelated path: " + key)
	}

	last := len(pathKey.Path)
	var fileNameLength int

	// handle user extension
	if last == 2 && pathKey.Path[last-1] == "users" {
		fileNameLength = len(pathKey.FileName) - USER_EXT_LEN
	}

	// handle todo extension
	if last == 3 && pathKey.Path[last-2] == "todos" {
		fileNameLength = len(pathKey.FileName) - TODO_EXT_LEN
	}

	return strings.Join(pathKey.Path, "/") + "/" + pathKey.FileName[:fileNameLength]
}
