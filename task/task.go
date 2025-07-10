package task

import (
	"time"
)


//Task represents a single task in the to-do list.
// It contains fields for the task ID, title, due date, and priority.
type Task struct{
	Id int
	Title string
	DueDate time.Time
	Priority string
}

