package lambda_fhe

import "C"
import (
	"bufio"
	"encoding/binary"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/holiman/uint256"
	"math/big"
	"os"
	"path/filepath"
	"strconv"
	"unsafe"
)

type LambdaFhe struct {
	name            string
	keyPath         string
	CryptoLabPtrMap map[uint64]unsafe.Pointer
}

func (z *LambdaFhe) GetType() string {
	return "LAMBDA"
}

func (l *LambdaFhe) InitKey() {

	l.CryptoLabPtrMap = make(map[uint64]unsafe.Pointer)
	var keysDirPath, present = os.LookupEnv("LAMBDA_KEYS_DIR")
	l.keyPath = keysDirPath
	fmt.Println("Lambda keysDirPath :", keysDirPath)
	if present {
		folders, err := initGlobalKeysFromFiles(keysDirPath)
		if err != nil {
			panic(err)
		}
		fmt.Println(" folders :", folders)
		for _, folder := range folders {
			keyDirPath := keysDirPath + strconv.Itoa(folder)
			cKeyDirPath := C.CString(keyDirPath)
			l.CryptoLabPtrMap[uint64(folder)] = createCrytoLabByKeyDir(cKeyDirPath)
			fmt.Println("l.CryptoLabPtrMap : ", l.CryptoLabPtrMap)

		}

		fmt.Println("INFO: global keys are initialized automatically using FHEVM_GO_KEYS_DIR env variable")
	} else {
		fmt.Println("INFO: global keys aren't initialized automatically (FHEVM_GO_KEYS_DIR env variable not set)")
	}
}

func (l *LambdaFhe) FheAddRun(input []byte) ([]byte, error) {
	if !checkInputDoubleCipherTextLength(input) {
		return nil, fmt.Errorf("input is invalid")
	}
	keyNumber, leftValue, rightValue := getCipherValues(input)
	keyNumberInt := *new(big.Int).SetBytes(keyNumber)
	cryptoLabPtr := l.CryptoLabPtrMap[keyNumberInt.Uint64()]
	lhs := toBytes(leftValue)
	rhs := toBytes(rightValue)

	cipherText := add(cryptoLabPtr, lhs, rhs)
	cipherText = toEVMBytes(cipherText)

	return toOutputBytes(cipherText), nil
}

func (l *LambdaFhe) FheAddScalarRun(input []byte) ([]byte, error) {
	keyNumber, leftValue, rightValue := getPlainAndCipherValue(input)
	keyNumberInt := *new(big.Int).SetBytes(keyNumber)
	cryptoLabPtr := l.CryptoLabPtrMap[keyNumberInt.Uint64()]

	if !checkInputOneCipherTextLength(input) {
		return nil, fmt.Errorf("input is invalid")
	}

	value := *new(big.Int).SetBytes(leftValue)
	lhs := value.Uint64()
	rhs := toBytes(rightValue)

	cipherText := addScalar(cryptoLabPtr, lhs, rhs)

	cipherText = toEVMBytes(cipherText)

	return toOutputBytes(cipherText), nil
}

func (l *LambdaFhe) FheSubRun(input []byte) ([]byte, error) {
	if !checkInputDoubleCipherTextLength(input) {
		return nil, fmt.Errorf("input is invalid")
	}

	keyNumber, leftValue, rightValue := getCipherValues(input)
	keyNumberInt := *new(big.Int).SetBytes(keyNumber)
	cryptoLabPtr := l.CryptoLabPtrMap[keyNumberInt.Uint64()]
	lhs := toBytes(leftValue)
	rhs := toBytes(rightValue)

	cipherText := sub(cryptoLabPtr, lhs, rhs)
	cipherText = toEVMBytes(cipherText)

	return toOutputBytes(cipherText), nil
}

func (l *LambdaFhe) FheSubScalarRun(input []byte) ([]byte, error) {
	if !checkInputOneCipherTextLength(input) {
		return nil, fmt.Errorf("input is invalid")
	}
	keyNumber, leftValue, rightValue := getPlainAndCipherValue(input)
	keyNumberInt := *new(big.Int).SetBytes(keyNumber)
	cryptoLabPtr := l.CryptoLabPtrMap[keyNumberInt.Uint64()]

	value := *new(big.Int).SetBytes(leftValue)
	lhs := value.Uint64()
	rhs := toBytes(rightValue)

	cipherText := subScalar(cryptoLabPtr, lhs, rhs)

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

func (l *LambdaFhe) RegisterKeyRun(accessibleState common.StateDBForPrecompiledContract, caller common.Address, addr common.Address, input []byte, isEthCall bool, isGasEstimation bool) ([]byte, error) {

	keyNumber := new(big.Int).SetBytes(input[0:32]).Uint64()
	keyRootStorage := new(big.Int).SetBytes(input[32:64])
	keyLength := new(big.Int).SetBytes(input[64:96]).Uint64()
	var keyByteArray []byte
	for i := 0; i < int(keyLength); i++ {

		plusValue := big.NewInt(int64(i))
		keyStorage := new(big.Int).Add(keyRootStorage, plusValue)

		keyByte := accessibleState.GetState(caller, common.BigToHash(keyStorage))

		keyByteArray = append(keyByteArray, keyByte.Bytes()...)

	}

	keyByteArray = common.TrimRightZeroes(keyByteArray)
	//keyDirPath := keysDirPath + "/" + strconv.Itoa(folder)
	newKeyPath := l.keyPath + strconv.Itoa(int(keyNumber))
	fmt.Println("newKeyPath : ", newKeyPath)
	err := createFileAndWriteBytes(newKeyPath+"/PK/MultKey.bin", keyByteArray)
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("key is invalid", err)
	}
	cKeyDirPath := C.CString(newKeyPath)
	l.CryptoLabPtrMap[keyNumber] = createCrytoLabByKeyDir(cKeyDirPath)
	fmt.Println("l.CryptoLabPtrMap : ", l.CryptoLabPtrMap)

	value := 0
	bytes := Uint256ToBytes(uint64(value))
	return bytes[:], nil
}

func (l *LambdaFhe) AddKeyBytesRun(accessibleState common.StateDBForPrecompiledContract, caller common.Address, addr common.Address, input []byte, isEthCall bool, isGasEstimation bool) ([]byte, error) {

	value := 0
	bytes := Uint256ToBytes(uint64(value))
	return bytes[:], nil
}

func getPlainAndCipherValue(input []byte) ([]byte, []byte, []byte) {
	keyNumber := input[0:32]
	leftValue := input[32:64]
	rightValue := input[96:]
	return keyNumber, leftValue, rightValue
}
func getCipherValues(input []byte) ([]byte, []byte, []byte) {
	keyNumber := input[0:32]
	input = input[96:]
	leftValue := input[:len(input)/2]
	rightValue := input[len(input)/2:]
	return keyNumber, leftValue, rightValue
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
	return len(input) == 524640
}

func checkInputDoubleCipherTextLength(input []byte) bool {
	return len(input) == 1049184
}

func uintToBytes(num uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, num)
	return b
}

func bytesToUint64(b []byte) uint64 {
	return binary.BigEndian.Uint64(b)
}

// createFileAndWriteBytes 함수는 주어진 파일 경로에 주어진 바이트 데이터를 씁니다.
func createFileAndWriteBytes(filePath string, byteData []byte) error {

	// 파일이 속할 폴더 경로 가져오기
	dir := filepath.Dir(filePath)

	// 폴더가 없으면 생성
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		return fmt.Errorf("폴더를 생성하는 중에 오류가 발생했습니다: %v", err)
	}

	// 파일 생성
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("파일을 생성하는 중에 오류가 발생했습니다: %v", err)
	}
	defer file.Close()

	// 파일에 바이트 데이터 쓰기
	writer := bufio.NewWriter(file)
	_, err = writer.Write(byteData)
	if err != nil {
		return fmt.Errorf("파일에 데이터를 쓰는 중에 오류가 발생했습니다: %v", err)
	}

	// 버퍼 비우고 파일에 쓴 내용 저장
	err = writer.Flush()
	if err != nil {
		return fmt.Errorf("파일에 데이터를 저장하는 중에 오류가 발생했습니다: %v", err)
	}

	return nil
}
