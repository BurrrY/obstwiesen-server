package mysql

import (
	"fmt"
	"github.com/BurrrY/obstwiesen-server/internal/config"
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"strings"
	"time"
)

var history = [...]string{
	"CREATE DATABASE `_DB_NAME_`",
	"USE `_DB_NAME_`;",

	"CREATE TABLE IF NOT EXISTS `events` (  `id` varchar(21) NOT NULL,  `parent_id` varchar(21) NOT NULL,  `title` varchar(255) DEFAULT NULL,  `description` text DEFAULT NULL,  `timestamp` datetime NOT NULL,  PRIMARY KEY (`id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;",
	"CREATE TABLE IF NOT EXISTS `meadows` (  `id` varchar(21) NOT NULL,  `name` varchar(255) NOT NULL,  PRIMARY KEY (`id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;",
	"CREATE TABLE IF NOT EXISTS `trees` (  `id` varchar(21) NOT NULL,  `name` varchar(255) NOT NULL,  `meadow_id` varchar(21) NOT NULL,  `created_at` datetime DEFAULT NULL,  `lat` decimal(10,7) DEFAULT NULL,  `lang` decimal(10,7) DEFAULT NULL,  PRIMARY KEY (`id`),  KEY `FK_trees_meadows` (`meadow_id`),  CONSTRAINT `FK_trees_meadows` FOREIGN KEY (`meadow_id`) REFERENCES `meadows` (`id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;",
	"CREATE TABLE IF NOT EXISTS `version` (  `version` int(11) DEFAULT 0,  `timestamp` datetime DEFAULT current_timestamp()) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;",
	"ALTER TABLE `meadows` ADD COLUMN `area` JSON NULL AFTER `name`;",
}

func updateDb() {
	currentVersion := 0

	rows, err := db.Query(fmt.Sprintf("USE `%s`;", viper.GetString(config.DB_NAME)))
	if err != nil {

		log.Warning("Error on getting Version!")

		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			if mysqlErr.Number == 1049 { //Unknown database 'xxx'
				log.Warning("DB needs init")
			} else {
				log.Fatal("MYSQL-Error connecting to DB", err.Error())
			}
		} else {
			log.Fatal("Error connecting to DB", err.Error())
		}
	} else {
		//re-connect to correct database
		db, err = sqlx.Connect("mysql", viper.GetString(config.DB_CONNSTR)+viper.GetString(config.DB_NAME))
		if err != nil {
			log.Fatalln(err)
		}

		rows, err = db.Query("SELECT version FROM version ORDER BY version DESC LIMIT 1")
		if err != nil {
			log.Warning("Error on getting Version!")
			if err.Error() == "no such table: version" {
				log.Warning("DB needs init")
			}
		} else {
			for rows.Next() {
				var version int
				err = rows.Scan(&version)
				if err != nil {
					log.Warning("version-Err: ", err)
				}
				log.Info("DB-Version: ", version)
				currentVersion = version
			}
		}
	}

	i := currentVersion
	for ; i < len(history); i++ {
		cmd := strings.Replace(history[i], "_DB_NAME_", viper.GetString(config.DB_NAME), -1)
		log.Warning("DB-Update: ", cmd)
		_, err = db.Exec(cmd)
		if err != nil {
			log.Error("DB-Update-Err: ", err)
			return
		}
	}

	if i > currentVersion {
		log.Info("Updated DB to version ", i)
		_, err = db.Exec("INSERT INTO version (version, timestamp) VALUES (?, ?)", i, time.Now().Format("2006-01-02T15:04:05"))
		if err != nil {
			log.Error("INSERT: ", err)
		}
	}
}
