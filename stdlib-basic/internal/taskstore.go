package taskstore

import (
	"sync"
	"time"
)

type Task struct {
	Id   int `json:"id"`
	Text string `json:"text"`
	Tags []string `json:"tags"`
	Due  time.Time `json:"due"`
}

type TaskStore struct {
	sync.Mutex

	tasks map[int]Task
	nextId int
}



func New() *TaskStore {
	ts := &TaskStore{}
	ts.tasks = make(map[int]Task)
	ts.nextId = 0
	return ts
}

func (ts *TaskStore) CreateTask(text string, tags []string, due time.Time) int {
	ts.Lock()
	defer ts.Unlock()


	task := Task {
		Id: ts.nextId,
		Text: text,
		Due: due,
	}

	task.Tags = make([]string, len(tags))
	copy(task.Tags, tags)

	ts.tasks[ts.nextId] = task
	ts.nextId++
	return task.Id
}

func (ts *TaskStore) GetTask(id int) (Task, error)

func (ts *TaskStore) DeleteTask(id int) error

func (ts *TaskStore) DeleteAllTasks() error

func (ts *TaskStore) GetAllTasks() []Task

func (ts *TaskStore) GetTasksByTag(tag string) []Task

func (ts *TaskStore) GetTasksByDueDate(year int, month time.Month, day int) []Task
