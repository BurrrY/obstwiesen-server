package mysql

import (
	"context"
	"encoding/json"
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

func (m stor) UpdateMeadow(ctx context.Context, id string, input model.MeadowInput) (*model.Meadow, error) {

	jsondata, _ := json.Marshal(input.Area)

	jsonStr := string(jsondata)

	_, err := db.Exec("UPDATE meadows SET name = ?, area = ? WHERE id = ?", input.Name, jsonStr, id)
	if err != nil {
		log.Warning("Error in UpdateMeadow: ", err)
		return nil, err
	}

	return m.GetMeadowByID(id)
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

func (m stor) UpdateTree(id string, input model.TreeInput) (*model.Tree, error) {
	_, err := db.Exec("UPDATE trees SET name = ?, lat = ?, lang = ? WHERE id = ?",
		input.Name, input.Lat, input.Lang, id)

	if err != nil {
		log.Warning("Error in AddTree: ", err)
		return nil, err
	}

	return m.GetTreeByID(id)
}

func (m stor) GetTreesOfMeadow(id string) ([]*model.Tree, error) {

	trees := []*model.Tree{}
	err := db.Select(&trees, "SELECT id, name, lang, lat FROM trees WHERE meadow_id = ? ORDER BY name", id)

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

	var data []*model.Event
	err := db.Select(&data, "SELECT id, title, description, timestamp FROM events WHERE parent_id = ? ORDER BY timestamp DESC", id)

	return data, err
}

func (m stor) GetMeadows() ([]*model.Meadow, error) {
	var meadows []*model.Meadow

	var raw []*GetMeadow
	err := db.Select(&raw, "SELECT id, name, area FROM meadows ORDER BY name")

	for _, meadow := range raw {
		meadows = append(meadows, toMeadowModel(meadow))
	}

	return meadows, err
}

func toMeadowModel(meadow *GetMeadow) *model.Meadow {

	data := &model.Meadow{
		ID:   meadow.ID,
		Name: meadow.Name,
		Area: nil,
	}
	if meadow.Area != nil {
		jsonString := *meadow.Area
		if len(jsonString) > 0 {
			err := json.Unmarshal([]byte(jsonString), &data.Area)
			if err != nil {
				return nil
			}
		}
	}
	return data

}

func (m stor) GetMeadowByID(id string) (*model.Meadow, error) {

	raw := &GetMeadow{}

	err := db.Get(raw, "SELECT id, name, area FROM meadows WHERE id = ?", id)
	if err != nil {
		log.Warning("GetMeadowByID", err)
	}

	return toMeadowModel(raw), err
}
