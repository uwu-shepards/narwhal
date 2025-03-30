package keeper

import (
	"github.com/uwupunks/narwhal/x/narwhal/types"
)

var _ types.QueryServer = Keeper{}
