package db

import (
	"os"

	"github.com/1set/todotxt"
	"github.com/peterbourgon/diskv/v3"
)

type Connection struct {
	diskv *diskv.Diskv
}

func Connect(basePath string) *Connection {
	dv := diskv.New(diskv.Options{
		BasePath:          basePath,
		AdvancedTransform: KeepaliveTransform,
		InverseTransform:  KeepaliveTransformInverse,
		CacheSizeMax:      1024 * 1024,
	})

	return &Connection{diskv: dv}
}

/*
 * TODO:
 * 		test that a method on a connection can recover from a panic
 *		and return an appropriate error. If that works, all methods*
 *		must implement a deferred recovery that propogrates an error.
 *
 *		* it is okay to omit the error if the method does not invoke
 *			a transform or it's inverse AND does not have another
 *			error to propogate up to the method caller.
 */

func (c *Connection) ReadTodos(username string) (todotxt.TaskList, error) {
	allTasks, err := c.diskv.ReadStream("keepalive/todos/"+username+"/todo", false)
	defer allTasks.Close()
	if err != nil {
		return nil, err
	}

	todoList := todotxt.NewTaskList()
	// LoadFromFile converts the "file" to a bufio.Scanner so any io.Reader is fine
	err = todoList.LoadFromFile(allTasks.(*os.File))
	if err != nil {
		return nil, err
	}

	return todoList, nil
}

func (c *Connection) WriteTodos(username string, todoList todotxt.TaskList) error {
	err := c.diskv.WriteString(
		"keepalive/todos/"+username+"/todo",
		todoList.String(),
	)

	if err != nil {
		return err
	}

	return nil
}
