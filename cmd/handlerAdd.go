package cmd

import (
	"flag"
	"fmt"
	"strings"
	"time"
	"github.com/austinwilson1296/to-do/task"
	"github.com/google/shlex"
)
// Command to add a new task to the in-memory task list. 
// Requires flags for title (-n), due date (-dd), and priority (-p).
func commandAdd(cfg *Config, args ...string) error {
	
	argsToString := strings.Join(args, " ")
	args, err := shlex.Split(argsToString)
	if err != nil {
		return fmt.Errorf("error parsing command: %v", err)
	}

	addCmd := flag.NewFlagSet("add", flag.ContinueOnError)
	title := addCmd.String("n", "", "Title of the task")
	dueDate := addCmd.String("dd", "", "Due date in YYYY-MM-DD format")
	priority := addCmd.String("p", "", "Priority of the task (low, medium, high)")

	err = addCmd.Parse(args)
	if err != nil {
		return err
	}

	if *title == "" || *dueDate == "" || *priority == "" {
		return fmt.Errorf("all flags -n, -dd, and -p are required")
	}

	validPriorities := map[string]bool{"low": true, "medium": true, "high": true}
	if !validPriorities[*priority] {
		return fmt.Errorf("invalid priority: must be 'low', 'medium', or 'high'")
	}

	parsedDueDate, err := time.Parse("2006-01-02", *dueDate)
	if err != nil {
		return fmt.Errorf("invalid due date format, please use YYYY-MM-DD")
	}

	newTask := task.Task{
		Id: 	 len(cfg.tasks) + 1, 
		Title:    *title,
		DueDate:  parsedDueDate,
		Priority: *priority,
	}

	cfg.tasks = append(cfg.tasks, newTask)
	fmt.Printf("Task added: %s\n", newTask.Title)
	return nil
}


