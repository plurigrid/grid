package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "grid/testutil/keeper"
	"grid/x/grid/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.GridKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
