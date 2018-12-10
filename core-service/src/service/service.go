package service

import (
	"log"
	"net/http"

	"github.com/atyagi9006/certificationapp/core-service/src/cache"
	"github.com/atyagi9006/certificationapp/core-service/src/config"
	"github.com/atyagi9006/certificationapp/core-service/src/router"
)

func InitCache() {
	db := config.Init()
	log.Println(db.DBConfig.DatabaseName)
	cache.Init(db)
}

func StartServer(port string) {
	r := router.NewRouter()
	// 1. router can be passed to http.handler
	http.Handle("/", r)
	log.Println("Starting HTTP service at : ", port)
	err := http.ListenAndServe(":"+port, nil) //router can also passed here
	if err != nil {
		log.Println("An error occured starting HTTP listener at port " + port)
		log.Println("Error: " + err.Error())
	}
}
