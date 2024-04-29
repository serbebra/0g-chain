package keeper

import (
	"context"

	"github.com/0glabs/0g-chain/x/dasigners/v1/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) EpochNumber(
	c context.Context,
	_ *types.QueryEpochNumberRequest,
) (*types.QueryEpochNumberResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	epochNumber, err := k.GetEpochNumber(ctx)
	if err != nil {
		return nil, err
	}
	return &types.QueryEpochNumberResponse{EpochNumber: epochNumber}, nil
}
