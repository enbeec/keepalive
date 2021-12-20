package main

import "github.com/1set/todotxt"

const usersPathPrefix = "keepalive/users"

// User is a struct that defines how a user's
//	data is to be represented to the client
type User struct {
	Username string           `json:"username"`
	FullName string           `json:"fullname"`
	Token    string           `json:"-"`
	Tasks    todotxt.TaskList `json:"tasks,omitempty"`
}

// dbUser is a struct that defines a user's
//	data is represented in the database
type dbUser struct {
	Username  string
	FirstName string
	LastName  string
	Token     string
}
