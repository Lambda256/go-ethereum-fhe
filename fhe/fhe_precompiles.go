package fhevm

import (
	"encoding/binary"
	"errors"
	"fmt"
)

func FheLibRequiredGas(input []byte) uint64 {

	return 10
}

func FheLibRun(input []byte) (ret []byte, err error) {
	fmt.Println("FheLibRun!!")

	if len(input) < 4 {
		err := errors.New("input must contain at least 4 bytes for method signature")
		return nil, err
	}
	// first 4 bytes are for the function signature
	signature := binary.BigEndian.Uint32(input[0:4])
	fmt.Println("signature ::", signature)
	fheLibMethod, found := GetFheLibMethod(signature)

	if !found {
		err := errors.New("precompile method not found")
		return nil, err
	}
	fmt.Println("fheLibMethod ::", fheLibMethod.name)
	// remove function signature
	input = input[4:]

	ret, err = fheLibMethod.Run(input)
	ret = ret[:]

	return
}
