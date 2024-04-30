package types

import (
	"encoding/hex"
	fmt "fmt"

	"github.com/0glabs/0g-chain/crypto/bn254"
)

func ValidateHexAddress(account string) error {
	addr, err := hex.DecodeString(account)
	if err != nil {
		return err
	}
	if len(addr) != 20 {
		return fmt.Errorf("invalid address length")
	}
	return nil
}

func (s *Signer) Validate() error {
	if len(s.PubkeyG1) != bn254.G1PubkeySize {
		return fmt.Errorf("invalid G1 pubkey length")
	}
	if len(s.PubkeyG2) != bn254.G2PubkeySize {
		return fmt.Errorf("invalid G2 pubkey length")
	}
	if err := ValidateHexAddress(s.Account); err != nil {
		return err
	}
	return nil
}
