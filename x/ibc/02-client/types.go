package client

import (
	"github.com/mosaicnetworks/chatterboxhub/x/ibc/23-commitment"
)

// TODO: types in this file should be (de/)serialized with proto in the future
// currently amkno codec handles it

// ConsensusState
type ConsensusState interface {
	Kind() Kind
	GetHeight() uint64
	GetRoot() commitment.Root
	Validate(Header) (ConsensusState, error) // ValidityPredicate
	Equivocation(Header, Header) bool        // EquivocationPredicate
}

/*
func Equal(client1, client2 ConsensusState) bool {
	return client1.Kind() == client2.Kind() &&
		client1.GetBase().Equal(client2.GetBase())
}
*/

type Header interface {
	Kind() Kind
	GetHeight() uint64
}

// XXX: Kind should be enum?

type Kind byte

const (
	Tendermint Kind = iota
)
