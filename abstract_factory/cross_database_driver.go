// --- products.go ---
package abstractfactory

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
)

// Abstract products
type Connection interface {
	Open(dsn string) (*sql.DB, error)
}
type QueryBuilder interface {
	Select(table string, cols ...string) string
}
type MigrationRunner interface {
	Run(db *sql.DB, migrationsPath string) error
}

// Abstract factory
type DBFactory interface {
	NewConnection() Connection
	NewQueryBuilder() QueryBuilder
	NewMigrationRunner() MigrationRunner
}

type mysqlConn struct{}

func (c *mysqlConn) Open(dsn string) (*sql.DB, error) {
	return sql.Open("mysql", dsn)
}

type mysqlQB struct{}

func (q *mysqlQB) Select(table string, cols ...string) string {
	return fmt.Sprintf("SELECT %s FROM `%s`", strings.Join(cols, ", "), table)
}

type mysqlMigrate struct{}

func (m *mysqlMigrate) Run(db *sql.DB, p string) error {
	// run .sql files in p in MySQL‑specific way…
	return nil
}

type MySQLFactory struct{}

func (MySQLFactory) NewConnection() Connection           { return &mysqlConn{} }
func (MySQLFactory) NewQueryBuilder() QueryBuilder       { return &mysqlQB{} }
func (MySQLFactory) NewMigrationRunner() MigrationRunner { return &mysqlMigrate{} }

func run(factory DBFactory, dsn string) {
	conn := factory.NewConnection()
	dbConn, err := conn.Open(dsn)
	if err != nil {
		log.Fatal(err)
	}

	qb := factory.NewQueryBuilder()
	sql := qb.Select("users", "id", "name")

	mig := factory.NewMigrationRunner()
	if err := mig.Run(dbConn, "./migrations"); err != nil {
		log.Fatal(err)
	}

	log.Println("SQL:", sql)
}

func RunCrossDatabaseDriverDemo() {
	// switch on config
	var factory DBFactory = MySQLFactory{}
	// or: factory = postgres.PostgresFactory{}

	run(factory, "user:pass@tcp(localhost)/appdb")
}
