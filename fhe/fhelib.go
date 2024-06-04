package fhevm

import (
	"encoding/binary"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// A method available in the fhelib precompile that can run and estimate gas
type FheLibMethod struct {
	// name of the fhelib function
	name string
	// types of the arguments that the fhelib function take. format is "(type1,type2...)" (e.g "(uint256,bytes1)")
	arg_types           string
	requiredGasFunction func(input []byte) uint64
	runFunction         func(accessibleState common.StateDBForPrecompiledContract, caller common.Address, addr common.Address, input []byte, isEthCall bool, isGasEstimation bool) ([]byte, error)
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

func (fheLibMethod *FheLibMethod) Run(accessibleState common.StateDBForPrecompiledContract, caller common.Address, addr common.Address, input []byte, isEthCall bool, isGasEstimation bool) ([]byte, error) {
	return fheLibMethod.runFunction(accessibleState, caller, addr, input, isEthCall, isGasEstimation)
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
		arg_types:           "(uint256,bytes,bytes)",
		requiredGasFunction: fheAddSubRequiredGas,
		runFunction:         fheAddRun,
	},
	{
		name:                "fheAddScalar",
		arg_types:           "(uint256,uint256,bytes)",
		requiredGasFunction: fheAddSubRequiredGas,
		runFunction:         fheAddScalarRun,
	},
	{
		name:                "fheSub",
		arg_types:           "(uint256,bytes,bytes)",
		requiredGasFunction: fheAddSubRequiredGas,
		runFunction:         fheSubRun,
	},
	{
		name:                "fheSubScalar",
		arg_types:           "(uint256,uint256,bytes)",
		requiredGasFunction: fheAddSubRequiredGas,
		runFunction:         fheSubScalarRun,
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
	{
		name:                "registerKey",
		arg_types:           "(uint256,uint256,uint256)",
		requiredGasFunction: registerKeyRequiredGas,
		runFunction:         registerKeyRun,
	},
	{
		name:                "addKeyBytes",
		arg_types:           "(uint256,bytes)",
		requiredGasFunction: registerKeyRequiredGas,
		runFunction:         addKeyBytesRun,
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
