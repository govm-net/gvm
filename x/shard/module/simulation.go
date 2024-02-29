package shard

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/govm-net/gvm/testutil/sample"
	shardsimulation "github.com/govm-net/gvm/x/shard/simulation"
	"github.com/govm-net/gvm/x/shard/types"
)

// avoid unused import issue
var (
	_ = shardsimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgCreateShard = "op_weight_msg_create_shard"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateShard int = 100

	opWeightMsgSendToShard = "op_weight_msg_send_to_shard"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSendToShard int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	shardGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&shardGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// ProposalContents doesn't return any content functions for governance proposals.
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateShard int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateShard, &weightMsgCreateShard, nil,
		func(_ *rand.Rand) {
			weightMsgCreateShard = defaultWeightMsgCreateShard
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateShard,
		shardsimulation.SimulateMsgCreateShard(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgSendToShard int
	simState.AppParams.GetOrGenerate(opWeightMsgSendToShard, &weightMsgSendToShard, nil,
		func(_ *rand.Rand) {
			weightMsgSendToShard = defaultWeightMsgSendToShard
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSendToShard,
		shardsimulation.SimulateMsgSendToShard(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateShard,
			defaultWeightMsgCreateShard,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				shardsimulation.SimulateMsgCreateShard(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgSendToShard,
			defaultWeightMsgSendToShard,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				shardsimulation.SimulateMsgSendToShard(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
