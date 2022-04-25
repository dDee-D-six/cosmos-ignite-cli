package scavenge_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "github.com/username/scavenge/testutil/keeper"
	"github.com/username/scavenge/testutil/nullify"
	"github.com/username/scavenge/x/scavenge"
	"github.com/username/scavenge/x/scavenge/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		ScavengeList: []types.Scavenge{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.ScavengeKeeper(t)
	scavenge.InitGenesis(ctx, *k, genesisState)
	got := scavenge.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.ScavengeList, got.ScavengeList)
	// this line is used by starport scaffolding # genesis/test/assert
}
