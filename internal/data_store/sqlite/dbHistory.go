package sqlite

import (
	log "github.com/sirupsen/logrus"
	"time"
)

var history = [...]string{
	"create table version ( version integer default 0, timestamp TEXT default TIMESTAMP);",
	"create table meadows(  id   TEXT not null constraint meadows_pk        primary key, name text not null);",
	`create table trees
		(
			id         TEXT
				constraint trees_pk
					primary key,
			name       TEXT,
			meadow_id  TEXT
				constraint trees_meadows_id_fk
					references meadows,
			created_at TEXT
		);
`,
	`alter table trees
    add lat REAL;

alter table trees
    add lang REAL;

`,
	`create table events
(
    id      TEXT
        constraint events_pk
            primary key,
    parent_id TEXT
);

`,
	`alter table events
    add title TEXT;

alter table events
    add description TEXT;

alter table events
    add timestamp TEXT;

`,
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
		log.Warning("DB-Update: ", history[i])
		_, err = db.Exec(history[i])
		if err != nil {
			log.Error("DB-Update-Err: ", err)
			return
		}
	}

	_, err = db.Exec("INSERT INTO version (version, timestamp) VALUES (?, ?)", i, time.Now().Format(time.RFC3339))
	if err != nil {
		log.Error("INSERT: ", err)
	}
}
