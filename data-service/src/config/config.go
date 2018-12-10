package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

func init() {
	log.Println("Config Loaded")
}

type Config struct {
	Type         string `yaml:"type"`
	URL          string `yaml:"url"`
	Port         int16  `yaml:"port"`
	DatabaseName string `yaml:"database-name"`
	UserType     string `yaml:"user-type"`
	UserName     string `yaml:"user-name"`
	Password     string `yaml:"password"`
	ArangoPort   int16  `yaml:"arangoPort"`
}

type Database struct {
	DBConfig Config `yaml:"database"`
}

func Init() *Database {
	log.Println("config.Init")
	d := &Database{}

	return d.GetConf(ConfigLocation)

}

func (d *Database) GetConf(confLoaction string) *Database {
	yamlFile, err := ioutil.ReadFile(confLoaction)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, &d)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return d
}
