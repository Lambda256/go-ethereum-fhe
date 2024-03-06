package fhevm

import (
	"encoding/binary"
	"errors"
	"github.com/ethereum/go-ethereum/crypto"
)

// A method available in the fhelib precompile that can run and estimate gas
type FheLibMethod struct {
	// name of the fhelib function
	name string
	// types of the arguments that the fhelib function take. format is "(type1,type2...)" (e.g "(uint256,bytes1)")
	arg_types           string
	requiredGasFunction func(input []byte) uint64
	runFunction         func(input []byte) ([]byte, error)
}

func (fheLibMethod *FheLibMethod) Name() string {
	return fheLibMethod.name
}

func makeKeccakSignature(input string) uint32 {
	return binary.BigEndian.Uint32(crypto.Keccak256([]byte(input))[0:4])
}

// Return the computed signature by concatenating the name and the arg types of the method
func (fheLibMethod *FheLibMethod) Signature() uint32 {
	return makeKeccakSignature(fheLibMethod.name + fheLibMethod.arg_types)
}

func (fheLibMethod *FheLibMethod) RequiredGas(input []byte) uint64 {
	return fheLibMethod.requiredGasFunction(input)
}

func (fheLibMethod *FheLibMethod) Run(input []byte) ([]byte, error) {
	return fheLibMethod.runFunction(input)
}

// Mapping between function signatures and the functions to call
var signatureToFheLibMethod = map[uint32]*FheLibMethod{}

func GetFheLibMethod(signature uint32) (fheLibMethod *FheLibMethod, found bool) {
	fheLibMethod, found = signatureToFheLibMethod[signature]
	return
}

// All methods available in the fhelib precompile
var fhelibMethods = []*FheLibMethod{
	{
		name:                "fheAdd",
		arg_types:           "(uint256,uint256,bytes1)",
		requiredGasFunction: fheAddSubRequiredGas,
		runFunction:         fheAddRun,
	},
	{
		name:                "fheSub",
		arg_types:           "(uint256,uint256,bytes1)",
		requiredGasFunction: fheAddSubRequiredGas,
		runFunction:         fheSubRun,
	},
	{
		name:                "fheLe",
		arg_types:           "(uint256,uint256,bytes1)",
		requiredGasFunction: fheLeRequiredGas,
		runFunction:         fheLeRun,
	},
	{
		name:                "fheLt",
		arg_types:           "(uint256,uint256,bytes1)",
		requiredGasFunction: fheLtRequiredGas,
		runFunction:         fheLtRun,
	},
	{
		name:                "fheEq",
		arg_types:           "(uint256,uint256,bytes1)",
		requiredGasFunction: fheEqRequiredGas,
		runFunction:         fheEqRun,
	},
	{
		name:                "fheGe",
		arg_types:           "(uint256,uint256,bytes1)",
		requiredGasFunction: fheGeRequiredGas,
		runFunction:         fheGeRun,
	},
	{
		name:                "fheGt",
		arg_types:           "(uint256,uint256,bytes1)",
		requiredGasFunction: fheGtRequiredGas,
		runFunction:         fheGtRun,
	},
	{
		name:                "fheNe",
		arg_types:           "(uint256,uint256,bytes1)",
		requiredGasFunction: fheNeRequiredGas,
		runFunction:         fheNeRun,
	},
	{
		name:                "fheNot",
		arg_types:           "(uint256)",
		requiredGasFunction: fheNotRequiredGas,
		runFunction:         fheNotRun,
	},
	{
		name:                "trivialEncrypt",
		arg_types:           "(uint256,bytes1)",
		requiredGasFunction: trivialEncryptRequiredGas,
		runFunction:         trivialEncryptRun,
	},
}

func init() {
	// create the mapping for every available fhelib method
	for _, method := range fhelibMethods {
		signatureToFheLibMethod[method.Signature()] = method
	}
}

func minInt(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func isScalarOp(input []byte) (bool, error) {
	if len(input) != 65 {
		return false, errors.New("input needs to contain two 256-bit sized values and 1 8-bit value")
	}
	isScalar := (input[64] == 1)
	return isScalar, nil
}
