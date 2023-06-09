package files

import (
	"bytes"
	"io"
	"os"
)

func LineCounter(r io.Reader) (int, error) {
	buf := make([]byte, 32*1024)
	count := 0
	lineSep := []byte{'\n'}

	for {
		c, err := r.Read(buf)
		count += bytes.Count(buf[:c], lineSep)

		switch {
		case err == io.EOF:
			return count, nil

		case err != nil:
			return count, err
		}
	}
}

func ResetFile(file *os.File) {
	_, err := file.Seek(0, io.SeekStart)
	if err != nil {
		panic(err)
	}
}

func GetWorkdir() string {
	ex, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return ex
}
