package factory

import "database/sql"

type DBConfig struct {
	Driver string
	DSN    string // Data Source Name
}

func OpenDatabase(cfg DBConfig) (*sql.DB, error) {
	db, err := sql.Open(cfg.Driver, cfg.DSN)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func RunDatabaseFactoryDemo() {
	cfg := DBConfig{
		Driver: "mysql",
		DSN:    "user:password@tcp(localhost:3306)/dbname",
	}

	db, err := OpenDatabase(cfg)
	if err != nil {
		println("Error opening database:", err.Error())
		return
	}
	defer db.Close()

	// cfg = DBConfig{Driver: "postgres", DSN: "user=foo dbname=bar sslmode=disable"}
	// dbConn, err := OpenDatabase(cfg)
	// if err != nil {
	// 	println("Error opening database:", err.Error())
	// 	return
	// }
	// defer dbConn.Close()
}
