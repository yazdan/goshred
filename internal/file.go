package internal

import (
	"crypto/rand"
	"log"
	"os"
)

func GetFileSize(name string) int64 {
	f, err := os.Open(name)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	stat, err := f.Stat()
	if err != nil {
		log.Fatal(err)
	}

	var size = stat.Size()

	return size
}

func Shred(name string) bool {
	size := GetFileSize(name)
	result := true

	for i := 0; i < 3; i++ {
		if !random_overwrite(name, size) {
			result = true
			break
		}
	}

	return result
}

func random_overwrite(name string, size int64) bool {
	// This will truncate the file
	result := false
	f, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	_, err = f.Write(random_bytes(size))

	if err != nil {
		log.Fatal(err)
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
		log.Fatalf("error while generating random string: %s", err)
	}
	return buf
}
