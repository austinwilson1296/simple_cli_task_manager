package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"github.com/austinwilson1296/to-do/task"
)

type Command struct {
	Name string
	handler func(*Config,...string) error
	help string
}
type Config struct {
	tasks []task.Task
}

func Run(cfg *Config) {
	
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Enter command: ")
		scanner.Scan()
		
		userInput := cleanInput(scanner.Text())

		commandName := userInput[0]
		args := []string{}
		if len(userInput) > 1 {
			args = userInput[1:]
		}
		command, exists := getCommands()[commandName]
		if exists {
			err := command.handler(cfg,args...)
			if err != nil {
				fmt.Printf("Error executing command '%s': %v\n", commandName, err)
			}
			continue
		}else{
			fmt.Printf("Command '%s' not found. Available commands: %v\n", commandName, getCommands())
		}

	

	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	fmt.Printf("%s\n", words)
	return words
}

func getCommands() map[string]Command {
	return map[string]Command{
		"exit": {
			Name: "exit",
			handler: commandExit,
			help: "usage : exit",
		},
		"add": {
			Name: "add",
			handler: commandAdd,
			help:"usage : add -n <name> -dd <dueDate> -p <priority>",
		},
		"list":{
			Name: "list",
			handler: commandList,
			help: "usage : list",
		},
		"remove": {
			Name: "remove",
			handler: commandRemove,
			help: "usage : remove <taskID>",
		},
	}
}
