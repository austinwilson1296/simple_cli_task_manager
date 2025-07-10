package main

import (
	"github.com/austinwilson1296/to-do/cmd"

)

func main(){
	cfg := &cmd.Config{}
	cmd.Run(cfg)
}