package main

import (
	"fmt"

	"github.com/HosseinForouzan/user-management/repository/psql"
	"github.com/HosseinForouzan/user-management/repository/psql/psqluser"
	"github.com/HosseinForouzan/user-management/service/userservice"
)

func main() {
	psqlrepo := psql.New()
	psqluser := psqluser.New(psqlrepo)

	defer psqlrepo.Conn().Close()


	fmt.Println(psqluser.GetUserByID(1))

	userSvc := userservice.New(psqluser)

	s, _ := userSvc.Register(userservice.RegisterRequest{
		Name: "Majid",
		PhoneNumber: "0912",
		Email: "m@m",
		Password: "123",
	})

	fmt.Println(s)

}