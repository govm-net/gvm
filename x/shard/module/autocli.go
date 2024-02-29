package shard

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "github.com/govm-net/gvm/api/gvm/shard"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: modulev1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "Shows the parameters of the module",
				},
				{
					RpcMethod: "ShardAll",
					Use:       "list-shard",
					Short:     "List all shard",
				},
				{
					RpcMethod:      "Shard",
					Use:            "show-shard [id]",
					Short:          "Shows a shard",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}},
				},
				{
					RpcMethod: "ToParentAll",
					Use:       "list-to-parent",
					Short:     "List all to-parent",
				},
				{
					RpcMethod:      "ToParent",
					Use:            "show-to-parent [id]",
					Short:          "Shows a to-parent by id",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				{
					RpcMethod: "ToLeftChildAll",
					Use:       "list-to-left-child",
					Short:     "List all to-left-child",
				},
				{
					RpcMethod:      "ToLeftChild",
					Use:            "show-to-left-child [id]",
					Short:          "Shows a to-left-child by id",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				{
					RpcMethod: "ToRightChildAll",
					Use:       "list-to-right-child",
					Short:     "List all to-right-child",
				},
				{
					RpcMethod:      "ToRightChild",
					Use:            "show-to-right-child [id]",
					Short:          "Shows a to-right-child by id",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				// this line is used by ignite scaffolding # autocli/query
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:              modulev1.Msg_ServiceDesc.ServiceName,
			EnhanceCustomCommand: true, // only required if you want to use the custom command
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "UpdateParams",
					Skip:      true, // skipped because authority gated
				},
				{
					RpcMethod:      "CreateShard",
					Use:            "create-shard [index]",
					Short:          "Send a create-shard tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}},
				},
				{
					RpcMethod:      "SendToShard",
					Use:            "send-to-shard [direction] [amount] [info]",
					Short:          "Send a send-to-shard tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "direction"}, {ProtoField: "amount"}, {ProtoField: "info"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
