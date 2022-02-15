package migrateconfig

import (
	"fmt"
	"io/ioutil"

	v1 "github.com/forbole/bdjuno/v2/cmd/migrate/v1"
	"github.com/forbole/bdjuno/v2/modules/pricefeed"
	"github.com/forbole/bdjuno/v2/types"
	junomigrate "github.com/forbole/juno/v2/cmd/migrate"
	junov1 "github.com/forbole/juno/v2/cmd/migrate/v1"
	"github.com/forbole/juno/v2/types/config"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

type Config struct {
	junomigrate.Config `yaml:",inline"`
	PricefeedConfig    *pricefeed.Config `yaml:"pricefeed"`
}

func NewMigrateConfigCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "config",
		RunE: func(cmd *cobra.Command, args []string) error {
			v2Config, err := migrateConfig()
			if err != nil {
				return nil
			}

			bz, err := yaml.Marshal(&v2Config)
			if err != nil {
				return fmt.Errorf("error while serializing v2 config: %s", err)
			}

			v2File := config.GetConfigFilePath()
			return ioutil.WriteFile(v2File, bz, 0600)
		},
	}

	return cmd
}

func migrateConfig() (Config, error) {
	bz, err := junov1.ReadConfig()
	if err != nil {
		return Config{}, err
	}

	v1JunoCfg, err := junomigrate.MigrateConfig()
	if err != nil {
		return Config{}, err
	}

	v1BDJunoCfg, err := v1.ParseConfig(bz)
	if err != nil {
		return Config{}, err
	}

	var pricefeedConfig *pricefeed.Config
	if v1BDJunoCfg.PricefeedConfig != nil {
		var tokens = make([]types.Token, len(v1BDJunoCfg.PricefeedConfig.Tokens))
		for tokenIndex, token := range v1BDJunoCfg.PricefeedConfig.Tokens {
			var units = make([]types.TokenUnit, len(token.Units))
			for unitIndex, unit := range token.Units {
				units[unitIndex] = types.NewTokenUnit(unit.Denom, unit.Exponent, unit.Aliases, unit.PriceID)
			}

			tokens[tokenIndex] = types.NewToken(token.Name, units)
		}

		pricefeedConfig = pricefeed.NewConfig(tokens)
	}

	return Config{
		Config:          v1JunoCfg,
		PricefeedConfig: pricefeedConfig,
	}, nil
}
