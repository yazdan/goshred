package internal

import (
	"testing"

	"github.com/spf13/afero"
)

func TestGetFileSizeWhenFileNotExist(t *testing.T) {
	var fs = afero.NewMemMapFs()

	size, err := GetFileSize("a.txt", fs)

	if size != -1 || err == nil {
		t.Fatalf("When file dose not exist, GetFileSize should return error ")
	}
}

func TestGetFileSizeWhenFileExists(t *testing.T) {
	var fs = afero.NewMemMapFs()

	var data = []byte("file a")
	afero.WriteFile(fs, "a.txt", data, 0644)

	size, err := GetFileSize("a.txt", fs)

	if size != int64(len(data)) || err != nil {
		t.Fatalf("When file exists, GetFileSize should not return error ")
	}
}

func TestShredWhenFileNotExist(t *testing.T) {
	var fs = afero.NewMemMapFs()

	res := Shred("a.txt", fs)

	if res {
		t.Fatalf("When file dose not exist, Shred should not return succesfull ")
	}
}

func TestShredWhenFileExist(t *testing.T) {
	var fs = afero.NewMemMapFs()

	var data = []byte("file a")
	afero.WriteFile(fs, "a.txt", data, 0644)

	res := Shred("a.txt", fs)

	if !res {
		t.Fatalf("When file dose not exist, Shred should not return succesfull ")
	}

	exist, _ := afero.Exists(fs, "a.txt")

	if exist {
		t.Fatalf("The file shouln't exist after Shred")
	}
}
