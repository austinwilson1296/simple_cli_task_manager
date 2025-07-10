package storage

import(
	"github.com/austinwilson1296/to-do/task"
)

type Store struct{
	Tasks []task.Task
}

func (s *Store) AddTask(t task.Task) error {
	s.Tasks = append(s.Tasks,t)
	return nil
}