package lambda_fhe

import (
	"os"
	"path/filepath"
	"sort"
	"strconv"
)

func initGlobalKeysFromFiles(keysDir string) ([]int, error) {

	return getNumericSubfolders(keysDir)
}

// 폴더 내의 숫자 폴더 목록을 가져오는 함수
func getNumericSubfolders(keysDir string) ([]int, error) {
	var numericFolders []int

	err := filepath.Walk(keysDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// 폴더 이름이 숫자인지 확인
		if info.IsDir() && isNumeric(filepath.Base(path)) {
			n, err := strconv.Atoi(filepath.Base(path))
			if err != nil {
				return err
			}
			numericFolders = append(numericFolders, n)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	// 숫자 폴더 목록 정렬
	sort.Ints(numericFolders)

	return numericFolders, nil
}

// 문자열이 숫자인지 확인하는 함수
func isNumeric(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}
