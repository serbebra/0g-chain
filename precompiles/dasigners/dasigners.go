package dasigners

import (
	"fmt"
	"math/big"
	"strings"

	precopmiles_common "github.com/0glabs/0g-chain/precompiles/common"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/evmos/ethermint/x/evm/statedb"
)

const (
	PrecompileAddress        = "0x0000000000000000000000000000000000001000"
	RequiredGasBasic  uint64 = 100

	DASignersFunctionEpochNumber = "epochNumber"
)

var _ vm.PrecompiledContract = &DASignersPrecompile{}

type DASignersPrecompile struct {
	abi abi.ABI
}

func NewDASignersPrecompile() (*DASignersPrecompile, error) {
	abi, err := abi.JSON(strings.NewReader(DASignersABI))
	if err != nil {
		return nil, err
	}
	return &DASignersPrecompile{
		abi: abi,
	}, nil
}

// Address implements vm.PrecompiledContract.
func (d *DASignersPrecompile) Address() common.Address {
	return common.HexToAddress(PrecompileAddress)
}

// RequiredGas implements vm.PrecompiledContract.
func (d *DASignersPrecompile) RequiredGas(input []byte) uint64 {
	return RequiredGasBasic
}

// Run implements vm.PrecompiledContract.
func (d *DASignersPrecompile) Run(evm *vm.EVM, contract *vm.Contract, readonly bool) ([]byte, error) {
	// parse input
	if len(contract.Input) < 4 {
		return nil, vm.ErrExecutionReverted
	}
	method, err := d.abi.MethodById(contract.Input[:4])
	if err != nil {
		return nil, vm.ErrExecutionReverted
	}
	args, err := method.Inputs.Unpack(contract.Input[4:])
	if err != nil {
		return nil, err
	}
	// get state db and context
	stateDB, ok := evm.StateDB.(*statedb.StateDB)
	if !ok {
		return nil, fmt.Errorf(precopmiles_common.ErrGetStateDB)
	}
	ctx := stateDB.GetContext()

	_ = ctx
	_ = args

	var res []byte
	switch method.Name {
	case DASignersFunctionEpochNumber:
		res, err = method.Outputs.Pack(big.NewInt(10))
		if err != nil {
			return nil, err
		}
	}
	return res, nil
}
