package main

import (
	"github.com/gin-gonic/gin"
	"github.com/maciejgaleja/gosimple/pkg/authenticator"
	"github.com/maciejgaleja/gosimple/pkg/keyvalue/json"
	"github.com/maciejgaleja/gosimple/pkg/server/storage"
	"github.com/maciejgaleja/gosimple/pkg/storage/filesystem"
)

const (
	dbFile     = "./workspace/authenticator.json"
	storageDir = "./workspace/storage/"
)

func NewAuthenticator() (*authenticator.Authenticator, error) {
	db, err := json.NewJsonStore(dbFile)
	if err != nil {
		return nil, err
	}
	return authenticator.NewAuthenticator(db)
}

func main() {
	s, err := filesystem.NewFilesystemStore(storageDir)
	if err != nil {
		panic(err)
	}

	ss := storage.NewStorageApi(s)

	g := gin.Default()
	ss.RegisterToGin(g, "objects")
	g.Run("0.0.0.0:8080")
}
