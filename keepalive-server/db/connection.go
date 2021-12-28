package db

import (
	"os"

	"github.com/1set/todotxt"
	"github.com/peterbourgon/diskv/v3"
)

// Connection represents a connection to the database
type Connection struct {
	diskvOptions diskv.Options
	diskv        *diskv.Diskv
}

// ConnectDiskv sets up a Connection backed by diskv/v3
//		Accepts configuration functions: CacheSizeMax, TempDir
func ConnectDiskv(basePath string, opts ...func(*Connection)) *Connection {
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

// CacheSizeMax enables overwriting the CacheSizeMax
//		of a referenced Connection
func CacheSizeMax(size uint64) func(*Connection) {
	return func(c *Connection) {
		c.diskvOptions.CacheSizeMax = size
	}
}

// TempDir enables overwriting the TempDir path
//		of a referenced Connection
func TempDir(path string) func(*Connection) {
	return func(c *Connection) {
		dirMustExist(path, c.diskvOptions.PathPerm)
		sameFilesystem(path, c.diskvOptions.BasePath)
		c.diskvOptions.TempDir = path
	}
}

/*
 * Right now the "general" Read/Write methods wrap diskv-specific operations
 *		because the intention is to expand this backend to other key/value
 *		persistence options like object storage like S3
 */

// ReadTodos retrieves a TaskList for a user.
//		Accepts a username string.
func (c *Connection) ReadTodos(username string) (todotxt.TaskList, error) {
	return c.readTodosDiskv(username)
}

// WriteTodos writes a TaskList for a user.
//		Accepts a username string and a TaskList
func (c *Connection) WriteTodos(username string, todoList todotxt.TaskList) error {
	return c.writeTodosDiskv(username, todoList)
}

// ReadUser retrieves a User by reference.
//		Accepts a username string.
func (c *Connection) ReadUser(username string) (*User, error) {
	return c.readUserDiskv(username)
}

// WriteUser writes a referenced User
func (c *Connection) WriteUser(user *User) error {
	return c.writeUserDiskv(user)
}
