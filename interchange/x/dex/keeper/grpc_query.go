package keeper

import (
	"github.com/username/interchange/x/dex/types"
)

var _ types.QueryServer = Keeper{}
