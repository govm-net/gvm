package keeper

import (
	"encoding/hex"
	"errors"
	"fmt"

	"cosmossdk.io/log"
	"cosmossdk.io/math"
	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/mempool"
	"github.com/govm-net/gvm/x/gvm/types"
)

type VoteExtHandler struct {
	dfHandler baseapp.DefaultProposalHandler
	cdc       codec.BinaryCodec
	logger    log.Logger
	Keeper    Keeper
	valStore  baseapp.ValidatorStore
}

func NewVoteExtHandler(
	logger log.Logger,
	keeper Keeper,
	valStore baseapp.ValidatorStore,
	mp mempool.Mempool,
	txVerifier baseapp.ProposalTxVerifier,
) *VoteExtHandler {
	return &VoteExtHandler{
		cdc:       keeper.cdc,
		logger:    logger,
		Keeper:    keeper,
		valStore:  valStore,
		dfHandler: *baseapp.NewDefaultProposalHandler(mp, txVerifier),
	}
}

// OracleVoteExtension defines the canonical vote extension structure.
type OracleVoteExtension struct {
	Height int64
	Prices map[string]math.LegacyDec
}

func (h *VoteExtHandler) ExtendVoteHandler() sdk.ExtendVoteHandler {
	return func(ctx sdk.Context, req *abci.RequestExtendVote) (*abci.ResponseExtendVote, error) {
		return &abci.ResponseExtendVote{VoteExtension: []byte("hello")}, nil
	}
}

func (h *VoteExtHandler) VerifyVoteExtensionHandler() sdk.VerifyVoteExtensionHandler {
	return func(ctx sdk.Context, req *abci.RequestVerifyVoteExtension) (*abci.ResponseVerifyVoteExtension, error) {
		return &abci.ResponseVerifyVoteExtension{Status: abci.ResponseVerifyVoteExtension_ACCEPT}, nil
	}
}

func (h *VoteExtHandler) PrepareProposal() sdk.PrepareProposalHandler {
	return func(ctx sdk.Context, req *abci.RequestPrepareProposal) (*abci.ResponsePrepareProposal, error) {
		resp, err := h.dfHandler.PrepareProposalHandler()(ctx, req)
		if err != nil {
			return nil, err
		}
		var proposalTxs [][]byte
		// proposalTxs := resp.Txs
		fmt.Println("PrepareProposal", req.Height)
		h.logger.Info("PrepareProposal", "Height", req.Height)
		if req.Height >= ctx.ConsensusParams().Abci.VoteExtensionsEnableHeight && ctx.ConsensusParams().Abci.VoteExtensionsEnableHeight != 0 {
			// err := baseapp.ValidateVoteExtensions(ctx, h.valStore, req.Height, ctx.ChainID(), req.LocalLastCommit)
			// if err != nil {
			// 	return nil, err
			// }
			voteExt := types.ShardBlock{
				Index:      uint64(req.Height),
				Parent:     hex.EncodeToString(ctx.HeaderHash()),
				LeftChild:  "1",
				RightChild: "2",
				Vdf:        "3",
			}
			bz, err := h.cdc.Marshal(&voteExt)
			if err != nil {
				h.logger.Error("failed to encode injected vote extension tx", "err", err)
				return nil, errors.New("failed to encode injected vote extension tx")
			}

			// Inject a "fake" tx into the proposal s.t. validators can decode, verify,
			// and store the canonical stake-weighted average prices.
			proposalTxs = append(proposalTxs, bz)
		}
		proposalTxs = append(proposalTxs, resp.Txs...)

		return &abci.ResponsePrepareProposal{
			Txs: proposalTxs,
		}, nil
	}
}

func (h *VoteExtHandler) ProcessProposal() sdk.ProcessProposalHandler {
	return func(ctx sdk.Context, req *abci.RequestProcessProposal) (*abci.ResponseProcessProposal, error) {
		h.logger.Info("ProcessProposal", "height", req.Height, "tx length", len(req.Txs))
		fmt.Println("ProcessProposal", req.Height)
		if req.Height >= ctx.ConsensusParams().Abci.VoteExtensionsEnableHeight && ctx.ConsensusParams().Abci.VoteExtensionsEnableHeight != 0 {
			if len(req.Txs) == 0 {
				return &abci.ResponseProcessProposal{Status: abci.ResponseProcessProposal_REJECT}, nil
			}
			tx := req.Txs[0]
			req.Txs = req.Txs[1:]
			var injectedVoteExtTx types.ShardBlock
			if err := h.cdc.Unmarshal(tx, &injectedVoteExtTx); err != nil {
				h.logger.Error("failed to decode injected vote extension tx", "err", err)
				return &abci.ResponseProcessProposal{Status: abci.ResponseProcessProposal_REJECT}, nil
			}
			h.logger.Info("ProcessProposal tx", "injectedVoteExtTx", injectedVoteExtTx, "parent", injectedVoteExtTx.Parent)
			fmt.Println("ProcessProposal", injectedVoteExtTx)
		}
		return h.dfHandler.ProcessProposalHandler()(ctx, req)
	}
}

func (h *VoteExtHandler) PreBlocker(ctx sdk.Context, req *abci.RequestFinalizeBlock) (*sdk.ResponsePreBlock, error) {
	res := &sdk.ResponsePreBlock{}
	if len(req.Txs) == 0 {
		return res, nil
	}

	var voteExt types.ShardBlock
	if err := h.cdc.Unmarshal(req.Txs[0], &voteExt); err != nil {
		h.logger.Error("failed to decode injected vote extension tx", "err", err)
		return nil, err
	}

	// set oracle prices using the passed in context, which will make these prices available in the current block
	// if err := h.keeper.SetOraclePrices(ctx, injectedVoteExtTx.StakeWeightedPrices); err != nil {
	// 	return nil, err
	// }

	h.logger.Info("PreBlocker", "voteExt", voteExt)
	fmt.Println("PreBlocker", voteExt, req.Time)
	return res, nil
}
