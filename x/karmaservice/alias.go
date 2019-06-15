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
	NewMsgSetMoniker = types.NewMsgSetMoniker
	NewMsgSetKarma   = types.NewMsgSetKarma
	NewKarmaRecord   = types.NewKarmaRecord
	ModuleCdc        = types.ModuleCdc
	RegisterCodec    = types.RegisterCodec
)

type (
	MsgSetMoniker = types.MsgSetMoniker
	MsgSetKarma   = types.MsgSetKarma
	QueryResKarma = types.KarmaRecordList
	KarmaRecord   = types.KarmaRecord
)
