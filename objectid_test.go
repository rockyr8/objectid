package objectid

import (
	"testing"
)

func TestNew(t *testing.T) {
	objectId := New()
	if objectId == nil {
		t.Error("cannot generiate objectid instance.")
	}
	t.Log(objectId)
	t.Log(objectId.CreationTime())
	t.Log(objectId.Machine())
}

func TestEqual(t *testing.T) {
	objectId := New()
	compare_objectId := Parse(objectId.String())
	if objectId.Equal(compare_objectId) == false {
		t.Error("two instance is not equal.")
	}
}
