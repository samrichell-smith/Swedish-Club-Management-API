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

func TestDelete(t *testing.T) {
	ts := New()
	id1 := ts.CreateTask("Foo", nil, time.Now())
	id2 := ts.CreateTask("Bar", nil, time.Now())

	if err := ts.DeleteTask(id1 + 1001); err == nil {
		t.Fatalf("delete task id=%d, got no error; want error", id1+1001)
	}

	if err := ts.DeleteTask(id1); err != nil {
		t.Fatal(err)
	}
	if err := ts.DeleteTask(id1); err == nil {
		t.Fatalf("delete task id=%d, got no error; want error", id1)
	}

	if err := ts.DeleteTask(id2); err != nil {
		t.Fatal(err)
	}
}

func TestDeleteAll(t *testing.T) {
	ts := New()
	ts.CreateTask("Foo", nil, time.Now())
	ts.CreateTask("Bar", nil, time.Now())

	if err := ts.DeleteAllTasks(); err != nil {
		t.Fatal(err)
	}

	tasks := ts.GetAllTasks()
	if len(tasks) > 0 {
		t.Fatalf("want no tasks remaining; got %v", tasks)
	}
}