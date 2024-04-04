package types

import (
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	// this line is used by starport scaffolding # 1
)

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreatePlatform{},
		&MsgUpdatePlatform{},
		&MsgDeletePlatform{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateClaim{},
		&MsgUpdateClaim{},
		&MsgDeleteClaim{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateBounty{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgClaimBounty{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCompleteBounty{},
	)
	// this line is used by starport scaffolding # 3

	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdateParams{},
	)
	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}
