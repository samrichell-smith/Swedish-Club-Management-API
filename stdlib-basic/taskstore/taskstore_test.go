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

	if task.Text != "Hola" {
		t.Errorf("got Text=%v, want %v", task.Text, "Hola")
	}

	allTasks := ts.GetAllTasks()
	if len(allTasks) != 1 || allTasks[0].Id != id {
		t.Errorf("got len(allTasks)=%d, allTasks[0].Id=%d; want 1, %d", len(allTasks), allTasks[0].Id, id)
	}

	_, err = ts.GetTask(id + 1)
	if err == nil {
		t.Fatal("got nil, want error")
	}

	ts.CreateTask("hey", nil, time.Now())
	allTasks2 := ts.GetAllTasks()
	if len(allTasks2) != 2 {
		t.Errorf("got len(allTasks2)=%d, want 2", len(allTasks2))
	}

}