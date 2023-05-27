package storage

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/maciejgaleja/gosimple/pkg/storage"
)

type StorageApi struct {
	s storage.Storage
}

func NewStorageApi(s storage.Storage) *StorageApi {
	return &StorageApi{s: s}
}

func (s *StorageApi) RegisterToGin(g *gin.Engine, rootPath string) {
	g.GET(buildPath(rootPath, "/"), func(c *gin.Context) {
		l, err := s.s.List()
		if err != nil {
			respondWithError(c, err)
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"objects": l,
		})
	})
}

func respondWithError(c *gin.Context, e error) {
	c.JSON(http.StatusInternalServerError, map[string]string{"error": e.Error()})
}

func buildPath(es ...string) string {
	return "/" + strings.Join(es, "/")
}
