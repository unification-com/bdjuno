package handlers

import (
	"encoding/json"
	"net/http"

	actionstypes "github.com/forbole/bdjuno/v2/cmd/actions/types"
	"github.com/forbole/bdjuno/v2/cmd/utils"
	"github.com/forbole/bdjuno/v2/database"
	"github.com/forbole/bdjuno/v2/modules"
	"github.com/forbole/bdjuno/v2/types/config"
	"github.com/forbole/juno/v2/cmd/parse"
	junoconfig "github.com/forbole/juno/v2/types/config"
)

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
