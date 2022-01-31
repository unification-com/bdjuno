package bank

import (
	// sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

type StakingModule interface {
	GetParams(height int64) (stakingtypes.Params, error)
}
