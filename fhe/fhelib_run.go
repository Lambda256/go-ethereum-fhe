package fhevm

import (
	"fmt"
	lambda_fhe "github.com/ethereum/go-ethereum/fhe/lambda"
	"github.com/holiman/uint256"
	"golang.org/x/crypto/chacha20"
	"os"
)

type fheLibInterface interface {
	FheAddRun(input []byte) ([]byte, error)
	FheSubRun(input []byte) ([]byte, error)
	FheLeRun(input []byte) ([]byte, error)
	FheLtRun(input []byte) ([]byte, error)
	FheEqRun(input []byte) ([]byte, error)
	FheGeRun(input []byte) ([]byte, error)
	FheGtRun(input []byte) ([]byte, error)
	FheNeRun(input []byte) ([]byte, error)
	FheNotRun(input []byte) ([]byte, error)
	TrivialEncryptRun(input []byte) ([]byte, error)
}

var fhelib fheLibInterface

func init() {

	var fheType, _ = os.LookupEnv("FHE_TYPE")

	switch fheType {
	case "ZAMA":
		fhelib = lambda_fhe.LambdaFhe{}
		fmt.Println("ZAMA FHE")
	case "LAMBDA":
		fhelib = lambda_fhe.LambdaFhe{}
		fmt.Println("LAMBDA FHE")
	default:
		fhelib = lambda_fhe.LambdaFhe{}
		fmt.Println("LAMBDA FHE")
	}

}

func fheAddRun(input []byte) ([]byte, error) {
	input = input[:minInt(65, len(input))]

	return fhelib.FheAddRun(input)
}

func fheSubRun(input []byte) ([]byte, error) {
	input = input[:minInt(65, len(input))]

	return fhelib.FheSubRun(input)
}

func fheLeRun(input []byte) ([]byte, error) {
	input = input[:minInt(65, len(input))]

	return fhelib.FheLeRun(input)
}

func fheLtRun(input []byte) ([]byte, error) {
	input = input[:minInt(65, len(input))]

	return fhelib.FheLtRun(input)
}

func fheEqRun(input []byte) ([]byte, error) {
	input = input[:minInt(65, len(input))]

	return fhelib.FheEqRun(input)
}

func fheGeRun(input []byte) ([]byte, error) {
	input = input[:minInt(65, len(input))]

	return fhelib.FheGeRun(input)
}

func fheGtRun(input []byte) ([]byte, error) {
	input = input[:minInt(65, len(input))]

	return fhelib.FheGtRun(input)
}

func fheNeRun(input []byte) ([]byte, error) {
	input = input[:minInt(65, len(input))]

	return fhelib.FheNeRun(input)
}

func fheNotRun(input []byte) ([]byte, error) {
	input = input[:minInt(65, len(input))]

	return fhelib.FheNotRun(input)
}

func trivialEncryptRun(input []byte) ([]byte, error) {
	input = input[:minInt(65, len(input))]
	return fhelib.TrivialEncryptRun(input)
}

var globalRngSeed []byte

var rngNonceKey [32]byte = uint256.NewInt(0).Bytes32()

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
