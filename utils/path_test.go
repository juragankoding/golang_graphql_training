package utils

import (
	"os"
	"testing"
)

func TestRootDir(t *testing.T) {
	path := RootDir()

	if path == "" {
		t.Error("this path is null")
	}

	file, err := os.Open(path)

	if err != nil {
		t.Error(err.Error())
	}

	defer file.Close()

	fi, err := file.Stat()

	if err != nil {
		t.Error(err.Error())
	}

	if fi.IsDir() == false {
		t.Error("file is not dir")
	}

}
