package cmd

import(
	"fmt"
)

func commandList(cfg *Config, args ...string) error {
	if len(cfg.tasks) == 0 {
		fmt.Println("No tasks available.")
		return nil
	}

	fmt.Println("Current Tasks:")
	for i, t := range cfg.tasks {
		fmt.Printf("%d: ID: %d Title: %s (Due: %s, Priority: %s)\n", i+1,t.Id, t.Title, t.DueDate.Format("2006-01-02"), t.Priority)
	}
	return nil
}