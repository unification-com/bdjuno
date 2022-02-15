package migratepartition

import (
	"fmt"

	"github.com/forbole/bdjuno/v2/database"
	"github.com/forbole/juno/v2/cmd/parse"
	"github.com/spf13/cobra"
)

func NewMigratePartitionCmd(parseCfg *parse.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use: "partition",
		RunE: func(cmd *cobra.Command, args []string) error {
			parseCtx, err := parse.GetParsingContext(parseCfg)
			if err != nil {
				return err
			}

			// Get the database
			db := database.Cast(parseCtx.Database)

			fmt.Println(db)

			return nil
		},
	}

	return cmd
}
