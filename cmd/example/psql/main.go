package main

import (
	"context"
	"fmt"
	"os"

	"github.com/HosseinForouzan/user-management/entity"
	"github.com/jackc/pgx/v5"
)



func main() {
	dbUrl := "postgres://hossein:secret@localhost:5431/user_management_db"
	conn, err := pgx.Connect(context.Background(), dbUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("CONNECTED!!!")
	defer conn.Close(context.Background())

	var user entity.User

	err = conn.QueryRow(context.Background(), "SELECT * FROM users WHERE id=$1", 1).
	Scan(&user.ID, &user.Name, &user.PhoneNumber, &user.Email, &user.Password)
	if err != nil {
	fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
	os.Exit(1)
	}

	fmt.Println(user)
}