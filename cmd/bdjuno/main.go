package main

import (
	"github.com/forbole/juno/v2/cmd"
	initcmd "github.com/forbole/juno/v2/cmd/init"
	parsecmd "github.com/forbole/juno/v2/cmd/parse"

	fixcmd "github.com/forbole/bdjuno/v2/cmd/fix"
	migratecmd "github.com/forbole/bdjuno/v2/cmd/migrate"
	parsegenesiscmd "github.com/forbole/bdjuno/v2/cmd/parse-genesis"

	"github.com/forbole/bdjuno/v2/types/config"

	"github.com/forbole/bdjuno/v2/database"
	"github.com/forbole/bdjuno/v2/modules"

	actionscmd "github.com/forbole/bdjuno/v2/cmd/actions"
	"github.com/forbole/bdjuno/v2/cmd/utils"
)

func main() {
	parseCfg := parsecmd.NewConfig().
		WithDBBuilder(database.Builder).
		WithEncodingConfigBuilder(config.MakeEncodingConfig(utils.GetBasicManagers())).
		WithRegistrar(modules.NewRegistrar(utils.GetAddressesParser()))

	cfg := cmd.NewConfig("bdjuno").
		WithParseConfig(parseCfg)

	// Run the command
	rootCmd := cmd.RootCmd(cfg.GetName())

	rootCmd.AddCommand(
		cmd.VersionCmd(),
		initcmd.InitCmd(cfg.GetInitConfig()),
		parsecmd.ParseCmd(cfg.GetParseConfig()),
		migratecmd.NewMigrateCmd(),
		fixcmd.NewFixCmd(cfg.GetParseConfig()),
		parsegenesiscmd.NewParseGenesisCmd(cfg.GetParseConfig()),
		actionscmd.NewActionsCmd(cfg.GetParseConfig()),
	)

	executor := cmd.PrepareRootCmd(cfg.GetName(), rootCmd)
	err := executor.Execute()
	if err != nil {
		panic(err)
	}
}
