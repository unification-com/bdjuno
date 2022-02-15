package migrate

import (
	"github.com/forbole/juno/v2/cmd/parse"
	"github.com/spf13/cobra"

	migrateconfig "github.com/forbole/bdjuno/v2/cmd/migrate/config"
	migratepartition "github.com/forbole/bdjuno/v2/cmd/migrate/partition"
)

// NewMigrateCmd returns the command to be run when migrating the config from v1 to v2
func NewMigrateCmd(parseCfg *parse.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use: "migrate",
	}

	cmd.AddCommand(
		migrateconfig.NewMigrateConfigCmd(),
		migratepartition.NewMigratePartitionCmd(parseCfg),
	)

	return cmd
}
