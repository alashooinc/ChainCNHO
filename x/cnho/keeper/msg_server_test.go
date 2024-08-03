package keeper_test

import (
	"context"
	"testing"

	keepertest "cnho/testutil/keeper"
	"cnho/x/cnho/keeper"
	"cnho/x/cnho/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.CnhoKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
