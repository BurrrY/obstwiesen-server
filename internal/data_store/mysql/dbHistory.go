package mysql

import (
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
	"os"
	"time"
)

var history = [...]string{
	"CREATE DATABASE `meadow`",
	"USE `meadow`;",

	"CREATE TABLE IF NOT EXISTS `events` (  `id` varchar(21) NOT NULL,  `parent_id` varchar(21) NOT NULL,  `title` varchar(255) DEFAULT NULL,  `description` text DEFAULT NULL,  `timestamp` datetime NOT NULL,  PRIMARY KEY (`id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;",
	"CREATE TABLE IF NOT EXISTS `meadows` (  `id` varchar(21) NOT NULL,  `name` varchar(255) NOT NULL,  PRIMARY KEY (`id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;",
	"CREATE TABLE IF NOT EXISTS `trees` (  `id` varchar(21) NOT NULL,  `name` varchar(255) NOT NULL,  `meadow_id` varchar(21) NOT NULL,  `created_at` datetime DEFAULT NULL,  `lat` decimal(10,7) DEFAULT NULL,  `lang` decimal(10,7) DEFAULT NULL,  PRIMARY KEY (`id`),  KEY `FK_trees_meadows` (`meadow_id`),  CONSTRAINT `FK_trees_meadows` FOREIGN KEY (`meadow_id`) REFERENCES `meadows` (`id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;",
	"CREATE TABLE IF NOT EXISTS `version` (  `version` int(11) DEFAULT 0,  `timestamp` datetime DEFAULT current_timestamp()) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;",
	"ALTER TABLE `events` ADD COLUMN `files` TEXT NULL AFTER `timestamp`;",
}

func updateDb() {
	current_version := 0

	rows, err := db.Query("USE `meadow`;")
	if err != nil {
		log.Warning("Error on getting Version!")
		if err.Error() == "Error 1049: Unknown database 'meadow'" {
			log.Warning("DB needs init")
		} else {
			log.Fatalln(err.Error())
		}
	} else {
		db, err = sqlx.Connect("mysql", os.Getenv("CON_STR")+"meadow")
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
				current_version = version
			}
		}
	}

	i := current_version
	for ; i < len(history); i++ {
		log.Warning("DB-Update: ", history[i])
		_, err = db.Exec(history[i])
		if err != nil {
			log.Error("DB-Update-Err: ", err)
			return
		}
	}

	if i > current_version {
		log.Info("Updated DB to version ", i)
		_, err = db.Exec("INSERT INTO version (version, timestamp) VALUES (?, ?)", i, time.Now().Format("2006-01-02T15:04:05"))
		if err != nil {
			log.Error("INSERT: ", err)
		}
	}
}
