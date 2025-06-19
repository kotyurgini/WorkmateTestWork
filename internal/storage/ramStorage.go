package storage

import (
	"fmt"
	"sync"

	"github.com/kotyurgini/WorkmateTestWork/internal/task"
)

type RAMStorage struct {
	nextID int
	tasks  map[int]task.Task
	mu     sync.RWMutex
}

func NewRAMStorage() *RAMStorage {
	return &RAMStorage{
		nextID: 1,
		tasks:  make(map[int]task.Task),
	}
}

func (s *RAMStorage) CreateTask() (TaskInfo, error) {
	newTask := task.NewNormalTask()
	err := s.saveTaskToCache(newTask)
	if err != nil {
		return TaskInfo{}, fmt.Errorf("failed to save task to cache: %w", err)
	}
	go newTask.Do()
	return TaskDataToTaskInfo(newTask.Data()), nil
}

func (s *RAMStorage) GetTask(ID int) (TaskInfo, error) {
	t, err := s.getTaskFromCache(ID)
	if err != nil {
		return TaskInfo{}, err
	}

	return TaskDataToTaskInfo(t.Data()), nil
}

func (s *RAMStorage) DeleteTask(ID int) error {
	t, err := s.getTaskFromCache(ID)
	if err != nil {
		return err
	}
	t.Cancel()
	s.removeTaskFromCache(ID)
	return nil
}

func (s *RAMStorage) saveTaskToCache(t task.Task) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if err := t.SetID(s.nextID); err != nil {
		return err
	}
	s.tasks[s.nextID] = t
	s.nextID++
	return nil
}

func (s *RAMStorage) getTaskFromCache(ID int) (task.Task, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	t, exists := s.tasks[ID]
	if !exists {
		return nil, fmt.Errorf("task with id - %d not found", ID)
	}
	return t, nil
}

func (s *RAMStorage) removeTaskFromCache(ID int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.tasks, ID)
}

func (s *RAMStorage) Close() {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, task := range s.tasks {
		task.Cancel()
	}

	clear(s.tasks)
}
