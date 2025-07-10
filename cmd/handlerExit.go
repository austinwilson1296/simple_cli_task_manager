package cmd

import(
	"os"
	"fmt"
)
//Command to exit the program. This function will be called when the user inputs the "exit" command.
// It prints a message to the console and then exits the program with a status code of 0, indicating a successful exit.
func commandExit(cfg *Config,args ...string) error{
	fmt.Printf("Exiting the program...\n")
	os.Exit(0)
	return nil
}