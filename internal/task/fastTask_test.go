package task

import (
	"testing"
	"time"
)

func TestFastTaskDo(t *testing.T) {
	task := NewFastTask()
	if task.status != "created" {
		t.Fatalf("task status should be created")
	}

	go task.Do()

	time.Sleep(1 * time.Second)
	if task.status != "in progress" {
		t.Fatalf("task status should be in progress")
	}
	time.Sleep(10 * time.Second)
	if task.status != "completed" {
		t.Fatalf("task status should be completed, got %s", task.status)
	}
}

func TestFastTaskCancel(t *testing.T) {
	task := NewFastTask()

	go task.Do()

	time.Sleep(2 * time.Second)

	task.cancelCh <- struct{}{}
	time.Sleep(1 * time.Second)
	if task.status != "cancelled" {
		t.Fatalf("task status should be cancelled, got %s", task.status)
	}
}
