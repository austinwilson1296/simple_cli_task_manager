package cmd

import(
	"os"
	"fmt"
)

func commandExit(cfg *Config,args ...string) error{
	fmt.Printf("Exiting the program...\n")
	os.Exit(0)
	return nil
}