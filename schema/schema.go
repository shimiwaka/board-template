package schema

import (
	"github.com/jinzhu/gorm"
)

type Board struct {
	gorm.Model `json:"-"`
	Owner      string         `json:"owner"`
}