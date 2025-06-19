package task

import (
	"fmt"
	"log"
	"time"
)

// Using for test
type NormalTask struct {
	id          int
	status      StatusType
	created     time.Time
	startHandle time.Time
	finished    time.Time
	result      string
	cancelCh    chan struct{}
}

func NewNormalTask() *NormalTask {
	return &NormalTask{
		status:   Created,
		created:  time.Now(),
		cancelCh: make(chan struct{}),
	}
}

func (t *NormalTask) SetID(ID int) error {
	if t.id != 0 {
		return fmt.Errorf("task ID is already set to %d", t.id)
	}
	t.id = ID
	return nil
}

func (t *NormalTask) Do() {
	if t.status != Created {
		return
	}

	delay := r.Intn(2) + 3

	log.Printf("Task %d started", t.id)
	t.status = "in progress"
	t.startHandle = time.Now()
	timer := time.NewTimer(time.Duration(delay) * time.Minute)

	select {
	case <-t.cancelCh:
		log.Printf("Task %d canceled", t.id)
		t.status = "cancelled"
		timer.Stop()
	case <-timer.C:
		log.Printf("Task %d completed", t.id)
		t.status = "completed"
		t.result = "done"
	}
	t.finished = time.Now()
}

func (t *NormalTask) Data() TaskData {
	return TaskData{
		ID:          t.id,
		Status:      t.status,
		Created:     t.created,
		StartHandle: t.startHandle,
		Finished:    t.finished,
		Result:      t.result,
	}
}

func (t *NormalTask) Cancel() {
	select {
	case t.cancelCh <- struct{}{}:
	default:
	}
}
