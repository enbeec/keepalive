package db

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/peterbourgon/diskv/v3"
)

func PathKeyToString(pk *diskv.PathKey) string {
	return strings.Join(append(pk.Path, pk.FileName), "/")
}

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
		t.Run(
			// subtest name
			fmt.Sprintf("%s:%s", name, testCase.StringPath),
			// subtest function
			func(t *testing.T) {
				gotPathKey := KeepaliveTransform(testCase.StringPath)
				if reflect.DeepEqual(
					gotPathKey, testCase.PathKey) == testCase.ShouldFail {
					t.Fatalf("%s => %s",
						testCase.StringPath, PathKeyToString(gotPathKey))
				}
			},
		)
	}
}

func TestKeepaliveTransformInverse(t *testing.T) {
	for name, testCase := range transforms {
		t.Run(
			// subtest name
			fmt.Sprintf("%s:%s", name, testCase.StringPath),
			// subtest function
			func(t *testing.T) {
				gotStringPath := KeepaliveTransformInverse(testCase.PathKey)
				if reflect.DeepEqual(
					gotStringPath, testCase.StringPath) == testCase.ShouldFail {
					t.Fatalf("%s <= %s",
						gotStringPath, PathKeyToString(testCase.PathKey))
				}
			},
		)
	}
}
