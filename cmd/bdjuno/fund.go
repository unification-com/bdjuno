package main

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	junomessages "github.com/forbole/juno/v4/modules/messages"
	beacontypes "github.com/unification-com/mainchain/x/beacon/types"
	enttypes "github.com/unification-com/mainchain/x/enterprise/types"
	wrkchaintypes "github.com/unification-com/mainchain/x/wrkchain/types"
)

// fundMessageAddressesParser represents a parser able to get the addresses of the involved
// account from a UND message
var fundMessageAddressesParser = junomessages.JoinMessageParsers(
	beaconMessageAddressesParser,
	wrkchainMessageAddressesParser,
	enterpriseMessageAddressesParser,
)

// beaconMessageAddressesParser represents a MessageAddressesParser for the x/beacon module
func beaconMessageAddressesParser(_ codec.Codec, cosmosMsg sdk.Msg) ([]string, error) {

	switch msg := cosmosMsg.(type) {
	case *beacontypes.MsgRegisterBeacon:
		return []string{msg.Owner}, nil
	case *beacontypes.MsgRecordBeaconTimestamp:
		return []string{msg.Owner}, nil
	}

	return nil, junomessages.MessageNotSupported(cosmosMsg)
}

// wrkchainMessageAddressesParser represents a MessageAddressesParser for the x/wrkchain module
func wrkchainMessageAddressesParser(_ codec.Codec, cosmosMsg sdk.Msg) ([]string, error) {

	switch msg := cosmosMsg.(type) {
	case *wrkchaintypes.MsgRegisterWrkChain:
		return []string{msg.Owner}, nil
	case *wrkchaintypes.MsgRecordWrkChainBlock:
		return []string{msg.Owner}, nil
	}

	return nil, junomessages.MessageNotSupported(cosmosMsg)
}

// enterpriseMessageAddressesParser represents a MessageAddressesParser for the x/enterrprise module
func enterpriseMessageAddressesParser(_ codec.Codec, cosmosMsg sdk.Msg) ([]string, error) {

	switch msg := cosmosMsg.(type) {
	case *enttypes.MsgWhitelistAddress:
		return []string{msg.Address, msg.Signer}, nil
	case *enttypes.MsgUndPurchaseOrder:
		return []string{msg.Purchaser}, nil
	case *enttypes.MsgProcessUndPurchaseOrder:
		return []string{msg.Signer}, nil
	}

	return nil, junomessages.MessageNotSupported(cosmosMsg)
}
