package storage

import "github.com/BurrrY/obstwiesen-server/graph/model"

type Storage interface {
	GetType() string
	StoreMeadow(meadow *model.Meadow)
	AddTree(tree *model.Tree, id string)
	GetMeadows() ([]*model.Meadow, error)
	GetMeadowByID(id string) (*model.Meadow, error)
	GetTreesOfMeadow(id string) ([]*model.Tree, error)
	GetTreeByID(id string) (*model.Tree, error)
	AddEvent(elemnt *model.Event, id string) error
	GetEventsOfTree(id string) ([]*model.Event, error)
}
