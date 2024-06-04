package fhevm

import (
	"encoding/binary"
	"errors"
	"github.com/ethereum/go-ethereum/common"
)

func FheLibRequiredGas(input []byte) uint64 {

	return 10
}

func FheLibRun(accessibleState common.StateDBForPrecompiledContract, caller common.Address, addr common.Address, input []byte, isEthCall bool, isGasEstimation bool) (ret []byte, err error) {
	if len(input) < 4 {
		err := errors.New("input must contain at least 4 bytes for method signature")
		return nil, err
	}
	// first 4 bytes are for the function signature
	signature := binary.BigEndian.Uint32(input[0:4])
	fheLibMethod, found := GetFheLibMethod(signature)

	if !found {
		err := errors.New("precompile method not found")
		return nil, err
	}
	// remove function signature
	input = input[4:]

	ret, err = fheLibMethod.Run(accessibleState, caller, addr, input, isEthCall, isGasEstimation)
	ret = ret[:]

	return
}
