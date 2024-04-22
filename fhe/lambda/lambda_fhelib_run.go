package lambda_fhe

import "C"
import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/holiman/uint256"
	"math/big"
	"os"
	"unsafe"
)

type LambdaFhe struct {
	sks       []byte
	name      string
	cryptoLab unsafe.Pointer
}

func (z *LambdaFhe) GetType() string {
	return "LAMBDA"
}

func (l *LambdaFhe) InitKey() {
	secretKeyPath := "/Users/kevin.park/fhe/HEaaN-0.3.0/examples/build/src/secretKeyDir/secretkey.bin"
	cSecretKeyPath := C.CString(secretKeyPath)
	keyDirPath := "/Users/kevin.park/fhe/HEaaN-0.3.0/examples/build/src/keyPackDir"
	cKeyDirPath := C.CString(keyDirPath)
	l.cryptoLab = createCrytoLab(cSecretKeyPath, cKeyDirPath)
	var keysDirPath, present = os.LookupEnv("LAMBDA_KEYS_DIR")
	fmt.Println("Lambda keysDirPath :", keysDirPath)
	fmt.Println("l.cryptoLab :", l.cryptoLab)
	if present {
		sksBytes, err := initGlobalKeysFromFiles(keysDirPath)
		if err != nil {
			panic(err)
		}
		l.sks = sksBytes
		fmt.Println("INFO: global keys are initialized automatically using FHEVM_GO_KEYS_DIR env variable")
	} else {
		fmt.Println("INFO: global keys aren't initialized automatically (FHEVM_GO_KEYS_DIR env variable not set)")
	}
}

func (l *LambdaFhe) FheAddRun(input []byte) ([]byte, error) {
	fmt.Println("FheAddRun")
	fmt.Println("len(input :", len(input))
	if !checkInputDoubleCipherTextLength(input) {
		return nil, fmt.Errorf("input is invalid")
	}
	leftValue, rightValue := getCipherValues(input)
	lhs := toBytes(leftValue)
	rhs := toBytes(rightValue)

	cipherText := add(l.cryptoLab, lhs, rhs)
	cipherText = toEVMBytes(cipherText)

	return toOutputBytes(cipherText), nil
}

func (l *LambdaFhe) FheAddScalarRun(input []byte) ([]byte, error) {
	fmt.Println("FheAddScalarRun")
	fmt.Println("len(input :", len(input))
	leftValue, rightValue := getPlainAndCipherValue(input)
	value := *new(big.Int).SetBytes(leftValue)
	lhs := value.Uint64()
	rhs := toBytes(rightValue)

	cipherText := addScalar(l.cryptoLab, lhs, rhs)

	cipherText = toEVMBytes(cipherText)
	return toOutputBytes(cipherText), nil
}

func (l *LambdaFhe) FheSubRun(input []byte) ([]byte, error) {

	if !checkInputDoubleCipherTextLength(input) {
		return nil, fmt.Errorf("input is invalid")
	}
	fmt.Println("FheSubRun")
	fmt.Println("len(input :", len(input))
	leftValue, rightValue := getCipherValues(input)
	lhs := toBytes(leftValue)
	rhs := toBytes(rightValue)

	cipherText := sub(l.cryptoLab, lhs, rhs)
	cipherText = toEVMBytes(cipherText)

	return toOutputBytes(cipherText), nil
}

func (l *LambdaFhe) FheSubScalarRun(input []byte) ([]byte, error) {
	fmt.Println("FheSubScalarRun")
	fmt.Println("len(input :", len(input))
	leftValue, rightValue := getPlainAndCipherValue(input)
	value := *new(big.Int).SetBytes(leftValue)
	lhs := value.Uint64()
	rhs := toBytes(rightValue)

	cipherText := subScalar(l.cryptoLab, lhs, rhs)

	cipherText = toEVMBytes(cipherText)
	return toOutputBytes(cipherText), nil
}

func (l *LambdaFhe) FheLeRun(input []byte) ([]byte, error) {
	//leftValue, rightValue := getLefetAndRightValue(input)
	value := 0

	bytes := Uint256ToBytes(uint64(value))
	return bytes[:], nil

}

func (l *LambdaFhe) FheLtRun(input []byte) ([]byte, error) {

	return nil, nil
}

func (l *LambdaFhe) FheEqRun(input []byte) ([]byte, error) {

	return nil, nil
}

func (l *LambdaFhe) FheGeRun(input []byte) ([]byte, error) {

	return nil, nil
}

func (l *LambdaFhe) FheGtRun(input []byte) ([]byte, error) {

	return nil, nil
}

func (l *LambdaFhe) FheNeRun(input []byte) ([]byte, error) {

	return nil, nil
}

func (l *LambdaFhe) FheNotRun(input []byte) ([]byte, error) {

	return nil, nil
}

func (l *LambdaFhe) TrivialEncryptRun(input []byte) ([]byte, error) {
	ret := []byte{0}
	return ret, nil
}

func getPlainAndCipherValue(input []byte) ([]byte, []byte) {

	leftValue := input[0:32]
	rightValue := input[64:]
	return leftValue, rightValue
}
func getCipherValues(input []byte) ([]byte, []byte) {
	input = input[64:]
	leftValue := input[:len(input)/2]
	rightValue := input[len(input)/2:]
	return leftValue, rightValue
}

func BytesToUint32(byteArray []byte) uint32 {
	var result uint32

	// 바이트 배열의 길이 계산
	byteLength := len(byteArray)

	// 바이트 배열의 크기가 4바이트 미만이면 0 반환
	if byteLength < 4 {
		return 0
	}

	// 바이트 배열을 uint32로 변환
	for i := 0; i < 4; i++ {
		shift := uint((3 - i) * 8)
		result |= uint32(byteArray[byteLength-4+i]) << shift
	}

	return result
}

func Uint256ToBytes(num uint64) []byte {
	uint256 := big.NewInt(int64(num))
	byteArray := uint256.Bytes()
	byteArray = append(make([]byte, 32-len(byteArray)), byteArray...)
	return byteArray
}

// apply padding to slice to the multiple of 32
func padArrayTo32Multiple(input []byte) []byte {
	modRes := len(input) % 32
	if modRes > 0 {
		padding := 32 - modRes
		for padding > 0 {
			padding--
			input = append(input, 0x0)
		}
	}
	return input
}

func toOutputBytes(input []byte) []byte {
	ret := make([]byte, 32, len(input)+32)
	ret[31] = 0x20
	ret = append(ret, input...)
	ret = padArrayTo32Multiple(ret)
	return ret
}

func toBytes(input []byte) []byte {
	fmt.Println("len(ret)1 : ", len(input))
	ret := common.TrimRightZeroes(input)
	fmt.Println("len(ret)2 : ", len(ret))
	fmt.Println("ret[:32] :", ret[:32])

	ret = fromEVMBytes(ret)
	return ret
}

func toEVMBytes(input []byte) []byte {
	arrLen := uint64(len(input))
	lenBytes32 := uint256.NewInt(arrLen).Bytes32()
	ret := make([]byte, 0, arrLen+32)
	ret = append(ret, lenBytes32[:]...)
	ret = append(ret, input...)
	return ret
}

func fromEVMBytes(input []byte) []byte {
	// 반환된 바이트 배열에서 처음 32바이트를 추출하여 uint256로 변환합니다.
	var arrLenBytes32 [32]byte
	copy(arrLenBytes32[:], input[:32])
	arrLen := new(big.Int).SetBytes(arrLenBytes32[:]).Uint64()

	// 원래의 바이트 배열을 추출합니다.
	return input[32 : 32+arrLen]
}

func checkInputOneCipherTextLength(input []byte) bool {
	return len(input) == 524608
}

func checkInputDoubleCipherTextLength(input []byte) bool {
	return len(input) == 1049152
}
