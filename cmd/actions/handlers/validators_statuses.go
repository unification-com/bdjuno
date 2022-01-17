package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	actionstypes "github.com/forbole/bdjuno/v2/cmd/actions/types"

	"github.com/forbole/bdjuno/v2/modules/slashing"
	"google.golang.org/grpc/codes"
)

func ValidatorsStatuses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "invalid payload", http.StatusBadRequest)
		return
	}

	var actionPayload actionstypes.ValidatorsStatusesPayload
	err = json.Unmarshal(reqBody, &actionPayload)
	if err != nil {
		http.Error(w, "invalid payload: failed to unmarshal json", http.StatusInternalServerError)
		return
	}

	result, err := getValidatorsStatuses(actionPayload.Input)
	if err != nil {
		graphQLError(w, err)
		return
	}

	data, _ := json.Marshal(result)
	w.Write(data)
}

func getValidatorsStatuses(input actionstypes.ValidatorsStatusesArgs) (statuses actionstypes.ValidatorsStatuses, err error) {
	parseCtx, sources, err := getCtxAndSources()
	if err != nil {
		return statuses, err
	}

	height := input.Height
	if height == 0 {
		// Get latest height if height input is empty
		height, err = parseCtx.Node.LatestHeight()
		if err != nil {
			return statuses, fmt.Errorf("error while getting chain latest block height: %s", err)
		}
	}

	// Get validators from the node with status if specified
	status := stakingtypes.BondStatus(input.Status).String()
	if input.Status == 0 {
		status = ""
	}
	validators, err := sources.StakingSource.GetValidatorsWithStatus(height, status)
	if err != nil {
		return statuses, err
	}

	slashingModule := slashing.NewModule(sources.SlashingSource, nil, nil, nil)

	res := make([]actionstypes.ValidatorStatus, len(validators))
	for index, val := range validators {
		// Get consesus address
		var pubKey cryptotypes.PubKey
		err := parseCtx.EncodingConfig.Marshaler.UnpackAny(val.ConsensusPubkey, &pubKey)
		if err != nil {
			return statuses, err
		}
		consAddr := sdk.ConsAddress(pubKey.Address())

		// Get signing info for tombstoned status
		valSigningInfo, err := slashingModule.GetSigningInfo(height, consAddr)
		if err != nil && !strings.Contains(err.Error(), codes.NotFound.String()) {
			return statuses, fmt.Errorf("error while getting validator signing info: %s", err)
		}

		res[index] = actionstypes.ValidatorStatus{
			ValidatorAddress: consAddr.String(),
			Status:           int(val.Status),
			Jailed:           val.Jailed,
			Tombstoned:       valSigningInfo.Tombstoned,
		}
	}

	statuses = actionstypes.ValidatorsStatuses{
		ValidatorsStatuses: res,
	}

	return statuses, nil
}
