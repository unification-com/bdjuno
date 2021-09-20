package utils

import (
	issuertypes "github.com/e-money/em-ledger/x/issuer/types"
	"github.com/forbole/bdjuno/database"
	"github.com/forbole/bdjuno/types"
)

func StoreEmoneyInflationFromMessage(
	height int64, msg *issuertypes.MsgSetInflation, db *database.Db,
) error {
	inflation := types.NewEmoneyInflation(msg.Issuer, msg.Denom, msg.InflationRate, height)

	//save  in emoney_inflation table and in mint table
	err := db.SaveInflation(inflation.Rate, inflation.Height)
	if err != nil {
		return err
	}

	return nil
}
