package ibc

import (
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/mosaicnetworks/chatterboxhub/x/ibc/02-client"
	"github.com/mosaicnetworks/chatterboxhub/x/ibc/02-client/tendermint"
	"github.com/mosaicnetworks/chatterboxhub/x/ibc/04-channel"
	"github.com/mosaicnetworks/chatterboxhub/x/ibc/23-commitment"
	"github.com/mosaicnetworks/chatterboxhub/x/ibc/23-commitment/merkle"
)

var ModuleCdc *codec.Codec

func init() {
	ModuleCdc = codec.New()
	RegisterCodec(ModuleCdc)
}

func RegisterCodec(cdc *codec.Codec) {
	client.RegisterCodec(cdc)
	tendermint.RegisterCodec(cdc)
	channel.RegisterCodec(cdc)
	commitment.RegisterCodec(cdc)
	merkle.RegisterCodec(cdc)
}
