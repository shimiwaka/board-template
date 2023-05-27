package main

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
	"github.com/jinzhu/gorm"
)

type Settings struct {
	Username string `yaml:"db_username"`
	Pass string `yaml:"db_pass"`
	Host string `yaml:"db_host"`
	Port string `yaml:"db_port"`
	Name string `yaml:"db_name"`
}

func main() {
	settings := Settings{}
	b, _ := os.ReadFile("../../config/db.yaml")
	yaml.Unmarshal(b, &settings)

	connectQuery := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		settings.Username, settings.Pass, settings.Host, settings.Port, settings.Name)

	db, err := gorm.Open("mysql", connectQuery)

	if err != nil {
		fmt.Printf("error: failed to connect DB")
		return
	}
	db.Exec("DELETE FROM boards")
	db.AutoMigrate(&Board{})
}