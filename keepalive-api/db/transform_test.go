package db

import (
	"reflect"
	"testing"

	"github.com/peterbourgon/diskv/v3"
)

type transform struct {
	ShouldFail bool
	StringPath string
	PathKey    *diskv.PathKey
}

var transforms = map[string]transform{
	"user1": {
		StringPath: "keepalive/users/val",
		PathKey: &diskv.PathKey{
			Path:     []string{"keepalive", "users"},
			FileName: "val.user",
		},
	},
	"todo1": {
		StringPath: "keepalive/todos/val/todo",
		PathKey: &diskv.PathKey{
			Path:     []string{"keepalive", "todos", "val"},
			FileName: "todo.txt",
		},
	},
}

func TestKeepaliveTransform(t *testing.T) {
	for name, testCase := range transforms {

		t.Logf("%s: transform", name)
		gotStringPath := KeepaliveTransformInverse(testCase.PathKey)
		t.Logf("%s: inverse transform", name)
		gotPathKey := KeepaliveTransform(testCase.StringPath)

		if !reflect.DeepEqual(gotStringPath, testCase.StringPath) {
			t.Fail()
		}

		if !reflect.DeepEqual(gotPathKey, testCase.PathKey) {
			t.Fail()
		}

		if t.Failed() {
			t.Logf("%s transformed to %s", testCase.StringPath, gotPathKey)
			t.Logf("%s transformed to %s", testCase.PathKey, gotStringPath)
			t.FailNow()
		}
	}
}
