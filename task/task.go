package task

import (
	"time"
)



type Task struct{
	Id int
	Title string
	DueDate time.Time
	Priority string
}

