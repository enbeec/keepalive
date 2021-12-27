package db

import (
	"os"

	"github.com/1set/todotxt"
)

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
	defer allTasks.Close()
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

func (c *Connection) readUserDiskv(username string) (*User, error) {
	jsonUser, err := c.diskv.ReadStream("keepalive/users/"+username, false)
	if err != nil {
		return nil, err
	}

	user, err := NewUserFromJSON(jsonUser)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (c *Connection) writeUserDiskv(user *User) error {
	userJSON, err := user.json()
	if err != nil {
		return err
	}

	if err := c.diskv.WriteString(
		"keepalive/users/"+user.Username, userJSON,
	); err != nil {
		return err
	}

	return nil
}
