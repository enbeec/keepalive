package db

import (
	"fmt"
	"strings"
	"testing"
)

type testUser struct {
	ShouldFail bool
	JSON       interface{}
	User       *User
}

const mdocado_json = `{
	"family_name": "Delgado",
	"given_name": "Maurice",
	"username": "mdocado"
}`

var mdocado_struct = User{
	FamilyName: "Delgado",
	GivenName:  "Maurice",
	Username:   "mdocado",
}

var testUsers = map[string]testUser{
	"string": testUser{
		JSON: mdocado_json,
		User: &mdocado_struct,
	},
	"io.Reader": testUser{
		JSON: strings.NewReader(mdocado_json),
		User: &mdocado_struct,
	},
}

func TestNewUserFromJSON(t *testing.T) {
	for name, testCase := range testUsers {
		t.Run(
			fmt.Sprintf("%s", name),
			func(t *testing.T) {
				_, err := NewUserFromJSON(testCase.JSON)
				if err != nil != testCase.ShouldFail {
					t.Fatalf(
						"got unexpected error value => %v", err)
				}
				// nothing else to test
				// *unless* I had more complex user creation logic
				// no need to test encoding/json's heavy lifting
			},
		)
	}
}

func TestUserJSON(t *testing.T) {
	for name, testCase := range testUsers {
		t.Run(
			fmt.Sprintf("%s", name),
			func(t *testing.T) {
				_, err := testCase.User.json()
				if err != nil != testCase.ShouldFail {
					t.Fatalf(
						"got unexpected error value => %v", err)
				}
				// nothing else to test
				// *unless* I had more complex user creation logic
				// no need to test encoding/json's heavy lifting
			},
		)
	}
}
