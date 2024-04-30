package types

import (
	"encoding/hex"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// ModuleName The name that will be used throughout the module
	ModuleName = "dasigners"

	// StoreKey Top level store key where all module items will be stored
	StoreKey = ModuleName
)

var (
	// prefix
	SignerKeyPrefix         = []byte{0x00}
	EpochSignerSetKeyPrefix = []byte{0x01}
	RegistrationKeyPrefix   = []byte{0x02}

	// keys
	ParamsKey      = []byte{0x05}
	EpochNumberKey = []byte{0x06}
)

func GetSignerKeyFromAccount(account string) ([]byte, error) {
	return hex.DecodeString(account)
}

func GetEpochSignerSetKeyFromEpoch(epoch uint64) []byte {
	return sdk.Uint64ToBigEndian(epoch)
}
