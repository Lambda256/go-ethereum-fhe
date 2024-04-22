package lambda_fhe

import (
	"fmt"
	"os"
	"path"
)

func initGlobalKeysFromFiles(keysDir string) ([]byte, error) {
	if _, err := os.Stat(keysDir); os.IsNotExist(err) {
		return nil, fmt.Errorf("init_keys: global keys directory doesn't exist (FHEVM_GO_KEYS_DIR): %s", keysDir)
	}
	// read keys from files
	var sksPath = path.Join(keysDir, "sks")
	sksBytes, err := os.ReadFile(sksPath)
	if err != nil {
		return nil, err
	}
	fmt.Println("INFO: global keys loaded from: " + keysDir)

	return sksBytes, nil
}
