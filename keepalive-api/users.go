package main

import "github.com/1set/todotxt"

const usersPathPrefix = "keepalive/users"

type User struct {
	Username  string           `json:"username"`
	FirstName string           `json:"-"`
	LastName  string           `json:"-"`
	Token     string           `json:"-"`
	Tasks     todotxt.TaskList `json:"tasks"`
}
