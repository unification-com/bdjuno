package types

import sdk "github.com/cosmos/cosmos-sdk/types"

type GraphQLError struct {
	Message string `json:"message"`
}

// ========================= Address Payload =========================
type AddressPayload struct {
	SessionVariables map[string]interface{} `json:"session_variables"`
	Input            Address                `json:"input"`
}

type Address struct {
	Address string `json:"address"`
}

// ========================= Account Balances Payload =========================
type AccountBalancesPayload struct {
	SessionVariables map[string]interface{} `json:"session_variables"`
	Input            AccountBalancesArgs    `json:"input"`
}

type AccountBalancesArgs struct {
	Address string `json:"address"`
	Height  int64  `json:"height"`
}

// ========================= Validators Statuses Payload =========================
type ValidatorsStatusesPayload struct {
	SessionVariables map[string]interface{} `json:"session_variables"`
	Input            ValidatorsStatusesArgs `json:"input"`
}

type ValidatorsStatusesArgs struct {
	Height int64 `json:"height"`
	Status int32 `json:"status"`
}

// ========================= Validators Statuses Response =========================
type ValidatorStatus struct {
	ValidatorAddress string `json:"validator_address"`
	Status           int    `json:"status"`
	Jailed           bool   `json:"jailed"`
	Tombstoned       bool   `json:"tombstoned"`
}

type ValidatorsStatuses struct {
	ValidatorsStatuses []ValidatorStatus `json:"validators_statuses"`
}

// ========================= Coins =========================
type Coins struct {
	Coins []sdk.Coin `json:"coins"`
}

type DecCoins struct {
	DecCoins []sdk.DecCoin `json:"dec_coins"`
}

// ========================= Delegator Rewards =========================

type DelegatorRewards struct {
	DecCoins   sdk.DecCoins `json:"dec_coins"`
	ValAddress string       `json:"validator_address"`
}
