package utils

import (
	"fmt"

	"github.com/forbole/juno/v2/cmd/parse"
)

// RemoveDuplicateValues removes the duplicated values from the given slice
func RemoveDuplicateValues(slice []string) []string {
	keys := make(map[string]bool)
	var list []string

	for _, entry := range slice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

// GetHeight uses the lastest height when the input height is empty from graphql request
func GetHeight(parseCtx *parse.Context, inputHeight int64) (int64, error) {
	if inputHeight == 0 {
		latestHeight, err := parseCtx.Node.LatestHeight()
		if err != nil {
			return 0, fmt.Errorf("error while getting chain latest block height: %s", err)
		}
		return latestHeight, nil
	}

	return inputHeight, nil
}
