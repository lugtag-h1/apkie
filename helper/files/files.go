package files

import (
	"archive/zip"
	"os"
)

func Exists(filename string) bool {
	fi, err := os.Stat(filename)
	if err != nil {
		return false
	}

	return !fi.IsDir() && fi.Size() > 0
}

func IsValidApk(filename string) bool {
	rd, err := zip.OpenReader(filename)
	if err == nil {
		defer rd.Close()
	}

	return err == nil
}
