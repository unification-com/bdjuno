package types

import sdk "github.com/cosmos/cosmos-sdk/types"

type GraphQLError struct {
	Message string `json:"message"`
}

// ========================= Account Balances =========================
type AddressPayload struct {
	SessionVariables map[string]interface{} `json:"session_variables"`
	Input            Address                `json:"input"`
}

type Address struct {
	Address string
}

// ========================= Coins =========================
type Coins struct {
	Coins []sdk.Coin
}

type DecCoins struct {
	DecCoins []sdk.DecCoin
}
