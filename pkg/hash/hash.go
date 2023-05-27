package hash

import (
	"crypto/sha512"
	"fmt"
	"io"
	"strings"
)

type Hash []byte

func HashString(s string) Hash {
	h := sha512.New()
	r := strings.NewReader(s)
	_, err := io.Copy(h, r)
	if err != nil {
		panic(err)
	}
	return Hash(fmt.Sprintf("%x", h.Sum(nil)))
}
