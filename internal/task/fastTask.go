package task

import (
	"fmt"
	"log"
	"time"
)

// Using for test
type FastTask struct {
	id          int
	status      StatusType
	created     time.Time
	startHandle time.Time
	finished    time.Time
	result      string
	cancelCh    chan struct{}
}

func NewFastTask() *FastTask {
	return &FastTask{
		status:   "created",
		created:  time.Now(),
		cancelCh: make(chan struct{}),
	}
}

func (t *FastTask) SetID(ID int) error {
	if t.id != 0 {
		return fmt.Errorf("task ID is already set to %d", t.id)
	}
	t.id = ID
	return nil
}

func (t *FastTask) Do() {
	if t.status != Created {
		return
	}

	delay := r.Intn(5) + 5

	log.Printf("Task %d started", t.id)
	t.startHandle = time.Now()
	t.status = "in progress"
	timer := time.NewTimer(time.Duration(delay) * time.Second)

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

func (t *FastTask) Data() TaskData {
	return TaskData{
		ID:          t.id,
		Status:      t.status,
		Created:     t.created,
		StartHandle: t.startHandle,
		Finished:    t.finished,
		Result:      t.result,
	}
}

func (t *FastTask) Cancel() {
	select {
	case t.cancelCh <- struct{}{}:
	default:
	}
}
