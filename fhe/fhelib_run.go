package fhevm

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	lambda_fhe "github.com/ethereum/go-ethereum/fhe/lambda"
	zama_fhe "github.com/ethereum/go-ethereum/fhe/zama"
	"golang.org/x/crypto/chacha20"
	"os"
)

type fheLibInterface interface {
	GetType() string
	FheAddRun(input []byte) ([]byte, error)
	FheAddScalarRun(input []byte) ([]byte, error)
	FheSubRun(input []byte) ([]byte, error)
	FheSubScalarRun(input []byte) ([]byte, error)
	FheLeRun(input []byte) ([]byte, error)
	FheLtRun(input []byte) ([]byte, error)
	FheEqRun(input []byte) ([]byte, error)
	FheGeRun(input []byte) ([]byte, error)
	FheGtRun(input []byte) ([]byte, error)
	FheNeRun(input []byte) ([]byte, error)
	FheNotRun(input []byte) ([]byte, error)
	InitKey()
	TrivialEncryptRun(input []byte) ([]byte, error)
	RegisterKeyRun(common.StateDBForPrecompiledContract, common.Address, common.Address, []byte, bool, bool) ([]byte, error)
	AddKeyBytesRun(common.StateDBForPrecompiledContract, common.Address, common.Address, []byte, bool, bool) ([]byte, error)
}

var fheLib fheLibInterface
var fheLibImplMap = map[string]fheLibInterface{}
var fheLibImpls = []fheLibInterface{&zama_fhe.ZamaFhe{}, &lambda_fhe.LambdaFhe{}}

func init() {
	insertUniqueFheLibTypeIntoMap()
	assignFheTypeImplementation()
	fheLib.InitKey()
}

func insertUniqueFheLibTypeIntoMap() {
	// check duplicate
	seen := make(map[string]bool)
	for _, fheLibImpl := range fheLibImpls {
		fheLibType := fheLibImpl.GetType()
		fmt.Println("fheLibType : ", fheLibType)
		if _, ok := seen[fheLibType]; ok {
			panic(fmt.Errorf("fheLibImpl Type is duplicated"))
		} else {
			seen[fheLibType] = true
		}
		fheLibImplMap[fheLibType] = fheLibImpl
	}
}

func assignFheTypeImplementation() {
	var fheType, ok = os.LookupEnv("FHE_TYPE")
	if !ok {
		fheType = "LAMBDA"
		//panic(fmt.Errorf("FHE_TYPE must be specified"))
	}

	if fheLibImpl, ok := fheLibImplMap[fheType]; ok {
		fheLib = fheLibImpl
	} else {
		panic(fmt.Errorf("FHE_TYPE does not exist"))
	}
}

func fheAddRun(accessibleState common.StateDBForPrecompiledContract, caller common.Address, addr common.Address, input []byte, isEthCall bool, isGasEstimation bool) ([]byte, error) {

	return fheLib.FheAddRun(input)
}

func fheAddScalarRun(accessibleState common.StateDBForPrecompiledContract, caller common.Address, addr common.Address, input []byte, isEthCall bool, isGasEstimation bool) ([]byte, error) {

	return fheLib.FheAddScalarRun(input)
}

func fheSubRun(accessibleState common.StateDBForPrecompiledContract, caller common.Address, addr common.Address, input []byte, isEthCall bool, isGasEstimation bool) ([]byte, error) {

	return fheLib.FheSubRun(input)
}

func fheSubScalarRun(accessibleState common.StateDBForPrecompiledContract, caller common.Address, addr common.Address, input []byte, isEthCall bool, isGasEstimation bool) ([]byte, error) {

	return fheLib.FheSubScalarRun(input)
}

func fheLeRun(accessibleState common.StateDBForPrecompiledContract, caller common.Address, addr common.Address, input []byte, isEthCall bool, isGasEstimation bool) ([]byte, error) {
	input = input[:minInt(65, len(input))]

	return fheLib.FheLeRun(input)
}

func fheLtRun(accessibleState common.StateDBForPrecompiledContract, caller common.Address, addr common.Address, input []byte, isEthCall bool, isGasEstimation bool) ([]byte, error) {
	input = input[:minInt(65, len(input))]

	return fheLib.FheLtRun(input)
}

func fheEqRun(accessibleState common.StateDBForPrecompiledContract, caller common.Address, addr common.Address, input []byte, isEthCall bool, isGasEstimation bool) ([]byte, error) {
	input = input[:minInt(65, len(input))]

	return fheLib.FheEqRun(input)
}

func fheGeRun(accessibleState common.StateDBForPrecompiledContract, caller common.Address, addr common.Address, input []byte, isEthCall bool, isGasEstimation bool) ([]byte, error) {
	input = input[:minInt(65, len(input))]

	return fheLib.FheGeRun(input)
}

func fheGtRun(accessibleState common.StateDBForPrecompiledContract, caller common.Address, addr common.Address, input []byte, isEthCall bool, isGasEstimation bool) ([]byte, error) {
	input = input[:minInt(65, len(input))]

	return fheLib.FheGtRun(input)
}

func fheNeRun(accessibleState common.StateDBForPrecompiledContract, caller common.Address, addr common.Address, input []byte, isEthCall bool, isGasEstimation bool) ([]byte, error) {
	input = input[:minInt(65, len(input))]

	return fheLib.FheNeRun(input)
}

func fheNotRun(accessibleState common.StateDBForPrecompiledContract, caller common.Address, addr common.Address, input []byte, isEthCall bool, isGasEstimation bool) ([]byte, error) {
	input = input[:minInt(65, len(input))]

	return fheLib.FheNotRun(input)
}

func trivialEncryptRun(accessibleState common.StateDBForPrecompiledContract, caller common.Address, addr common.Address, input []byte, isEthCall bool, isGasEstimation bool) ([]byte, error) {
	input = input[:minInt(65, len(input))]
	return fheLib.TrivialEncryptRun(input)
}

func registerKeyRun(accessibleState common.StateDBForPrecompiledContract, caller common.Address, addr common.Address, input []byte, isEthCall bool, isGasEstimation bool) ([]byte, error) {
	if !isGasEstimation && !isEthCall {
		return fheLib.RegisterKeyRun(accessibleState, caller, addr, input, isEthCall, isGasEstimation)
	}
	return input[:], nil
}

func addKeyBytesRun(accessibleState common.StateDBForPrecompiledContract, caller common.Address, addr common.Address, input []byte, isEthCall bool, isGasEstimation bool) ([]byte, error) {
	//accessibleState.GetState()
	if !isGasEstimation && !isEthCall {
		return fheLib.AddKeyBytesRun(accessibleState, caller, addr, input, isEthCall, isGasEstimation)
	}
	return input[:], nil
}

var globalRngSeed []byte

func init() {
	if chacha20.NonceSizeX != 24 {
		panic("expected 24 bytes for NonceSizeX")
	}

	// TODO: Since the current implementation is not FHE-based and, hence, not private,
	// we just initialize the global seed with non-random public data. We will change
	// that once the FHE version is available.
	globalRngSeed = make([]byte, chacha20.KeySize)
	for i := range globalRngSeed {
		globalRngSeed[i] = byte(1 + i)
	}
}
