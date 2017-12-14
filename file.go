package main

import (
	"os"
)

func CreateDir(path string) error {
	stat, err := os.Stat(path)

	if os.IsNotExist(err) || !stat.IsDir() {
		err = os.MkdirAll(path, 0755)

		return err
	}

	return err
}
