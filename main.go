package main

import (

	"github.com/HosseinForouzan/user-management/delivery/httpserver"
	"github.com/HosseinForouzan/user-management/repository/psql"
	"github.com/HosseinForouzan/user-management/repository/psql/psqluser"
	"github.com/HosseinForouzan/user-management/service/userservice"
)


func main() {


	userSvc := setupServices()

	server := httpserver.New(userSvc)
	server.SetRoutes()
	
}

func setupServices() userservice.Service {
	psqlrepo := psql.New()
	psqluser := psqluser.New(psqlrepo)

	userSvc := userservice.New(psqluser)

	return userSvc

}
