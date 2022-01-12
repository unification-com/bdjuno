package actions

import (
	"encoding/json"
	"fmt"
	"net/http"

	sdk "github.com/cosmos/cosmos-sdk/types"

	actionstypes "github.com/forbole/bdjuno/v2/cmd/actions/types"
	"github.com/forbole/bdjuno/v2/cmd/utils"
	"github.com/forbole/bdjuno/v2/database"
	"github.com/forbole/bdjuno/v2/modules"
	"github.com/forbole/bdjuno/v2/modules/bank"
	"github.com/forbole/bdjuno/v2/types/config"
	"github.com/forbole/juno/v2/cmd/parse"
	junoconfig "github.com/forbole/juno/v2/types/config"
)

func getAccountBalances(input actionstypes.AccountBalancesArgs) ([]sdk.Coins, error) {
	parseCtx, sources, err := getCtxAndSources()
	if err != nil {
		return []sdk.Coins{}, err
	}

	bankModule := bank.NewModule(nil, sources.BankSource, parseCtx.EncodingConfig.Marshaler, nil)

	height := input.Height
	if height == 0 {
		// Get latest height if height input is empty
		height, err = parseCtx.Node.LatestHeight()
		if err != nil {
			return []sdk.Coins{}, fmt.Errorf("error while getting chain latest block height: %s", err)
		}
	}

	balances, err := bankModule.Keeper.GetBalances([]string{input.Address}, height)
	if err != nil {
		return []sdk.Coins{}, err
	}

	coins := make([]sdk.Coins, len(balances))
	for index, bal := range balances {
		coins[index] = bal.Balance
	}

	return coins, nil
}

func getDelegatorRewards(address string) (response []actionstypes.DelegatorRewards, err error) {
	parseCtx, sources, err := getCtxAndSources()
	if err != nil {
		return response, err
	}

	// Get latest node height
	height, err := parseCtx.Node.LatestHeight()
	if err != nil {
		return response, fmt.Errorf("error while getting chain latest block height: %s", err)
	}
	rewards, err := sources.DistrSource.DelegatorTotalRewards(address, height)
	if err != nil {
		return response, err
	}

	res := make([]actionstypes.DelegatorRewards, len(rewards))
	for index, rew := range rewards {
		res[index] = actionstypes.DelegatorRewards{
			DecCoins:   rew.Reward,
			ValAddress: rew.ValidatorAddress,
		}
	}

	return res, nil
}

func getValidatorCommission(address string) (response actionstypes.DecCoins, err error) {
	parseCtx, sources, err := getCtxAndSources()
	if err != nil {
		return response, err
	}

	// Get latest node height
	height, err := parseCtx.Node.LatestHeight()
	if err != nil {
		return response, fmt.Errorf("error while getting chain latest block height: %s", err)
	}
	commission, err := sources.DistrSource.ValidatorCommission(address, height)
	if err != nil {
		return response, err
	}

	response = actionstypes.DecCoins{
		DecCoins: commission,
	}

	return response, nil
}

func getTotalSupply() (response actionstypes.Coins, err error) {
	parseCtx, sources, err := getCtxAndSources()
	if err != nil {
		return response, err
	}

	bankModule := bank.NewModule(nil, sources.BankSource, parseCtx.EncodingConfig.Marshaler, nil)

	// Get latest height
	height, err := parseCtx.Node.LatestHeight()
	if err != nil {
		return response, fmt.Errorf("error while getting chain latest block height: %s", err)
	}

	coins, err := bankModule.Keeper.GetSupply(height)
	if err != nil {
		return response, err
	}

	response = actionstypes.Coins{
		Coins: coins,
	}

	return response, nil
}

func getCtxAndSources() (*parse.Context, *modules.Sources, error) {
	parseCfg := parse.NewConfig().
		WithDBBuilder(database.Builder).
		WithEncodingConfigBuilder(config.MakeEncodingConfig(utils.GetBasicManagers())).
		WithRegistrar(modules.NewRegistrar(utils.GetAddressesParser()))

	parseCtx, err := parse.GetParsingContext(parseCfg)
	if err != nil {
		return nil, nil, err
	}

	sources, err := modules.BuildSources(junoconfig.Cfg.Node, parseCtx.EncodingConfig)
	if err != nil {
		return nil, nil, err
	}

	return parseCtx, sources, nil
}

func graphQLError(w http.ResponseWriter, err error) {
	errorObject := actionstypes.GraphQLError{
		Message: err.Error(),
	}
	errorBody, _ := json.Marshal(errorObject)
	w.WriteHeader(http.StatusBadRequest)
	w.Write(errorBody)
}
