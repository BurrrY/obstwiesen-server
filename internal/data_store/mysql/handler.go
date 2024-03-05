package mysql

import (
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
	"os"
)

type stor struct {
	ConnectionString string
	ConnectionError  error
}

var db *sqlx.DB
var Connection stor

func (m stor) GetType() string {
	return "sqlite"
}

func init() {

	if os.Getenv("DB_PROVIDER") != "mysql" {
		log.New().Info("Skip mysql Init by Config: " + os.Getenv("DB_PROVIDER"))
		Connection.ConnectionError = errors.New("mysql disabled")
		return
	}

	var err error

	db, err = sqlx.Connect("mysql", os.Getenv("DB_CONNSTR"))
	if err != nil {
		log.Fatalln(err)
	}

	updateDb()

}
