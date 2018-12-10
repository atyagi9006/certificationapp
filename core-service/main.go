package main

import (
	"log"

	"github.com/atyagi9006/certificationapp/core-service/src/cache"

	"github.com/atyagi9006/certificationapp/core-service/src/service"
)

func main() {
	log.Printf("Starting %s \n", "core-service")
	service.InitCache()
	service.StartServer("8080")
}

func main_() {

	service.InitCache()
	client := cache.RedisClient()
	defer client.RClient.Close()
	//fmt.Println(db.DBConfig.DatabaseName)
	/* 	cqoe := models.CurrentQuestionOptionElement{
	   		QuestionID: "123",
	   		OptionIDs:  []string{"a", "b", "c", "d"},
	   	}
	   	marshal, _ := json.Marshal(cqoe.OptionIDs)
	   	client.Write(cqoe.QuestionID, marshal) */

	client.Read("134184")
	//client.Delete(cqoe.QuestionID)

	//	cache.Init()
}
