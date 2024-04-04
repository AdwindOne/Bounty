package platform

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "Bounty/api/bounty/platform"
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
				{
					RpcMethod: "ClaimAll",
					Use:       "list-claim",
					Short:     "List all Claim",
				},
				{
					RpcMethod:      "Claim",
					Use:            "show-claim [id]",
					Short:          "Shows a Claim",
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
					Use:            "create-platform [index] [title] [description] [reward] [creator1]",
					Short:          "Create a new platform",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}, {ProtoField: "title"}, {ProtoField: "description"}, {ProtoField: "reward"}, {ProtoField: "creator1"}},
				},
				{
					RpcMethod:      "UpdatePlatform",
					Use:            "update-platform [index] [title] [description] [reward] [creator1]",
					Short:          "Update platform",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}, {ProtoField: "title"}, {ProtoField: "description"}, {ProtoField: "reward"}, {ProtoField: "creator1"}},
				},
				{
					RpcMethod:      "DeletePlatform",
					Use:            "delete-platform [index]",
					Short:          "Delete platform",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}},
				},
				{
					RpcMethod:      "CreateClaim",
					Use:            "create-claim [index] [bountyId] [hacker] [status]",
					Short:          "Create a new Claim",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}, {ProtoField: "bountyId"}, {ProtoField: "hacker"}, {ProtoField: "status"}},
				},
				{
					RpcMethod:      "UpdateClaim",
					Use:            "update-claim [index] [bountyId] [hacker] [status]",
					Short:          "Update Claim",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}, {ProtoField: "bountyId"}, {ProtoField: "hacker"}, {ProtoField: "status"}},
				},
				{
					RpcMethod:      "DeleteClaim",
					Use:            "delete-claim [index]",
					Short:          "Delete Claim",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}},
				},
				{
					RpcMethod:      "CreateBounty",
					Use:            "create-bounty [title] [description] [reward]",
					Short:          "Send a CreateBounty tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "title"}, {ProtoField: "description"}, {ProtoField: "reward"}},
				},
				{
					RpcMethod:      "ClaimBounty",
					Use:            "claim-bounty [bounty-id]",
					Short:          "Send a ClaimBounty tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "bountyId"}},
				},
				{
					RpcMethod:      "CompleteBounty",
					Use:            "complete-bounty [claim-id]",
					Short:          "Send a CompleteBounty tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "claimId"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
