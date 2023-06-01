package keeper

import (
	"context"

	"github.com/alice/checkers/x/checkers/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Help(goCtx context.Context, msg *types.MsgHelp) (*types.MsgHelpResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgHelpResponse{}, nil
}
