package sqlite

import (
	log "github.com/sirupsen/logrus"
	"time"
)

var history = [...]string{
	"create table version ( version integer default 0, timestamp TEXT default TIMESTAMP);",
}

func updateDb() {
	current_version := 0
	rows, err := db.Query("SELECT version FROM version ORDER BY version DESC LIMIT 1")
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

	i := current_version
	for ; i < len(history); i++ {
		log.Warning("DB-Update: ", history[0])
		_, err = db.Exec(history[0])
		if err != nil {
			log.Warning("DB-Update-Err: ", err)
		}
	}
	t := time.Now()
	_, err = db.Exec("INSERT INTO version (version, timestamp) VALUES (?, ?)", i, t.Format(time.DateTime))
	if err != nil {
		log.Warning("INSERT: ", err)
	}
}
