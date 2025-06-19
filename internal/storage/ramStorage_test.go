package storage

import (
	"testing"
)

func TestCreateTask(t *testing.T) {
	storage := NewRAMStorage()
	newTask, err := storage.CreateTask()
	if err != nil {
		t.Fatalf("failed to create task: %v", err)
	}
	if newTask.ID == 0 {
		t.Error("task ID should not be zero")
	}
	if newTask.Status != "created" {
		t.Errorf("expected task status 'created', got '%s'", newTask.Status)
	}
}

func TestGetTask(t *testing.T) {
	storage := NewRAMStorage()
	newTask, err := storage.CreateTask()
	if err != nil {
		t.Fatalf("failed to create task: %v", err)
	}

	ti, err := storage.GetTask(newTask.ID)
	if err != nil {
		t.Fatalf("failed to get task: %v", err)
	}

	if ti.ID != newTask.ID {
		t.Errorf("expected task ID %d, got %d", newTask.ID, ti.ID)
	}

	ti, err = storage.GetTask(99)
	if err == nil {
		t.Error("expected error when getting non-existent task, got nil")
	}
}

func TestDeleteTask(t *testing.T) {
	storage := NewRAMStorage()
	newTask, err := storage.CreateTask()
	if err != nil {
		t.Fatalf("failed to create task: %v", err)
	}

	err = storage.DeleteTask(newTask.ID)
	if err != nil {
		t.Fatalf("failed to delete task: %v", err)
	}

	err = storage.DeleteTask(99)
	if err == nil {
		t.Fatalf("expected error when delete non-existent task, got nil")
	}
}
