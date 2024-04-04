package bounty

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "Bounty/api/bounty/bounty"
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
					RpcMethod: "PlatformAll",
					Use:       "list-platform",
					Short:     "List all platform",
				},
				{
					RpcMethod:      "Platform",
					Use:            "show-platform [id]",
					Short:          "Shows a platform",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}},
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
					RpcMethod:      "CreatePlatform",
					Use:            "create-platform [index]",
					Short:          "Create a new platform",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}},
				},
				{
					RpcMethod:      "UpdatePlatform",
					Use:            "update-platform [index]",
					Short:          "Update platform",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}},
				},
				{
					RpcMethod:      "DeletePlatform",
					Use:            "delete-platform [index]",
					Short:          "Delete platform",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
