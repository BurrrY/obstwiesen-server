package mysql

import (
	"errors"
	"github.com/BurrrY/obstwiesen-server/internal/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type stor struct {
	ConnectionString string
	ConnectionError  error
	SetupDone        bool
}

func (m stor) Setup() {

	if m.SetupDone == true {
		return
	}

	if viper.GetString(config.DB_PROVIDER) != "mysql" {
		log.New().Info("Skip mysql Init by Config: " + viper.GetString(config.DB_PROVIDER))
		Connection.ConnectionError = errors.New("mysql disabled")
		return
	}

	var err error

	db, err = sqlx.Connect("mysql", viper.GetString(config.DB_CONNSTR))
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
