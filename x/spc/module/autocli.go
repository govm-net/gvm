package spc

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "github.com/govm-net/gvm/api/gvm/spc"
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
					RpcMethod:      "Transfer",
					Use:            "transfer [app-id] [authenticator-data] [client-data] [credential] [signature]",
					Short:          "Send a transfer tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "appId"}, {ProtoField: "authenticatorData"}, {ProtoField: "clientData"}, {ProtoField: "credential"}, {ProtoField: "signature"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
