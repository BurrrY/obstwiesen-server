package graph

import (
	"github.com/BurrrY/obstwiesen-server/graph/model"
	str "github.com/BurrrY/obstwiesen-server/internal/data_store"
	fstr "github.com/BurrrY/obstwiesen-server/internal/file_store"
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

func init() {

}
