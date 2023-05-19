package keeper

import (
	"fmt"

	"github.com/cometbft/cometbft/libs/log"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/osmosis-labs/bech32-ibc/x/bech32ibc/types"
)

type (
	Keeper struct {
		channelKeeper types.ChannelKeeper

		cdc      codec.Codec
		storeKey storetypes.StoreKey

		tk types.TransferKeeper
	}
)

func NewKeeper(
	channelKeeper types.ChannelKeeper,
	cdc codec.Codec,
	storeKey storetypes.StoreKey,
	tk types.TransferKeeper,
) *Keeper {
	return &Keeper{
		channelKeeper: channelKeeper,
		cdc:           cdc,
		storeKey:      storeKey,
		tk:            tk,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
