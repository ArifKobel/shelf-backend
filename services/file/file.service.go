package file_service

import (
	"os"

	"github.com/google/uuid"
)

func GenerateFileName(name string) string {
	uuid := uuid.New()
	return uuid.String() + "-" + name
}

func CreatePathIfNotExists(path string) error {
	exist, err := exists(path)
	if err != nil {
		return err
	}
	if !exist {
		err := os.MkdirAll(path, 0755)
		if err != nil {
			return err
		}
	}
	return nil
}

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
