package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"go-gin-api/boot"
	"go-gin-api/pkg/console"
	"go-gin-api/pkg/file"
	"strings"
)

func genCommand() *cobra.Command {
	gen := &cobra.Command{
		Use:   "gen",
		Short: "generate file and code",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			boot.Boot.Init()
		},
	}

	gen.AddCommand(
		genMigrateCommand(),
	)

	return gen
}

func generateFile(filePath string, templateStr string, variables ...map[string]string) {
	replaces := make(map[string]string)
	if len(variables) > 0 {
		replaces = variables[0]
	}

	// check whether the file exists
	if file.Exists(filePath) {
		console.Exit(filePath + " already exists!")
	}

	// variable replacement for template content
	for search, replace := range replaces {
		templateStr = strings.ReplaceAll(templateStr, search, replace)
	}

	err := file.Put([]byte(templateStr), filePath)
	console.ExitIf(err)

	console.Successp(fmt.Sprintf("[%s] created.", filePath))
}
