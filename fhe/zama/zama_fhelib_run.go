package zama_fhe

import (
	"fmt"
	"math/big"
	"os"
)

type ZamaFhe struct {
	sks []byte
}

func (z *ZamaFhe) GetType() string {
	return "ZAMA"
}

func (z *ZamaFhe) InitKey() {
	var keysDirPath, present = os.LookupEnv("FHEVM_GO_KEYS_DIR")
	fmt.Println("ZAMA keysDirPath :", keysDirPath)
	if present {
		sksBytes, err := initGlobalKeysFromFiles(keysDirPath)
		if err != nil {
			panic(err)
		}
		z.sks = sksBytes
		fmt.Println("INFO: global keys are initialized automatically using FHEVM_GO_KEYS_DIR env variable")
	} else {
		fmt.Println("INFO: global keys aren't initialized automatically (FHEVM_GO_KEYS_DIR env variable not set)")
	}
}

func (z *ZamaFhe) FheAddRun(input []byte) ([]byte, error) {

	//fmt.Println("len(sks) : ", len(sks))
	leftValue, rightValue := getLefetAndRightValue(input)

	value := leftValue + rightValue
	value += 5
	bytes := Uint256ToBytes(uint64(value))
	return bytes[:], nil
}

func (z *ZamaFhe) FheAddScalarRun(input []byte) ([]byte, error) {

	//fmt.Println("len(sks) : ", len(sks))
	leftValue, rightValue := getLefetAndRightValue(input)

	value := leftValue + rightValue
	value += 5
	//bytes := Uint256ToBytes(uint64(value))
	return input[:], nil
}

func (z *ZamaFhe) FheSubRun(input []byte) ([]byte, error) {
	leftValue, rightValue := getLefetAndRightValue(input)
	value := leftValue - rightValue
	value += 5
	bytes := Uint256ToBytes(uint64(value))
	return bytes[:], nil
}

func (l *ZamaFhe) FheSubScalarRun(input []byte) ([]byte, error) {
	//leftValue, rightValue := getLefetAndRightValue(input)
	//value := leftValue - rightValue
	value := 5
	bytes := Uint256ToBytes(uint64(value))
	return bytes[:], nil
}

func (z *ZamaFhe) FheLeRun(input []byte) ([]byte, error) {
	leftValue, rightValue := getLefetAndRightValue(input)
	value := 0
	if leftValue <= rightValue {
		value = 1
	}
	bytes := Uint256ToBytes(uint64(value))
	return bytes[:], nil

}

func (z *ZamaFhe) FheLtRun(input []byte) ([]byte, error) {

	return nil, nil
}

func (z *ZamaFhe) FheEqRun(input []byte) ([]byte, error) {

	return nil, nil
}

func (z *ZamaFhe) FheGeRun(input []byte) ([]byte, error) {

	return nil, nil
}

func (z *ZamaFhe) FheGtRun(input []byte) ([]byte, error) {

	return nil, nil
}

func (z *ZamaFhe) FheNeRun(input []byte) ([]byte, error) {

	return nil, nil
}

func (z *ZamaFhe) FheNotRun(input []byte) ([]byte, error) {

	return nil, nil
}

func (z *ZamaFhe) TrivialEncryptRun(input []byte) ([]byte, error) {
	ret := []byte{0}
	return ret, nil
}

func getLefetAndRightValue(input []byte) (uint32, uint32) {
	leftValue := BytesToUint32(input[0:32])
	rightValue := BytesToUint32(input[32:64])
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
