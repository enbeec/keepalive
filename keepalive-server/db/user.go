package db

import (
	"encoding/json"
	"fmt"
	"io"
)

// User represents a single user object in the database
type User struct {
	Username string `json:"username"`

	// more translateable than first and last name
	GivenName  string `json:"given_name"`
	FamilyName string `json:"family_name"`
}

// NewUser returns a reference to a blank User
func NewUser() *User {
	return &User{
		Username:   "",
		GivenName:  "",
		FamilyName: "",
	}
}

// NewUserFromJSON returns a reference to a user created from JSON
//		Inputs must be of string type or satisfy the io.Reader interface
//
// Example input:
//		{
//			"username": "steve",
//			"given_name": "steve",
//			"family_name": "crustman",
//		}
func NewUserFromJSON(fromJSON interface{}) (*User, error) {
	user := &User{
		Username:   "",
		GivenName:  "",
		FamilyName: "",
	}

	if j, isReader := fromJSON.(io.Reader); isReader {
		jsonBytes, err := io.ReadAll(j)
		if err != nil {
			return nil, err
		}
		if err := json.Unmarshal(jsonBytes, user); err != nil {
			return nil, err
		}
	} else if j, isString := fromJSON.(string); isString {
		if err := json.Unmarshal([]byte(j), user); err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("Cannot unmarshal %v", fromJSON)
	}

	return user, nil
}

func (u *User) json() (string, error) {
	json, err := json.Marshal(u)
	if err != nil {
		return "", err
	}

	return string(json), nil
}
