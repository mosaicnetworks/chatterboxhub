package karmaservice

import (
	"github.com/mosaicnetworks/chatterboxhub/x/karmaservice/types"
)

const (
	ModuleName = types.ModuleName
	RouterKey  = types.RouterKey
	StoreKey   = types.StoreKey
)

var (
	NewMsgBuyName = types.NewMsgBuyName
	NewMsgSetName = types.NewMsgSetName
	NewWhois      = types.NewWhois
	ModuleCdc     = types.ModuleCdc
	RegisterCodec = types.RegisterCodec
)

type (
	MsgSetName      = types.MsgSetName
	MsgBuyName      = types.MsgBuyName
	QueryResResolve = types.QueryResResolve
	QueryResNames   = types.QueryResNames
	Whois           = types.Whois
)
