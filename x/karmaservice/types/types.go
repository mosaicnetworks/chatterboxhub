package types

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type KarmaRecord struct {
	Address sdk.AccAddress `json:"address"`
	Moniker string         `json:"moniker"`
	Karma   int            `json:"karma"`
}

type KarmaRecordList struct {
	Karma []KarmaRecord
}

func NewKarmaRecord() KarmaRecord {
	return KarmaRecord{
		Karma: 0,
	}
}

// implement fmt.Stringer
func (w KarmaRecord) String() string {
	return strings.TrimSpace(fmt.Sprintf(`Address: %s
Moniker: %s
Karma: %d`, w.Address, w.Moniker, w.Karma))
}
