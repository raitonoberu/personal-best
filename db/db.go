package db

import (
	"fmt"

	"github.com/glebarez/sqlite"
	"github.com/raitonoberu/personal-best/app/model"
	"gorm.io/gorm"
)

var db *gorm.DB

func Get() *gorm.DB {
	return db
}

func Init(dsn string) error {
	d, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("couldn't open db: %x", err)
	}

	for _, m := range []any{
		(*model.User)(nil),
		(*model.Competition)(nil),
		(*model.Game)(nil),
		(*model.Player)(nil),
	} {
		if err := d.AutoMigrate(m); err != nil {
			return fmt.Errorf("couldn't migrate %#v: %x", m, err)
		}
	}
	db = d
	return nil
}
