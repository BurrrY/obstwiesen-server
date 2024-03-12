package graph

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/BurrrY/obstwiesen-server/graph/model"
	str "github.com/BurrrY/obstwiesen-server/internal/data_store"
	fstr "github.com/BurrrY/obstwiesen-server/internal/file_store"
	gonanoid "github.com/matoous/go-nanoid/v2"
	log "github.com/sirupsen/logrus"
	"strings"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

var storage str.Storage
var filestore fstr.FileStorage

type Resolver struct {
	meadows []*model.Meadow
}

func (r *Resolver) Setup() {
	tmp, _ := str.GetProvider()
	storage = *tmp

	tmp2, _ := fstr.GetProvider()
	filestore = *tmp2
}

func storeFile(file *graphql.Upload, eventID string) (error, string) {
	fileID, _ := gonanoid.New()

	idx := strings.LastIndex(file.Filename, ".")
	ending := file.Filename[idx:]
	log.Debug("Filename: " + file.Filename)
	log.Debug("End: ", ending)
	err, newPath := filestore.StoreFile(file, eventID, fileID+ending)
	return err, newPath
}

func init() {

}
