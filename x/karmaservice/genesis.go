package karmaservice

import (
	types "github.com/mosaicnetworks/chatterboxhub/x/karmaservice/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

type GenesisState struct {
	//	WhoisRecords []Whois `json:"whois_records"`
	KarmaRecords types.KarmaRecordList `json:"karma_records"`
}

func NewGenesisState(karmaRecords types.KarmaRecordList) GenesisState {
	return GenesisState{}
}

func ValidateGenesis(data GenesisState) error {
	//	for _, record := range data.KarmaRecords.Karma {
	//		if record.Owner == nil {
	//			return fmt.Errorf("Invalid WhoisRecord: Value: %s. Error: Missing Owner", record.Value)
	//		}
	//		if record.Value == "" {
	//			return fmt.Errorf("Invalid WhoisRecord: Owner: %s. Error: Missing Value", record.Owner)
	//		}
	//		if record.Price == nil {
	//			return fmt.Errorf("Invalid WhoisRecord: Value: %s. Error: Missing Price", record.Value)
	//		}
	//	}
	return nil
}

func DefaultGenesisState() GenesisState {
	return GenesisState{KarmaRecords: types.KarmaRecordList{}}
}

func InitGenesis(ctx sdk.Context, keeper Keeper, data GenesisState) []abci.ValidatorUpdate {
	for _, record := range data.KarmaRecords.Karma {
		keeper.SetKarma(ctx, record.Address, record.Karma)
	}
	return []abci.ValidatorUpdate{}
}

func ExportGenesis(ctx sdk.Context, k Keeper) GenesisState {
	var records types.KarmaRecordList
	iterator := k.GetKarmaRecordsIterator(ctx)
	for ; iterator.Valid(); iterator.Next() {
		name := string(iterator.Key())
		addr, _ := sdk.AccAddressFromBech32(name)
		var karma KarmaRecord
		karma = k.GetKarma(ctx, addr)
		records.Karma = append(records.Karma, karma)
	}
	return GenesisState{KarmaRecords: records}
}
