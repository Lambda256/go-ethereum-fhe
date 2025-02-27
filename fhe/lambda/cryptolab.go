package lambda_fhe

/*
#cgo LDFLAGS: -L. -llambda256_cryptolab -lstdc++ -lHEaaN -lomp
#include "cryptoLab.h"
*/
import "C"
import (
	"fmt"
)
import "unsafe"

func mainTest() {

	keyDirPath := "/Users/kevin.park/fhe/HEaaN-0.3.0/examples/build/src/keyPackDir"
	cKeyDirPath := C.CString(keyDirPath)

	//generateKey(cSecretKeyPath, cKeyDirPath)
	cryptoLab := createCrytoLabByKeyDir(cKeyDirPath)
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

func createCrytoLabBySeceryKeyAndKeyDir(cSecretKeyPath *C.char, cKeyDirPath *C.char) unsafe.Pointer {

	return C.createCrytoLabBySeceryKeyAndKeyDir(cSecretKeyPath, cKeyDirPath)
}

func createCrytoLabByKeyDir(cKeyDirPath *C.char) unsafe.Pointer {

	return C.createCrytoLabByKeyDir(cKeyDirPath)
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
