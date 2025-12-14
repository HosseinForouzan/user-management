package main

import (
	"fmt"

	"github.com/HosseinForouzan/user-management/repository/psql"
	"github.com/HosseinForouzan/user-management/repository/psql/psqluser"
)

func main() {
	psqlrepo := psql.New()
	psqluser := psqluser.New(psqlrepo)

	defer psqlrepo.Conn().Close()


	fmt.Println(psqluser.GetUserByID(1))


}