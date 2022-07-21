package main

import (
	"go-gin-api/app/cmd"
	"go-gin-api/pkg/console"
)

func main() {
	c := cmd.RootCommand()
	err := c.Execute()
	console.ExitIf(err)
}
