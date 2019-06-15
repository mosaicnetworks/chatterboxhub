package karmaservice

import (
	"github.com/cosmos/cosmos-sdk/codec"

	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

// query endpoints supported by the karmaservice Querier
const (
	QueryKarma    = "karma"
	QueryAllKarma = "allkarma"
)

// NewQuerier is the module level router for state queries
func NewQuerier(keeper Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err sdk.Error) {
		switch path[0] {
		case QueryKarma:
			return queryKarma(ctx, path[1:], req, keeper)
		case QueryAllKarma:
			return queryAllKarma(ctx, req, keeper)
		default:
			return nil, sdk.ErrUnknownRequest("unknown karmaservice query endpoint")
		}
	}
}

func queryAllKarma(ctx sdk.Context, req abci.RequestQuery, keeper Keeper) ([]byte, sdk.Error) {
	var namesList QueryResKarma

	iterator := keeper.GetKarmaRecordsIterator(ctx)

	for ; iterator.Valid(); iterator.Next() {
		namesList = append(namesList, string(iterator.Key()))
	}

	res, err := codec.MarshalJSONIndent(keeper.cdc, namesList)
	if err != nil {
		panic("could not marshal result to JSON")
	}

	return res, nil
}

func queryKarma(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) (res []byte, err sdk.Error) {
	addr, _ := sdk.AccAddressFromBech32(path[0])

	// returns "yo" or "neeein!"
	value := keeper.GetKarma(ctx, addr)

	bz, err2 := codec.MarshalJSONIndent(keeper.cdc, value)
	if err2 != nil {
		panic("could not marshal result to JSON")
	}

	return bz, nil
}
