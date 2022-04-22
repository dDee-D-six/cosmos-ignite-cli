package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/username/blog/x/blog/types"
)

func (k msgServer) DeleteComment(goCtx context.Context, msg *types.MsgDeleteComment) (*types.MsgDeleteCommentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgDeleteCommentResponse{}, nil
}
