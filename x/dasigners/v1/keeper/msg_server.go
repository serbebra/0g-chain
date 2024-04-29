package keeper

import (
	"github.com/0glabs/0g-chain/x/dasigners/v1/types"
)

var _ types.MsgServer = &Keeper{}
