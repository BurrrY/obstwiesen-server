package mysql

import (
	"github.com/BurrrY/obstwiesen-server/graph/model"
	log "github.com/sirupsen/logrus"
	"time"
)

func (m stor) StoreMeadow(meadow *model.Meadow) {
	_, err := db.Exec("INSERT INTO meadows ( id, name) VALUES (?, ?)", meadow.ID, meadow.Name)
	if err != nil {
		log.Warning("Error in StoreMeadow: ", err)
		return
	}

}
func (m stor) AddEvent(elemnt *model.Event, id string) error {
	_, err := db.Exec("INSERT INTO events ( id, parent_id, description, title, timestamp) VALUES (?, ?,?,?,?)",
		elemnt.ID, id, elemnt.Description, elemnt.Title, elemnt.Timestamp)
	if err != nil {
		log.Warning("Error in AddEvent: ", err)
		return err
	}

	return nil
}

func (m stor) AddTree(tree *model.Tree, id string) {
	_, err := db.Exec("INSERT INTO trees ( id, name, meadow_id, created_at) VALUES (?, ?, ?, ?)",
		tree.ID, tree.Name, id, time.Now().Format("2006-01-02T15:04:05"))
	if err != nil {
		log.Warning("Error in AddTree: ", err)
		return
	}

}

func (m stor) GetTreesOfMeadow(id string) ([]*model.Tree, error) {

	trees := []*model.Tree{}
	err := db.Select(&trees, "SELECT id, name FROM trees WHERE meadow_id = ? ORDER BY name", id)

	return trees, err
}

func (m stor) GetTreeByID(id string) (*model.Tree, error) {
	d := model.Tree{}
	err := db.Get(&d, "SELECT id, name, lang, lat FROM trees WHERE id = ?", id)
	if err != nil {
		log.Warning("GetTreeByID", err)
	}

	log.Info("Tree!", d)
	return &d, err
}

func (m stor) GetEventsOfTree(id string) ([]*model.Event, error) {

	data := []*model.Event{}
	err := db.Select(&data, "SELECT id, title, description, timestamp FROM events WHERE parent_id = ? ORDER BY timestamp DESC", id)

	return data, err
}

func (m stor) GetMeadows() ([]*model.Meadow, error) {
	meadows := []*model.Meadow{}
	err := db.Select(&meadows, "SELECT id, name FROM meadows ORDER BY name")

	return meadows, err
}

func (m stor) GetMeadowByID(id string) (*model.Meadow, error) {
	meadow := model.Meadow{}
	err := db.Get(&meadow, "SELECT id, name FROM meadows WHERE id = ?", id)
	if err != nil {
		log.Warning("GetMeadowByID", err)
	}

	return &meadow, err
}
