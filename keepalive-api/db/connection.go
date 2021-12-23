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

func (c *Connection) ReadTodos(username string) (todotxt.TaskList, error) {
	return c.readTodosDiskv(username)
}

func (c *Connection) WriteTodos(username string, todoList todotxt.TaskList) error {
	return c.writeTodosDiskv(username, todoList)
}

// TODO: ReadUser
// TODO: WriteUser

/* TODO: ====================== diskv specific  =====================\
 * |	try recovering from a panic occuring in a transform function |
 * |	and return an appropriate error. If that works, all methods* |
 * |	must implement a deferred recovery that propogrates an error.|
 * |                                                                 |
 * |	* it is okay to omit the error if the method does not invoke |
 * |		a transform or it's inverse AND does not have another    |
 * |		error to propogate up to the method caller.              |
 * \=================================================================/
 */

func (c *Connection) readTodosDiskv(username string) (todotxt.TaskList, error) {
	allTasks, err := c.diskv.ReadStream("keepalive/todos/"+username+"/todo", false)
	defer allTasks.Close() // ReadStream returns a ReadCloser (has .Read() and .Close())
	if err != nil {
		return nil, err
	}

	todoList := todotxt.NewTaskList()
	// LoadFromFile converts the "file" to a bufio.Scanner so any io.Reader is fine
	//		however, todotxt requires *os.File specifically so we're asserting that
	err = todoList.LoadFromFile(allTasks.(*os.File))
	if err != nil {
		return nil, err
	}

	return todoList, nil
}

func (c *Connection) writeTodosDiskv(username string, todoList todotxt.TaskList) error {
	if err := c.diskv.WriteString(
		"keepalive/todos/"+username+"/todo", todoList.String(),
	); err != nil {
		return err
	}

	return nil
}
