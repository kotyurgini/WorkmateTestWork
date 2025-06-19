package storage

import "time"

type TaskInfo struct {
	ID             int        `json:"id"`
	Status         string     `json:"status"`
	Created        time.Time  `json:"created"`
	HandleDuration string     `json:"handle_duration"`
	Finished       *time.Time `json:"finished"`
	Result         string     `json:"result"`
}
