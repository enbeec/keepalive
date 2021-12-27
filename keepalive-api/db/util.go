package db

import "os"

func sameFilesystem(pathA, pathB string) bool {
	infoA, errA := os.Stat(pathA)
	infoB, errB := os.Stat(pathB)

	if errA != nil && errB != nil {
		return false
	} else if infoA.Sys() != infoB.Sys() {
		return false
	}

	return true
}

func dirMustExist(path string, mode os.FileMode) error {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		os.Mkdir(path, mode) // create if doesn't exist
	} else if info.Mode() != mode {
		os.Chmod(path, mode) // correct wrong FileMode
	} else {
		return err
	}

	return nil
}
