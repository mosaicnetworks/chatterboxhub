package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

var ModuleCdc = codec.New()

func init() {
	RegisterCodec(ModuleCdc)
}

// RegisterCodec registers concrete types on the Amino codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgSetMoniker{}, "karmaservice/SetMoniker", nil)
	cdc.RegisterConcrete(MsgSetKarma{}, "karmaservice/SetKarma", nil)
}
