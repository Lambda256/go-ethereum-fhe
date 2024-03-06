package lambda_fhe

import "C"
import (
	"fmt"
	"os"
	"path"
)

var sks []byte

func init() {
	var keysDirPath, present = os.LookupEnv("FHEVM_GO_KEYS_DIR")
	fmt.Println("keysDirPath :", keysDirPath)
	if present {
		err := InitGlobalKeysFromFiles(keysDirPath)
		if err != nil {
			panic(err)
		}
		fmt.Println("INFO: global keys are initialized automatically using FHEVM_GO_KEYS_DIR env variable")
	} else {
		fmt.Println("INFO: global keys aren't initialized automatically (FHEVM_GO_KEYS_DIR env variable not set)")
	}
}

func InitGlobalKeysFromFiles(keysDir string) error {
	if _, err := os.Stat(keysDir); os.IsNotExist(err) {
		return fmt.Errorf("init_keys: global keys directory doesn't exist (FHEVM_GO_KEYS_DIR): %s", keysDir)
	}
	// read keys from files
	var sksPath = path.Join(keysDir, "sks")
	sksBytes, err := os.ReadFile(sksPath)
	if err != nil {
		return err
	}
	sks = sksBytes

	fmt.Println("INFO: global keys loaded from: " + keysDir)

	return nil
}
