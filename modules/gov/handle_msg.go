package gov

import (
	"fmt"
	"time"

	"strconv"

	"github.com/forbole/bdjuno/v4/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypesv1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"

	gov "github.com/cosmos/cosmos-sdk/x/gov/types"
	govtypesv1beta1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
	juno "github.com/forbole/juno/v4/types"
)

// HandleMsg implements modules.MessageModule
func (m *Module) HandleMsg(index int, msg sdk.Msg, tx *juno.Tx) error {
	if len(tx.Logs) == 0 {
		return nil
	}

	switch cosmosMsg := msg.(type) {
	case *govtypesv1beta1.MsgSubmitProposal:
		return m.handleMsgSubmitProposalV1Beta1(tx, index, cosmosMsg)

	case *govtypesv1beta1.MsgVote:
		return m.handleMsgVoteV1Beta1(tx, cosmosMsg)

	case *govtypesv1beta1.MsgDeposit:
		return m.handleMsgDepositV1Beta1(tx, cosmosMsg)

	case *govtypesv1.MsgDeposit:
		return m.handleMsgDeposit(tx, cosmosMsg)

	case *govtypesv1.MsgVote:
		return m.handleMsgVote(tx, cosmosMsg)
	}

	return nil
}

// handleMsgSubmitProposalV1Beta1 allows to properly handle a handleMsgSubmitProposalV1Beta1
func (m *Module) handleMsgSubmitProposalV1Beta1(tx *juno.Tx, index int, msg *govtypesv1beta1.MsgSubmitProposal) error {
	// Get the proposal id
	event, err := tx.FindEventByType(index, gov.EventTypeSubmitProposal)
	if err != nil {
		return fmt.Errorf("error while searching for EventTypeSubmitProposal: %s", err)
	}

	id, err := tx.FindAttributeByKey(event, gov.AttributeKeyProposalID)
	if err != nil {
		return fmt.Errorf("error while searching for AttributeKeyProposalID: %s", err)
	}

	proposalID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return fmt.Errorf("error while parsing proposal id: %s", err)
	}

	// Get the proposal
	proposal, err := m.source.Proposal(tx.Height, proposalID)
	if err != nil {
		return fmt.Errorf("error while getting proposal: %s", err)
	}

	// Store the proposal
	proposalObj := types.NewProposal(
		proposal.ProposalId,
		msg.GetContent().ProposalRoute(),
		msg.GetContent().ProposalType(),
		msg.GetContent(),
		proposal.Status.String(),
		proposal.SubmitTime,
		proposal.DepositEndTime,
		proposal.VotingStartTime,
		proposal.VotingEndTime,
		msg.Proposer,
	)

	err = m.db.SaveProposals([]types.Proposal{proposalObj})
	if err != nil {
		return err
	}

	txTimestamp, err := time.Parse(time.RFC3339, tx.Timestamp)
	if err != nil {
		return fmt.Errorf("error while parsing time: %s", err)
	}

	// Store the deposit
	deposit := types.NewDeposit(proposal.ProposalId, msg.Proposer, msg.InitialDeposit, txTimestamp, tx.Height)
	return m.db.SaveDeposits([]types.Deposit{deposit})
}

// handleMsgDeposit allows to properly handle a handleMsgDeposit
func (m *Module) handleMsgDeposit(tx *juno.Tx, msg *govtypesv1.MsgDeposit) error {
	deposit, err := m.source.ProposalDeposit(tx.Height, msg.ProposalId, msg.Depositor)
	if err != nil {
		return fmt.Errorf("error while getting proposal deposit: %s", err)
	}
	txTimestamp, err := time.Parse(time.RFC3339, tx.Timestamp)
	if err != nil {
		return fmt.Errorf("error while parsing time: %s", err)
	}

	return m.db.SaveDeposits([]types.Deposit{
		types.NewDeposit(msg.ProposalId, msg.Depositor, deposit.Amount, txTimestamp, tx.Height),
	})
}

// handleMsgDeposit allows to properly handle a handleMsgDeposit
func (m *Module) handleMsgDepositV1Beta1(tx *juno.Tx, msg *govtypesv1beta1.MsgDeposit) error {
	deposit, err := m.source.ProposalDeposit(tx.Height, msg.ProposalId, msg.Depositor)
	if err != nil {
		return fmt.Errorf("error while getting proposal deposit: %s", err)
	}
	txTimestamp, err := time.Parse(time.RFC3339, tx.Timestamp)
	if err != nil {
		return fmt.Errorf("error while parsing time: %s", err)
	}

	return m.db.SaveDeposits([]types.Deposit{
		types.NewDeposit(msg.ProposalId, msg.Depositor, deposit.Amount, txTimestamp, tx.Height),
	})
}

// handleMsgVote allows to properly handle a handleMsgVote
func (m *Module) handleMsgVote(tx *juno.Tx, msg *govtypesv1.MsgVote) error {
	txTimestamp, err := time.Parse(time.RFC3339, tx.Timestamp)
	if err != nil {
		return fmt.Errorf("error while parsing time: %s", err)
	}

	vote := types.NewVote(msg.ProposalId, msg.Voter, msg.Option, txTimestamp, tx.Height)

	return m.db.SaveVote(vote)
}

// handleMsgVoteV1Beta1 allows to properly handle a handleMsgVoteV1Beta1
func (m *Module) handleMsgVoteV1Beta1(tx *juno.Tx, msg *govtypesv1beta1.MsgVote) error {
	txTimestamp, err := time.Parse(time.RFC3339, tx.Timestamp)
	if err != nil {
		return fmt.Errorf("error while parsing time: %s", err)
	}

	vote := types.NewVote(msg.ProposalId, msg.Voter, govtypesv1.VoteOption(msg.Option), txTimestamp, tx.Height)

	return m.db.SaveVote(vote)
}
