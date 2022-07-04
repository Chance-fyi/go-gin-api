package cmd

import (
	"github.com/spf13/cobra"
	"go-gin-api/pkg/migrate"
)

func migrateRollbackCommand() *cobra.Command {
	rollback := &cobra.Command{
		Use:   "rollback",
		Short: "rollback the last migration operation",
		Run: func(cmd *cobra.Command, args []string) {
			migrate.Run(migrate.Rollback)
		},
	}

	return rollback
}
