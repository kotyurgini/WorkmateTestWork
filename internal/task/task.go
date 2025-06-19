package task

import (
	"math/rand"
	"time"
)

type Task interface {
	Do()
	Data() TaskData
	Cancel()
	SetID(ID int) error
}

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

type TaskData struct {
	ID          int
	Status      StatusType
	Created     time.Time
	StartHandle time.Time
	Finished    time.Time
	Result      string
}

type StatusType string

const (
	Created    StatusType = "created"
	InProgress StatusType = "in progress"
	Completed  StatusType = "completed"
	Cancelled  StatusType = "cancelled"
)
