package storage

import (
	"github.com/BurrrY/obstwiesen-server/internal/config"
	"github.com/BurrrY/obstwiesen-server/internal/file_store/disk"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func GetProvider() (*FileStorage, error) {
	var result FileStorage

	dbProvider := viper.GetString(config.FILE_PROVIDER)

	if dbProvider == "disk" {
		result = &disk.Thing
	} else {
		log.WithFields(log.Fields{
			"connectionData": dbProvider,
		}).Fatal("Cannot Connect to FILE_PROVIDER, No 'FILE_PROVIDER' configured!")
	}

	result.Setup()
	return &result, nil
}
