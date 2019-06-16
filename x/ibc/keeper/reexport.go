package ibc

import (
	"github.com/mosaicnetworks/chatterboxhub/x/ibc/02-client"
	"github.com/mosaicnetworks/chatterboxhub/x/ibc/04-channel"
	"github.com/mosaicnetworks/chatterboxhub/x/ibc/23-commitment"
)

type (
	Proof          = commitment.Proof
	ConsensusState = client.ConsensusState
	Packet         = channel.Packet
)
