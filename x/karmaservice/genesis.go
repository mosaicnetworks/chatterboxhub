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

func NewGenesisState(karmaRecords KarmaRecordList) GenesisState {
	return GenesisState{KarmaRecords: nil}
}

func ValidateGenesis(data GenesisState) error {
	for _, record := range data.KarmaRecords {
		//		if record.Owner == nil {
		//			return fmt.Errorf("Invalid WhoisRecord: Value: %s. Error: Missing Owner", record.Value)
		//		}
		//		if record.Value == "" {
		//			return fmt.Errorf("Invalid WhoisRecord: Owner: %s. Error: Missing Value", record.Owner)
		//		}
		//		if record.Price == nil {
		//			return fmt.Errorf("Invalid WhoisRecord: Value: %s. Error: Missing Price", record.Value)
		//		}
	}
	return nil
}

func DefaultGenesisState() GenesisState {
	return GenesisState{KarmaRecords: types.KarmaRecordList{}}
}

func InitGenesis(ctx sdk.Context, keeper Keeper, data GenesisState) []abci.ValidatorUpdate {
	for _, record := range data.KarmaRecords {
		keeper.SetKarma(ctx, record.Value, record)
	}
	return []abci.ValidatorUpdate{}
}

func ExportGenesis(ctx sdk.Context, k Keeper) GenesisState {
	var records []Whois
	iterator := k.GetNamesIterator(ctx)
	for ; iterator.Valid(); iterator.Next() {
		name := string(iterator.Key())
		var karma KarmaRecord
		karma = k.GetKarma(ctx, name)
		records = append(records, karma)
	}
	return GenesisState{KarmaRecords: records}
}
