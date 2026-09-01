package task

import (
	"testing"
)

func TestNew(t *testing.T) {
	got, err := New("Learn Go")
	if err != nil {
		t.Fatalf("New() returned an unexpected error: %v", err)
	}

	if got.Title != "Learn Go" {
		t.Errorf("Title = %q, want %q", got.Title, "Learn Go")
	}

	if got.Done {
		t.Errorf("Done = true, want false")
	}
}

func TestNewRejectsEmptyTitle(t *testing.T) {
	_, err := New("   ")
	if err == nil {
		t.Fatal("New() expected an error for an empty title")
	}
}

func TestComplete(t *testing.T) {
	task, err := New("Learn Go")
	if err != nil {
		t.Fatalf("New() returned an unexpected error: %v", err)
	}

	task.Complete()

	if !task.Done {
		t.Error("Done = false, want true")
	}

	if task.Status() != "done" {
		t.Errorf("Status() = %q, want %q", task.Status(), "done")
	}
}

func TestListAdd(t *testing.T) {
	var list List

	err := list.Add("Learn Go")
	if err != nil {
		t.Fatalf("Add() returned an unexpected error: %v", err)
	}

	tasks := list.All()

	if len(tasks) != 1 {
		t.Fatalf("len(tasks) = %d, want 1", len(tasks))
	}

	if tasks[0].Title != "Learn Go" {
		t.Errorf("tasks[0].Title = %q, want %q", tasks[0].Title, "Learn Go")
	}

	if tasks[0].ID != 1 {
		t.Errorf("tasks[0].ID = %d, want 1", tasks[0].ID)
	}
}

func TestListAddRejectsEmptyTitle(t *testing.T) {
	var list List

	err := list.Add("   ")
	if err == nil {
		t.Fatal("Add() expected an error for an empty title")
	}

	if len(list.All()) != 0 {
		t.Errorf("len(list.All()) = %d, want 0", len(list.All()))
	}
}

func TestListComplete(t *testing.T) {
	var list List

	if err := list.Add("Learn Go"); err != nil {
		t.Fatalf("Add() returned an unexpected error: %v", err)
	}

	if err := list.Add("Build CLI app"); err != nil {
		t.Fatalf("Add() returned an unexpected error: %v", err)
	}

	ok := list.Complete(2)
	if !ok {
		t.Fatal("Complete() = false, want true")
	}

	tasks := list.All()

	if !tasks[1].Done {
		t.Error("tasks[1].Done = false, want true")
	}
}

func TestListCompleteUnknownID(t *testing.T) {
	var list List

	if err := list.Add("Learn Go"); err != nil {
		t.Fatalf("Add() returned an unexpected error: %v", err)
	}

	ok := list.Complete(999)

	if ok {
		t.Error("Complete() = true, want false")
	}
}

func TestListAddUsesNextAvailableID(t *testing.T) {
	var list List

	if err := list.Add("Task A"); err != nil {
		t.Fatalf("Add() returned an unexpected error: %v", err)
	}

	if err := list.Add("Task B"); err != nil {
		t.Fatalf("Add() returned an unexpected error: %v", err)
	}

	if err := list.Add("Task C"); err != nil {
		t.Fatalf("Add() returned an unexpected error: %v", err)
	}

	if !list.Delete(2) {
		t.Fatal("Delete() = false, want true")
	}

	if err := list.Add("Task D"); err != nil {
		t.Fatalf("Add() returned an unexpected error: %v", err)
	}

	tasks := list.All()

	if tasks[2].ID != 4 {
		t.Errorf("tasks[2].ID = %d, want 4", tasks[2].ID)
	}
}

func TestListDelete(t *testing.T) {
	var list List

	if err := list.Add("Task A"); err != nil {
		t.Fatalf("Add() returned an unexpected error: %v", err)
	}

	if err := list.Add("Task B"); err != nil {
		t.Fatalf("Add() returned an unexpected error: %v", err)
	}

	if !list.Delete(1) {
		t.Fatal("Delete() = false, want true")
	}

	tasks := list.All()

	if len(tasks) != 1 {
		t.Fatalf("len(tasks) = %d, want 1", len(tasks))
	}

	if tasks[0].Title != "Task B" {
		t.Errorf("tasks[0].Title = %q, want %q", tasks[0].Title, "Task B")
	}
}

func TestListSaveAndLoad(t *testing.T) {
	var list List

	if err := list.Add("Learn Go"); err != nil {
		t.Fatalf("Add() returned an unexpected error: %v", err)
	}

	if err := list.Add("Build CLI"); err != nil {
		t.Fatalf("Add() returned an unexpected error: %v", err)
	}

	filename := t.TempDir() + "/tasks.json"

	if err := list.Save(filename); err != nil {
		t.Fatalf("Save() returned an unexpected error: %v", err)
	}

	var loaded List

	if err := loaded.Load(filename); err != nil {
		t.Fatalf("Load() returned an unexpected error: %v", err)
	}

	tasks := loaded.All()

	if len(tasks) != 2 {
		t.Fatalf("len(tasks) = %d, want 2", len(tasks))
	}

	if tasks[0].Title != "Learn Go" {
		t.Errorf("tasks[0].Title = %q, want %q", tasks[0].Title, "Learn Go")
	}

	if tasks[1].Title != "Build CLI" {
		t.Errorf("tasks[1].Title = %q, want %q", tasks[1].Title, "Build CLI")
	}
}

func TestListAllReturnsCopy(t *testing.T) {
	var list List

	if err := list.Add("learn go fast"); err != nil {
		t.Fatalf("add() returned an unexpected error: %v", err)
	}

	tasks := list.All()

	tasks[0].Title = "changed title"

	original := list.All()

	if original[0].Title == tasks[0].Title {
		t.Fatalf("original task title %s, want %q", original[0].Title, "Learn go fast")
	}
}

func TestListFind(t *testing.T) {
	var list List

	if err := list.Add("we learning go 1"); err != nil {
		t.Fatalf("Add() returned an unexpected error: %v", err)
	}

	if err := list.Add("we learning go 2"); err != nil {
		t.Fatalf("Add() returned an unexpected error: %v", err)
	}

	task, ok := list.Find(2)

	if !ok {
		t.Fatalf("Find() equals to false, want true")
	}

	if task.Title != "we learning go 2" {
		t.Fatalf("task.Title = %q want %q", task.Title, "we learning go 2")
	}
}

func TestFindUnknownID(t *testing.T) {
	var list List
	_, ok := list.Find(2)
	if ok {
		t.Fatalf("Find() equals true, want false")
	}
}
