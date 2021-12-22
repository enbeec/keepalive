package db

import (
	"bufio"
	"strings"

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
	allTasks := c.diskv.ReadString("keepalive/todos/" + username + "/todo")
	taskScanner := bufio.NewScanner(strings.NewReader(allTasks))

	todoList := todotxt.NewTaskList()

	for taskScanner.Scan() {
		task, err := todotxt.ParseTask(taskScanner.Text())
		if err != nil {
			return nil, err
		}
		todoList.AddTask(task)
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
