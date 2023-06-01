package base

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/skip-mev/pob/blockbuster"
)

// PrepareLane will prepare a partial proposal for the base lane.
func (l *DefaultLane) PrepareLane(ctx sdk.Context, maxTxBytes int64, selectedTxs map[string][]byte) ([][]byte, error) {
	txs := make([][]byte, 0)
	txsToRemove := make(map[sdk.Tx]struct{}, 0)

	// Select all transactions in the mempool that are valid and not already in the
	// partial proposal.
	for iterator := l.Mempool.Select(ctx, nil); iterator != nil; iterator = iterator.Next() {
		tx := iterator.Tx()

		txBytes, err := l.TxEncoder(tx)
		if err != nil {
			txsToRemove[tx] = struct{}{}
			continue
		}

		// if the transaction is already in the (partial) block proposal, we skip it.
		hash, err := blockbuster.GetTxHashStr(l.TxEncoder, tx)
		if err != nil {
			txsToRemove[tx] = struct{}{}
			continue
		}
		if _, ok := selectedTxs[hash]; ok {
			continue
		}

		// Verify the transaction.
		if err := l.VerifyTx(ctx, tx); err != nil {
			txsToRemove[tx] = struct{}{}
			continue
		}

		txs = append(txs, txBytes)
	}

	// Remove all transactions that were invalid during the creation of the partial proposal.
	if err := blockbuster.RemoveTxsFromLane(txsToRemove, l.Mempool); err != nil {
		return nil, fmt.Errorf("failed to remove txs from mempool for lane %s: %w", l.Name(), err)
	}

	return txs, nil
}

// ProcessLane verifies the default lane's portion of a block proposal.
func (l *DefaultLane) ProcessLane(ctx sdk.Context, proposalTxs [][]byte, next blockbuster.ProcessLanesHandler) (sdk.Context, error) {
	for index, tx := range proposalTxs {
		tx, err := l.TxDecoder(tx)
		if err != nil {
			return ctx, fmt.Errorf("failed to decode tx: %w", err)
		}

		if l.Match(tx) {
			if err := l.VerifyTx(ctx, tx); err != nil {
				return ctx, fmt.Errorf("failed to verify tx: %w", err)
			}
		} else {
			return next(ctx, proposalTxs[index:])
		}
	}

	// This means we have processed all transactions in the proposal.
	return ctx, nil
}

func (l *DefaultLane) VerifyTx(ctx sdk.Context, tx sdk.Tx) error {
	if l.AnteHandler != nil {
		_, err := l.AnteHandler(ctx, tx, false)
		return err
	}

	return nil
}
