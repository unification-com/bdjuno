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
