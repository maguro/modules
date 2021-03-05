package gorm

import (
	"gorm.io/gorm"

	"github.com/maguro/modules/platform/pg"
)

type Connection struct {
	db *gorm.DB
}

func (c Connection) Load(query string) interface{} {
	panic("implement me")
}

var _ pg.Connection = &Connection{}
