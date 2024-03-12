package sqlite

import (
	"errors"
	"github.com/BurrrY/obstwiesen-server/graph/model"
	"github.com/BurrrY/obstwiesen-server/internal/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type stor struct {
	ConnectionString string
	ConnectionError  error
	SetupDone        bool
}

func (m stor) UpdateTree(id string, input model.TreeInput) (*model.Tree, error) {
	//TODO implement me
	panic("implement me")
}

func (m stor) Setup() {

	if m.SetupDone == true {
		return
	}

	if viper.GetString(config.DB_PROVIDER) != "sqlite" {
		log.New().Info("Skip sqlite Init by Config: " + viper.GetString(config.DB_PROVIDER))
		Connection.ConnectionError = errors.New("sqlite disabled")
		return
	}

	var err error

	db, err = sqlx.Connect("sqlite3", viper.GetString(config.DB_CONNSTR))
	if err != nil {
		log.Fatalln(err)
	}

	updateDb()
	m.SetupDone = true
}

var db *sqlx.DB
var Connection stor

func (m stor) GetType() string {
	return "sqlite"
}

func init() {

}
