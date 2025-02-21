package taskstore

import (
	"testing"
	"time"
)

func TestCreateAndGet(t *testing.T) {
	ts := New()
	id := ts.CreateTask("Hola", nil, time.Now())

	task, err := ts.GetTask(id)

	if err != nil {
		t.Fatal(err)
	}

	if task.Id != id {
		t.Errorf("got task.Id=%d, id=%d", task.Id, id)
	}
}