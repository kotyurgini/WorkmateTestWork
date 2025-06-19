package storage

import (
	"time"

	"github.com/kotyurgini/WorkmateTestWork/internal/task"
)

func TaskDataToTaskInfo(data task.TaskData) TaskInfo {
	ti := TaskInfo{
		ID:      data.ID,
		Status:  string(data.Status),
		Created: data.Created,
		Result:  data.Result,
	}

	if data.StartHandle.IsZero() {
		ti.HandleDuration = "0s"
	} else {
		if data.Finished.IsZero() {
			ti.HandleDuration = time.Since(data.StartHandle).String()
		} else {
			ti.HandleDuration = data.Finished.Sub(data.StartHandle).String()
		}
	}

	if data.Finished.IsZero() {
		ti.Finished = nil
	} else {
		ti.Finished = &data.Finished
	}

	return ti
}
