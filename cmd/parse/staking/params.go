package staking

import (
	"fmt"
	"github.com/forbole/bdjuno/v4/database"
	"github.com/forbole/bdjuno/v4/modules/staking"
	modulestypes "github.com/forbole/bdjuno/v4/modules/types"
	parsecmdtypes "github.com/forbole/juno/v4/cmd/parse/types"
	"github.com/forbole/juno/v4/types/config"
	"github.com/spf13/cobra"
)

// paramsCmd returns a Cobra command that allows to fix the staking params.
func paramsCmd(parseConfig *parsecmdtypes.Config) *cobra.Command {
	return &cobra.Command{
		Use:   "params",
		Short: "Fix the information about staking params from the latest known height",
		RunE: func(cmd *cobra.Command, args []string) error {
			parseCtx, err := parsecmdtypes.GetParserContext(config.Cfg, parseConfig)
			if err != nil {
				return err
			}

			sources, err := modulestypes.BuildSources(config.Cfg.Node, parseCtx.EncodingConfig)
			if err != nil {
				return err
			}

			// Get the database
			db := database.Cast(parseCtx.Database)

			// Build the staking module
			stakingModule := staking.NewModule(sources.StakingSource, parseCtx.EncodingConfig.Codec, db)

			// Get latest height
			height, err := parseCtx.Node.LatestHeight()
			if err != nil {
				return fmt.Errorf("error while getting latest block height: %s", err)
			}

			// refresh params
			err = stakingModule.UpdateParams(height)
			if err != nil {
				return fmt.Errorf("error while updating params: %s", err)
			}

			return nil
		},
	}
}
