package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	actionstypes "github.com/forbole/bdjuno/v2/cmd/actions/types"
	"github.com/forbole/bdjuno/v2/modules/bank"
)

func TotalSupply(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	result, err := getTotalSupply()
	if err != nil {
		graphQLError(w, err)
		return
	}

	data, _ := json.Marshal(result)
	w.Write(data)
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

	return actionstypes.Coins{
		Coins: coins,
	}, nil
}
