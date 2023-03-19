package internal

import (
	"crypto/rand"
	"log"
	"os"

	"github.com/spf13/afero"
)

func GetFileSize(name string, fs afero.Fs) (int64, error) {
	f, err := fs.Open(name)
	var size = int64(-1)

	if err != nil {
		log.Print(err)

		return -1, err
	}

	defer f.Close()

	stat, err := f.Stat()
	if err != nil {
		log.Print(err)
	} else {
		size = stat.Size()
	}

	return size, err
}

func Shred(name string, fs afero.Fs) bool {
	size, err := GetFileSize(name, fs)
	result := true

	if err == nil {
		for i := 0; i < 3; i++ {
			if !random_overwrite(name, size, fs) {
				result = false
				break
			}
		}

		err = fs.Remove(name)

		if err != nil {
			log.Print(err)
			result = false
		}
	} else {
		result = false
	}

	return result
}

func random_overwrite(name string, size int64, fs afero.Fs) bool {
	// This will truncate the file
	result := false
	f, err := fs.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		log.Print(err)
	}

	defer f.Close()

	_, err = f.Write(random_bytes(size))

	if err != nil {
		log.Print(err)
	}
	result = true

	return result
}

// Note: This can exhaust ram if file size is too big
// improvement point
func random_bytes(size int64) []byte {
	buf := make([]byte, size)
	_, err := rand.Read(buf)
	if err != nil {
		log.Printf("error while generating random string: %s", err)
	}
	return buf
}
