package psql

import (
	"context"
	"fmt"
	"os"

	pgx "github.com/jackc/pgx/v5/pgxpool"
)

type PsqlDB struct {
	db *pgx.Pool
}

func (p *PsqlDB) Conn() *pgx.Pool {
	return p.db
}

func New() *PsqlDB {
	dbUrl := "postgres://hossein:secret@localhost:5431/user_management_db"
	db, err := pgx.New(context.Background(), dbUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	

	return &PsqlDB{db: db}
}