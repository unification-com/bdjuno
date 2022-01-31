package source

import (
	minttypes "github.com/MonikaCat/osmosis/v6/x/mint/types"
)

type Source interface {
	Params(height int64) (minttypes.Params, error)
}