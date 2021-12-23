package user

import (
	"github.com/1set/todotxt"
	"github.com/enbeec/keepalive/keepalive-api/db"
)

const usersPathPrefix = "keepalive/users"

// User is a struct that defines how a user's
//	data is to be represented to the client
type User struct {
	Username string           `json:"username"`
	Token    string           `json:"-"`
	Tasks    todotxt.TaskList `json:"tasks,omitempty"`

	// more translateable than first and last name
	GivenName  string `json:"givenName"`
	FamilyName string `json:"familyName"`
}

// NewUser returns a reference to a blank User with an empty TaskList
func NewUser() *User {
	return &User{
		Tasks:      todotxt.NewTaskList(),
		Username:   "",
		GivenName:  "",
		FamilyName: "",
		Token:      "",
	}
}

// LoadUser returns a fully initialized user from the database
func LoadUser(db db.Connection, username string) (*User, error) {
	tasks, err := db.ReadTodos(username)
	if err != nil {
		return nil, err
	}

	return &User{
		Username: username,
		Tasks:    tasks,
		// TODO: add the other fields from the db
		//			(blocked on db.Connection method)
	}, nil
}

// Save persists a user using a given database connection
func (u *User) Save(db db.Connection) error {
	if err := db.WriteTodos(u.Username); err != nil {
		return err
	}

	// TODO: add the non-task fields from the db
	//			(blocked on db.Connection method)
	return nil
}

// HasTasks is a helper method that returns true if the user has tasks
func (u *User) HasTasks() bool {
	return len(u.Tasks) > 0
}
