package graph

import (
	"github.com/BurrrY/obstwiesen-server/graph/model"
	str "github.com/BurrrY/obstwiesen-server/internal/data_store"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

var storage str.Storage

type Resolver struct {
	meadows []*model.Meadow
}

func init() {
	tmp, _ := str.GetProvider()
	storage = *tmp
}
