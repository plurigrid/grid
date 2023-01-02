package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"grid/x/grid/types"
)

func (k msgServer) SubmitPower(goCtx context.Context, msg *types.MsgSubmitPower) (*types.MsgSubmitPowerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgSubmitPowerResponse{}, nil
}
