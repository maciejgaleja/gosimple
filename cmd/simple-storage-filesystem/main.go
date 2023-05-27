package main

import (
	"fmt"

	"github.com/maciejgaleja/gosimple/pkg/authenticator"
	"github.com/maciejgaleja/gosimple/pkg/hash"
	"github.com/maciejgaleja/gosimple/pkg/keyvalue/json"
)

const (
	dbFile = "./workspace/authenticator.json"
)

func NewAuthenticator() (*authenticator.Authenticator, error) {
	db, err := json.NewJsonStore(dbFile)
	if err != nil {
		return nil, err
	}
	return authenticator.NewAuthenticator(db)
}

func main() {
	fmt.Println("Hello, world")
	a, err := NewAuthenticator()
	if err != nil {
		panic(err)
	}

	err = a.Register(authenticator.Entry{authenticator.HashedUsername(hash.HashString("user")), authenticator.HashedPassword(hash.HashString("password"))})

	b, err := a.Verify(authenticator.Entry{authenticator.HashedUsername(hash.HashString("usesr")), authenticator.HashedPassword(hash.HashString("password"))})
	if err != nil {
		panic(err)
	}
	fmt.Println(b)
}
