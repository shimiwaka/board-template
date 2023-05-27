package main

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
	// "github.com/jinzhu/gorm"
	// "github.com/shimiwaka/board-template/schema"
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
	fmt.Printf("%s\n", settings.Username)
}