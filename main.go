package main

import (
	"fmt"
	"os"

	"github.com/coolzilj/renrenbk/cmd"
)

func main() {
	renrenbkCmd := cmd.NewRenrenbkCmd()
	if err := renrenbkCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
