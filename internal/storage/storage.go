package storage

type Storage interface {
	CreateTask() (TaskInfo, error)
	GetTask(ID int) (TaskInfo, error)
	DeleteTask(ID int) error
	Close()
}
