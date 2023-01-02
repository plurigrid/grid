package power_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "grid/testutil/keeper"
	"grid/testutil/nullify"
	"grid/x/power"
	"grid/x/power/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.PowerKeeper(t)
	power.InitGenesis(ctx, *k, genesisState)
	got := power.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
