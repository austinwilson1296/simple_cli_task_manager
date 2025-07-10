package cmd

import (
	"fmt"
)

func commandRemove(cfg *Config, args ...string) error {
	if len(args) < 1 {
		return fmt.Errorf("task ID is required")
	}

	taskID := args[0]
	for i, t := range cfg.tasks {
		if fmt.Sprintf("%d", t.Id) == taskID {
			cfg.tasks = append(cfg.tasks[:i], cfg.tasks[i+1:]...)
			fmt.Printf("Task with ID %s has been removed.\n", taskID)
			return nil
		}
	}

	return fmt.Errorf("task with ID %s not found", taskID)
}