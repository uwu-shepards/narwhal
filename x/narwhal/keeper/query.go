package keeper

import (
	"narwhal/x/narwhal/types"
)

var _ types.QueryServer = Keeper{}
