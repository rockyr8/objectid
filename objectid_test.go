package objectid

import (
	"testing"
)

func TestNew(t *testing.T) {
	objectId := New()
	t.Log(objectId)
	t.Log(objectId.Machine())
}

func TestEqual(t *testing.T) {
	objectId := New()
	compare_objectId, _ := Parse(objectId.String())
	if objectId != compare_objectId {
		t.Error("two instance is not equal.")
	}
}
