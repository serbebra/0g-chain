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

func (k Keeper) EpochSignerSet(c context.Context, request *types.QueryEpochSignerSetRequest) (*types.QueryEpochSignerSetResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	epochSignerSet := make([]*types.Signer, 0)
	signers, found := k.GetEpochSignerSet(ctx, request.EpochNumber)
	if !found {
		return &types.QueryEpochSignerSetResponse{Signers: epochSignerSet}, nil
	}
	for _, account := range signers.Signers {
		signer, found, err := k.GetSigner(ctx, account)
		if err != nil {
			return nil, err
		}
		if !found {
			return nil, types.ErrSignerNotFound
		}
		epochSignerSet = append(epochSignerSet, &signer)
	}
	return &types.QueryEpochSignerSetResponse{Signers: epochSignerSet}, nil
}
