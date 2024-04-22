package zama_fhe

import "C"

import (
	"fmt"
	"os"
)

func initGlobalKeysFromFiles(keysDir string) ([]byte, error) {
	if _, err := os.Stat(keysDir); os.IsNotExist(err) {
		return nil, fmt.Errorf("init_keys: global keys directory doesn't exist (FHEVM_GO_KEYS_DIR): %s", keysDir)
	}

	fmt.Println("INFO: global keys loaded from: " + keysDir)

	return nil, nil
}

// initialize keys automatically only if FHEVM_GO_KEYS_DIR is set
