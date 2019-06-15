package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type KarmaRecord struct {
	Address sdk.AccAddress `json:"address"`
	Moniker string         `json:"moniker"`
	Karma   int            `json:"karma"`
}

type KarmaRecordList struct {
	Accounts []KarmaRecord
}

func NewKarmaRecord() KarmaRecord {
	return KarmaRecord{
		Karma: 0,
	}
}
