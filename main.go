package main

import (
	"fmt"
	"go-gin-api/app/cmd"
	"os"
)

func main() {
	c := cmd.RootCommand()

	if err := c.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
