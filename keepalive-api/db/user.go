package db

import (
	"encoding/json"
	"fmt"
)

// User contains the name and username of a user
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

// NewUserFromJSON returns a reference to a blank User
func NewUserFromJSON(fromJSON interface{}) (*User, error) {
	user := &User{
		Username:   "",
		GivenName:  "",
		FamilyName: "",
	}

	if j, isString := fromJSON.(string); isString {
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
