package psql

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

type PsqlDB struct {
	db *pgx.Conn
}

func (p *PsqlDB) Conn() *pgx.Conn {
	return p.db
}

func New() *PsqlDB {
	dbUrl := "postgres://hossein:secret@localhost:5431/user_management_db"
	db, err := pgx.Connect(context.Background(), dbUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("CONNECTED!!!")
	defer db.Close(context.Background())

	return &PsqlDB{db: db}
}