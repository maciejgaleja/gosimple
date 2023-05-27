package storage

import (
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/maciejgaleja/gosimple/pkg/storage"
)

type PostResponse struct {
	Key  storage.Key
	Size int64
}

type GetResponse struct {
	Key  storage.Key
	Size int64
}

type DeleteResponse struct {
	Key storage.Key
}

type InfoResponse struct {
	Key    storage.Key
	Exists bool
}

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
	g.GET(buildPath(rootPath, ":key", "info"), func(c *gin.Context) {
		k := storage.Key(c.Param("key"))
		e := s.s.Exists(k)
		c.JSON(http.StatusOK, InfoResponse{Key: k, Exists: e})
	})
	g.GET(buildPath(rootPath, ":key"), func(c *gin.Context) {
		k := storage.Key(c.Param("key"))

		r, err := s.s.Reader(k)
		if err != nil {
			respondWithError(c, err)
			return
		}

		n, err := io.Copy(c.Writer, r)
		if err != nil {
			respondWithError(c, err)
			return
		}

		c.JSON(http.StatusOK, GetResponse{Key: k, Size: n})
	})
	g.POST(buildPath(rootPath, ":key"), func(c *gin.Context) {
		k := storage.Key(c.Param("key"))
		file, err := c.FormFile("file")
		if err != nil {
			c.String(http.StatusBadRequest, "get form err: %s", err.Error())
			return
		}

		fd, err := file.Open()
		if err != nil {
			respondWithError(c, err)
			return
		}

		w, err := s.s.Create(k)
		if err != nil {
			respondWithError(c, err)
			return
		}

		n, err := io.Copy(w, fd)
		if err != nil {
			respondWithError(c, err)
			return
		}

		c.JSON(http.StatusCreated, PostResponse{Key: k, Size: n})
	})
	g.DELETE(buildPath(rootPath, ":key"), func(c *gin.Context) {
		k := storage.Key(c.Param("key"))

		err := s.s.Delete(k)
		if err != nil {
			respondWithError(c, err)
			return
		}

		c.JSON(http.StatusCreated, DeleteResponse{Key: k})
	})
}

func respondWithError(c *gin.Context, e error) {
	c.JSON(http.StatusInternalServerError, map[string]string{"error": e.Error()})
}

func buildPath(es ...string) string {
	return "/" + strings.Join(es, "/")
}
