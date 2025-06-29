package singleton

import (
	"database/sql"
	"sync"

	_ "github.com/lib/pq"
)

var (
	dbInstance *sql.DB
	onceDB     sync.Once
)

func GetDB(dsn string) (*sql.DB, error) {
	var err error
	onceDB.Do(func() {
		dbInstance, err = sql.Open("postgres", dsn)
		if err != nil {
			return
		}
		// optional: configure pool settings
		dbInstance.SetMaxOpenConns(25)
		dbInstance.SetMaxIdleConns(5)
		err = dbInstance.Ping()
	})
	return dbInstance, err
}
