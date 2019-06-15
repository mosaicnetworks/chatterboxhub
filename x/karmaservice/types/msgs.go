package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const RouterKey = ModuleName // this was defined in your key.go file

type MsgSetMoniker struct {
	Moniker string         `json:"moniker"`
	Owner   sdk.AccAddress `json:"owner"`
}
type MsgSetKarma struct {
	Karma int            `json:"karma"`
	Owner sdk.AccAddress `json:"owner"`
}

func NewMsgSetKarma(karma int, owner sdk.AccAddress) MsgSetKarma {
	return MsgSetKarma{
		Karma: karma,
		Owner: owner,
	}
}

func NewMsgSetMoniker(moniker string, owner sdk.AccAddress) MsgSetMoniker {
	return MsgSetMoniker{
		Moniker: moniker,
		Owner:   owner,
	}
}

// Route should return the name of the module
func (msg MsgSetKarma) Route() string   { return RouterKey }
func (msg MsgSetMoniker) Route() string { return RouterKey }

// Type should return the action
func (msg MsgSetMoniker) Type() string { return "set_moniker" }
func (msg MsgSetKarma) Type() string   { return "set_karma" }

// ValidateBasic runs stateless checks on the message
func (msg MsgSetMoniker) ValidateBasic() sdk.Error {
	if msg.Owner.Empty() {
		return sdk.ErrInvalidAddress(msg.Owner.String())
	}
	//	if len(msg.Moniker) == 0 || len(msg.Value) == 0 {
	//		return sdk.ErrUnknownRequest("Name and/or Value cannot be empty")
	//	}
	return nil
}

func (msg MsgSetKarma) ValidateBasic() sdk.Error {
	if msg.Owner.Empty() {
		return sdk.ErrInvalidAddress(msg.Owner.String())
	}
	//	if len(msg.Moniker) == 0 || len(msg.Value) == 0 {
	//		return sdk.ErrUnknownRequest("Name and/or Value cannot be empty")
	//	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgSetMoniker) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgSetMoniker) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
}

// GetSignBytes encodes the message for signing
func (msg MsgSetKarma) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgSetKarma) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
}
