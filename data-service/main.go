package main

import (
	"github.com/atyagi9006/certificationapp/data-service/src/service"
)

func main() {
	service.DbInit()
	service.StartServer("50054")
}



