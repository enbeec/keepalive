package db

import (
	"os"

	"github.com/1set/todotxt"
	"github.com/peterbourgon/diskv/v3"
)

type Connection struct {
	diskvOptions diskv.Options
	diskv        *diskv.Diskv
}

func Connect(basePath string, opts ...func(*Connection)) *Connection {
	defaultOptions := diskv.Options{
		BasePath:          basePath,
		AdvancedTransform: KeepaliveTransform,
		InverseTransform:  KeepaliveTransformInverse,
		CacheSizeMax:      1024 * 1024,
		PathPerm:          os.FileMode(int(0770)),
	}

	dirMustExist(defaultOptions.BasePath, defaultOptions.PathPerm)

	// prepare dummy connection for configuration
	c := &Connection{diskv: nil, diskvOptions: defaultOptions}

	// configure the connection options (if any were passed)
	for _, opt := range opts {
		opt(c)
	}

	// finally, use the dummy connection's options to connect
	c.diskv = diskv.New(c.diskvOptions)
	return c
}

func CacheSizeMax(size uint64) func(*Connection) {
	return func(c *Connection) {
		c.diskvOptions.CacheSizeMax = size
	}
}

func TempDir(path string) func(*Connection) {
	return func(c *Connection) {
		dirMustExist(path, c.diskvOptions.PathPerm)
		sameFilesystem(path, c.diskvOptions.BasePath)
		c.diskvOptions.TempDir = path
	}
}

// Right now the "general" Read/Write methods wrap diskv-specific operations
//		because the intention is to expand this backend to other key/value
//		persistence options like object storage (e.g. S3) etc.

func (c *Connection) ReadTodos(username string) (todotxt.TaskList, error) {
	return c.readTodosDiskv(username)
}

func (c *Connection) WriteTodos(username string, todoList todotxt.TaskList) error {
	return c.writeTodosDiskv(username, todoList)
}

func (c *Connection) ReadUser(username string) (*User, error) {
	return c.readUserDiskv(username)
}

func (c *Connection) WriteUser(user *User) error {
	return c.writeUserDiskv(user)
}
