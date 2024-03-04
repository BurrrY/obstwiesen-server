package storage

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/BurrrY/obstwiesen-server/graph/model"
)

type FileStorage interface {
	GetType() string
	StoreFile(file *graphql.Upload, parentID string, fileID string) error
	GetFiles(parentId string) ([]*model.File, error)
}
