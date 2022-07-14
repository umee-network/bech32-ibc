package bech32ibc

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	gov1b1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"

	"github.com/osmosis-labs/bech32-ibc/x/bech32ibc/keeper"
	"github.com/osmosis-labs/bech32-ibc/x/bech32ibc/types"
)

// NewHandler returns claim module messages
func NewHandler(k keeper.Keeper) sdk.Handler {

	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		switch msg := msg.(type) {

		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", types.ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}

func NewBech32IBCProposalHandler(k keeper.Keeper) gov1b1.Handler {
	return func(ctx sdk.Context, content gov1b1.Content) error {
		switch c := content.(type) {
		case *types.UpdateHrpIbcChannelProposal:
			return handleUpdateHrpIbcChannelProposal(ctx, k, c)

		default:
			return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized bech32 ibc proposal content type: %T", c)
		}
	}
}

func handleUpdateHrpIbcChannelProposal(ctx sdk.Context, k keeper.Keeper, p *types.UpdateHrpIbcChannelProposal) error {
	return k.HandleUpdateHrpIbcChannelProposal(ctx, p)
}
