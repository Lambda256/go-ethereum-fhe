package lambda_fhe

/*
#cgo LDFLAGS: -L. -lex1-encrypt-decrypt -lstdc++ -lHEaaN -lomp
#include "example.h"
*/
import "C"
import (
	"fmt"
)
import "unsafe"

func mainTest() {

	secretKeyPath := "/Users/kevin.park/fhe/HEaaN-0.3.0/examples/build/src/secretKeyDir/secretkey.bin"
	cSecretKeyPath := C.CString(secretKeyPath)
	keyDirPath := "/Users/kevin.park/fhe/HEaaN-0.3.0/examples/build/src/keyPackDir"
	cKeyDirPath := C.CString(keyDirPath)

	//generateKey(cSecretKeyPath, cKeyDirPath)
	cryptoLab := createCrytoLab(cSecretKeyPath, cKeyDirPath)
	fmt.Println(cryptoLab)
	//encryptResult := encrypt(cryptoLab, 100)
	//fmt.Println("len(encryptResult) : ", len(encryptResult))

	//addScalarResult := addScalar(cryptoLab,encryptResult,1 );

	//finalResult := decrypt(cryptoLab,addScalarResult)

	//addResult := add(cryptoLab,encryptResult,encryptResult );

	//finalResult4 := decrypt(cryptoLab,addResult)

	//fmt.Println("finalResult4 : ", finalResult4)

}

func generateKey(cSecretKeyPath *C.char, cKeyDirPath *C.char) {
	C.generateKey(cSecretKeyPath, cKeyDirPath)
}

func createCrytoLab(cSecretKeyPath *C.char, cKeyDirPath *C.char) unsafe.Pointer {

	return C.createCrytoLab(cSecretKeyPath, cKeyDirPath)
}

func encrypt(cryptoLab unsafe.Pointer, plainText uint64) []byte {
	out := &C.ByteArray{}
	ciphertext := C.encrypt(cryptoLab, C.uint64_t(plainText), out)
	defer C.freeByteArray(ciphertext)
	return C.GoBytes(unsafe.Pointer(ciphertext.data), C.int(ciphertext.length))
}

func add(cryptoLab unsafe.Pointer, goBytes1 []uint8, goBytes2 []uint8) []byte {

	var ciphertext1 C.ByteArray
	ciphertext1.data = (*C.uchar)(unsafe.Pointer(&goBytes1[0]))
	ciphertext1.length = C.int(len(goBytes1))

	var ciphertext2 C.ByteArray
	ciphertext2.data = (*C.uchar)(unsafe.Pointer(&goBytes2[0]))
	ciphertext2.length = C.int(len(goBytes2))

	addResult := C.add(cryptoLab, ciphertext1, ciphertext2)
	defer C.freeByteArray(addResult)

	return C.GoBytes(unsafe.Pointer(addResult.data), C.int(addResult.length))
}

func addScalar(cryptoLab unsafe.Pointer, plainText uint64, goBytes []uint8) []byte {

	var ciphertext2 C.ByteArray
	ciphertext2.data = (*C.uchar)(unsafe.Pointer(&goBytes[0]))
	ciphertext2.length = C.int(len(goBytes))
	addResult := C.addScalar(cryptoLab, ciphertext2, C.uint64_t(plainText))
	defer C.freeByteArray(addResult)

	return C.GoBytes(unsafe.Pointer(addResult.data), C.int(addResult.length))
}

func sub(cryptoLab unsafe.Pointer, goBytes1 []uint8, goBytes2 []uint8) []byte {

	var ciphertext1 C.ByteArray
	ciphertext1.data = (*C.uchar)(unsafe.Pointer(&goBytes1[0]))
	ciphertext1.length = C.int(len(goBytes1))

	var ciphertext2 C.ByteArray
	ciphertext2.data = (*C.uchar)(unsafe.Pointer(&goBytes2[0]))
	ciphertext2.length = C.int(len(goBytes2))

	addResult := C.sub(cryptoLab, ciphertext1, ciphertext2)
	defer C.freeByteArray(addResult)

	return C.GoBytes(unsafe.Pointer(addResult.data), C.int(addResult.length))
}

func subScalar(cryptoLab unsafe.Pointer, plainText uint64, goBytes []uint8) []byte {

	var ciphertext2 C.ByteArray
	ciphertext2.data = (*C.uchar)(unsafe.Pointer(&goBytes[0]))
	ciphertext2.length = C.int(len(goBytes))
	addResult := C.subScalar(cryptoLab, ciphertext2, C.uint64_t(plainText))
	defer C.freeByteArray(addResult)

	return C.GoBytes(unsafe.Pointer(addResult.data), C.int(addResult.length))
}

func decrypt(cryptoLab unsafe.Pointer, goBytes []uint8) int {
	var ciphertext3 C.ByteArray
	ciphertext3.data = (*C.uchar)(unsafe.Pointer(&goBytes[0]))
	ciphertext3.length = C.int(len(goBytes))
	finalResult := C.decrypt(cryptoLab, ciphertext3)
	fmt.Println("finalResult1 : ", finalResult)
	fmt.Println("finalResult2 : ", int(finalResult))
	return int(finalResult)
}
