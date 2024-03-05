package storage

import (
	"github.com/BurrrY/obstwiesen-server/internal/config"
	"github.com/BurrrY/obstwiesen-server/internal/data_store/mysql"
	"github.com/BurrrY/obstwiesen-server/internal/data_store/sqlite"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func GetProvider() (*Storage, error) {
	var result Storage

	dbProvider := viper.GetString(config.DB_PROVIDER)

	if dbProvider == "sqlite" {
		result = &sqlite.Connection
	} else if dbProvider == "mysql" {
		result = &mysql.Connection
	} else {
		log.WithFields(log.Fields{
			"connectionData": dbProvider,
		}).Fatal("Cannot Connect to any DB, No 'dbProvider' configured!")
	}
	result.Setup()
	return &result, nil
}
