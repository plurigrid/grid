package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"grid/x/grid/types"
)

func (k msgServer) SubmitEstimate(goCtx context.Context, msg *types.MsgSubmitEstimate) (*types.MsgSubmitEstimateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgSubmitEstimateResponse{}, nil
}
