package gvm

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/govm-net/gvm/x/gvm/keeper"
	"github.com/govm-net/gvm/x/gvm/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the shardBlock
	for _, elem := range genState.ShardBlockList {
		k.SetShardBlock(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.ShardBlockList = k.GetAllShardBlock(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
