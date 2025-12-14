package psqluser

import "github.com/HosseinForouzan/user-management/repository/migrations/psql"

type DB struct {
	conn *psql.PsqlDB
}

func New(conn *psql.PsqlDB) *DB {
	return &DB{conn: conn}
}