package karmaservice

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewHandler returns a handler for "karmaservice" type messages.
func NewHandler(keeper Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {
		case MsgSetMoniker:
			return handleMsgSetMoniker(ctx, keeper, msg)
		case MsgSetKarma:
			return handleMsgSetName(ctx, keeper, msg)
		default:
			errMsg := fmt.Sprintf("Unrecognized karmaservice Msg type: %v", msg.Type())
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}

// Handle a message to set name
func handleMsgSetMoniker(ctx sdk.Context, keeper Keeper, msg MsgSetMoniker) sdk.Result {
	//	if !msg.Owner.Equals(keeper.GetOwner(ctx, msg.Name)) { // Checks if the the msg sender is the same as the current owner
	//		return sdk.ErrUnauthorized("Incorrect Owner").Result() // If not, throw an error
	//	}
	keeper.SetMoniker(ctx, msg.Moniker, msg.Owner) // If so, set the name to the value specified in the msg.
	return sdk.Result{}                            // return
}

// Handle a message to set name
func handleMsgSetKarma(ctx sdk.Context, keeper Keeper, msg MsgSetKarma) sdk.Result {
	//	if !msg.Owner.Equals(keeper.GetOwner(ctx, msg.Name)) { // Checks if the the msg sender is the same as the current owner
	//		return sdk.ErrUnauthorized("Incorrect Owner").Result() // If not, throw an error
	//	}
	keeper.SetKarma(ctx, msg.Owner, msg.Karma) // If so, set the name to the value specified in the msg.
	return sdk.Result{}                        // return
}
