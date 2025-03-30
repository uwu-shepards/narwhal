package keeper_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/uwupunks/narwhal/x/narwhal/keeper"

	keepertest "github.com/uwupunks/narwhal/testutil/keeper"
	"github.com/uwupunks/narwhal/x/narwhal/types"
)

func setupMsgServer(t testing.TB) (keeper.Keeper, types.MsgServer, context.Context) {
	k, ctx := keepertest.NarwhalKeeper(t)
	return k, keeper.NewMsgServerImpl(k), ctx
}

func TestMsgServer(t *testing.T) {
	k, ms, ctx := setupMsgServer(t)
	require.NotNil(t, ms)
	require.NotNil(t, ctx)
	require.NotEmpty(t, k)
}
