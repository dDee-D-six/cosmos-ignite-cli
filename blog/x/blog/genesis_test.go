package blog_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "github.com/username/blog/testutil/keeper"
	"github.com/username/blog/testutil/nullify"
	"github.com/username/blog/x/blog"
	"github.com/username/blog/x/blog/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		CommentList: []types.Comment{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		CommentCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.BlogKeeper(t)
	blog.InitGenesis(ctx, *k, genesisState)
	got := blog.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.CommentList, got.CommentList)
	require.Equal(t, genesisState.CommentCount, got.CommentCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
