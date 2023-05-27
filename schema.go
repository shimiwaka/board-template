package schema

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type Board struct {
	gorm.Model `json:"-"`
	Owner      string         `json:"owner"`
}

func test() {
	fmt.Println("test")
}