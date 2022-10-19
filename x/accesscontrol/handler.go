package accesscontrol

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	"github.com/cosmos/cosmos-sdk/x/accesscontrol/keeper"
	"github.com/cosmos/cosmos-sdk/x/accesscontrol/types"
)

func HandleMsgUpdateResourceDependencyMappingProposal(ctx sdk.Context, k *keeper.Keeper, p *types.MsgUpdateResourceDependencyMappingProposal) error {
	for _, resourceDepMapping := range p.MessageDependencyMapping {
		k.SetResourceDependencyMapping(ctx, resourceDepMapping)
	}
	return nil
}

func HandleMsgUpdateWasmDependencyMappingProposal(ctx sdk.Context, k *keeper.Keeper, p *types.MsgUpdateWasmDependencyMappingProposal) error {
	err := k.SetWasmDependencyMapping(ctx, sdk.AccAddress(p.ContractAddress), p.WasmDependencyMapping)
	if err != nil {
		return err
	}
	return nil
}

func NewProposalHandler(k keeper.Keeper) govtypes.Handler {
	return func(ctx sdk.Context, content govtypes.Content) error {
		switch c := content.(type) {
		case *types.MsgUpdateResourceDependencyMappingProposal:
			return HandleMsgUpdateResourceDependencyMappingProposal(ctx, &k, c)
		case *types.MsgUpdateWasmDependencyMappingProposal:
			return HandleMsgUpdateWasmDependencyMappingProposal(ctx, &k, c)
		default:
			return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized accesscontrol proposal content type: %T", c)
		}
	}
}
