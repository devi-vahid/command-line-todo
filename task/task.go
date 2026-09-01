package task

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

type List struct {
	tasks []Task
}

type Task struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Done      bool      `json:"done"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func New(title string) (Task, error) {
	title = strings.TrimSpace(title)

	if title == "" {
		return Task{}, fmt.Errorf("task title cannot be empty")
	}
	now := time.Now()
	return Task{
		Title:     title,
		CreatedAt: now,
		UpdatedAt: now,
	}, nil
}

func (l *List) Add(title string) error {
	task, err := New(title)
	if err != nil {
		return err
	}
	maxID := 0
	for _, task := range l.tasks {
		if task.ID > maxID {
			maxID = task.ID
		}
	}

	task.ID = maxID + 1
	l.tasks = append(l.tasks, task)

	return nil
}

func (l *List) Complete(id int) bool {
	for i := range l.tasks {
		if l.tasks[i].ID == id {
			l.tasks[i].Complete()
			return true
		}
	}
	return false
}

func (l List) All() []Task {
	tasks := make([]Task, len(l.tasks))
	copy(tasks, l.tasks)
	return tasks
}

func (t Task) Status() string {
	if t.Done {
		return "done"
	}

	return "pending"
}

func (t *Task) Complete() {
	t.Done = true
	t.UpdatedAt = time.Now()
}

func (l *List) Delete(id int) bool {
	for i := range l.tasks {
		if l.tasks[i].ID == id {
			l.tasks = append(l.tasks[:i], l.tasks[i+1:]...)
			return true
		}
	}
	return false
}

func (l List) Save(filename string) error {
	data, err := json.MarshalIndent(l.tasks, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal tasks: %w", err)
	}

	if err := os.WriteFile(filename, data, 0o644); err != nil {
		return fmt.Errorf("write tasks file: %w", err)
	}

	return nil
}

func (l *List) Load(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}

		return fmt.Errorf("read tasks file: %w", err)
	}

	if err := json.Unmarshal(data, &l.tasks); err != nil {
		return fmt.Errorf("unmarshal tasks: %w", err)
	}

	return nil
}

func (l List) Find(id int) (Task, bool) {
	for i := range l.tasks {
		if l.tasks[i].ID == id {
			return l.tasks[i], true
		}
	}

	return Task{}, false
}

func (l *List) UpdateTitle(id int, title string) (Task, error){
	task, ok := l.Find(id)
	if !ok {
		return Task{}, fmt.Errorf("task-id: %d, not found", id)
	}

	title = strings.TrimSpace(title)
	if title == "" {
		return Task{}, fmt.Errorf("title can't be empty"")
	}

	task.Title = title

	return task, nil
}
