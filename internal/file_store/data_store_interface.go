package storage

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/BurrrY/obstwiesen-server/graph/model"
)

type FileStorage interface {
	GetType() string
	Setup()
	StoreFile(file *graphql.Upload, parentID string, fileID string) (error, string)
	GetFiles(parentId string) ([]*model.File, error)
	GetImage(file string, dir string, width int) string
	GetFile(id string) (*model.File, error)
}
