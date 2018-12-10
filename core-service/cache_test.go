package main

import (
	"encoding/json"
	"testing"

	"github.com/atyagi9006/certificationapp/core-service/src/cache"
	"github.com/atyagi9006/certificationapp/core-service/src/config"

	"github.com/atyagi9006/certificationapp/core-service/src/models"
)

func TestWrite(t *testing.T) {
	cqoe := models.CurrentQuestionOptionElement{
		QuestionID: "123",
		OptionIDs:  []string{"a", "b", "c", "d"},
	}
	db := config.Init()
	client := cache.RedisClient(db)
	marshal, _ := json.Marshal(cqoe.OptionIDs)
	client.Write(cqoe.QuestionID, marshal)
}

func TestRead(t *testing.T) {
	db := config.Init()
	client := cache.RedisClient(db)
	client.Read("123")
}
