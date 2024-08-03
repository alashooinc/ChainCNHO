package cnho_test

import (
	"testing"

	keepertest "cnho/testutil/keeper"
	"cnho/testutil/nullify"
	"cnho/x/cnho"
	"cnho/x/cnho/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CnhoKeeper(t)
	cnho.InitGenesis(ctx, *k, genesisState)
	got := cnho.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
