package gov

import (
	"github.com/QOSGroup/qbase/context"
	"github.com/QOSGroup/qbase/txs"
	btypes "github.com/QOSGroup/qbase/types"
	gtypes "github.com/QOSGroup/qos/module/gov/types"
	"github.com/QOSGroup/qos/types"
)

type TxDeposit struct {
	ProposalID uint64         `json:"proposal_id"` // ID of the proposal
	Depositor  btypes.Address `json:"depositor"`   // Address of the depositor
	Amount     uint64         `json:"amount"`      // Amount of QOS to add to the proposal's deposit
}

var _ txs.ITx = (*TxDeposit)(nil)

func (tx TxDeposit) ValidateData(ctx context.Context) error {
	if len(tx.Depositor) == 0 {
		return ErrInvalidInput("depositor is empty")
	}

	if tx.Amount == 0 {
		return ErrInvalidInput("amount of deposit is zero")
	}

	proposal, ok := GetGovMapper(ctx).GetProposal(ctx, tx.ProposalID)
	if !ok {
		return ErrUnknownProposal(tx.ProposalID)
	}

	if (proposal.Status != gtypes.StatusDepositPeriod) && (proposal.Status != gtypes.StatusVotingPeriod) {
		return ErrFinishedProposal(tx.ProposalID)
	}

	return nil
}

func (tx TxDeposit) Exec(ctx context.Context) (result btypes.Result, crossTxQcp *txs.TxQcp) {
	result = btypes.Result{
		Code: btypes.CodeOK,
	}

	err, votingStarted := GetGovMapper(ctx).AddDeposit(ctx, tx.ProposalID, tx.Depositor, tx.Amount)
	if err != nil {
		result = btypes.Result{Code: btypes.CodeInternal, Codespace: btypes.CodespaceType(err.Error())}
	}

	resTags := btypes.NewTags(
		Depositor, []byte(tx.Depositor.String()),
		ProposalID, tx.ProposalID,
	)

	if votingStarted {
		resTags = resTags.AppendTag(VotingPeriodStart, types.Uint64ToBigEndian(tx.ProposalID))
	}
	result.Tags = result.Tags.AppendTags(resTags)

	return
}

func (tx TxDeposit) GetSigner() []btypes.Address {
	return []btypes.Address{tx.Depositor}
}

func (tx TxDeposit) CalcGas() btypes.BigInt {
	return btypes.ZeroInt()
}

func (tx TxDeposit) GetGasPayer() btypes.Address {
	return tx.Depositor
}

func (tx TxDeposit) GetSignData() (ret []byte) {
	ret = append(ret, types.Uint64ToBigEndian(tx.ProposalID)...)
	ret = append(ret, tx.Depositor...)
	ret = append(ret, types.Uint64ToBigEndian(tx.Amount)...)

	return
}
