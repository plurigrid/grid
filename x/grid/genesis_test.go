package grid_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "grid/testutil/keeper"
	"grid/testutil/nullify"
	"grid/x/grid"
	"grid/x/grid/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.GridKeeper(t)
	grid.InitGenesis(ctx, *k, genesisState)
	got := grid.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
