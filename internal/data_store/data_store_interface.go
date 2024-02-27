package storage

import "github.com/BurrrY/obstwiesen-server/graph/model"

type Storage interface {
	GetType() string
	StoreMeadow(meadow *model.Meadow)
}
