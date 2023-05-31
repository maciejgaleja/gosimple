package hash

import (
	"bytes"
	"crypto/sha512"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/maciejgaleja/gosimple/pkg/types"
)

type Sha512 []byte

func Hash(r io.Reader) Sha512 {
	h := sha512.New()
	_, err := io.Copy(h, r)
	if err != nil {
		panic(err)
	}
	return Sha512(fmt.Sprintf("%x", h.Sum(nil)))
}

func HashBytes(d []byte) Sha512 {
	return Hash(bytes.NewBuffer(d))
}

func HashString(s string) Sha512 {
	r := strings.NewReader(s)
	return Hash(r)
}

func HashFile(f types.FilePath) (Sha512, error) {
	fd, err := os.Open(string(f))
	if err != nil {
		return Sha512{}, err
	}

	return Hash(fd), nil
}
