package dao

import (
	"context"
	"fmt"
	"log"
	"time"

	cfg "github.com/atyagi9006/certificationapp/data-service/src/config"
)

var Config *cfg.Database
var Reader ReadDataInterface
var Writer WriteDataInterface

func init() {
	fmt.Println("Hello World")
}

func Init(conf *cfg.Database) {

	Config = conf
	Reader = NewDBReader()
	Writer = NewDBWriter()

	Reader.GetAdmin(context.Background(), "test-project-db")
	//Reader.GetAdmin(context.Background(), "admin")

	//go Process(Config.DBConfig.Type)
}

func GetDBReader() *ReadDataInterface {
	return &Reader
}
func GetDBWriter() *WriteDataInterface {
	return &Writer
}

func Process(dbType string) {
	time.Sleep(1 * time.Second)
	log.Println("Hello World", dbType)
}
