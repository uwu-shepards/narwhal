package narwhal_test

import (
	"testing"

	keepertest "github.com/uwupunks/narwhal/testutil/keeper"
	"github.com/uwupunks/narwhal/testutil/nullify"
	narwhal "github.com/uwupunks/narwhal/x/narwhal/module"
	"github.com/uwupunks/narwhal/x/narwhal/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.NarwhalKeeper(t)
	narwhal.InitGenesis(ctx, k, genesisState)
	got := narwhal.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
