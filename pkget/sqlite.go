package pkget

import (
	"log"

	"github.com/jinzhu/gorm"
	// sqlite3 added
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type sqlite struct {
	path string
	db   *gorm.DB
}

// Package : the main sqlite table
type Package struct {
	gorm.Model
	Name        string
	Version     uint
	Arch        string
	OS          string
	Repo        string
	HomePage    string
	Size        string
	Publisher   string
	Section     string
	Description string
	PkVersion   string
	Depends     []Package `gorm:"many2many:package_package;"`
}

func (s sqlite) openConnect() {
	db, err := gorm.Open("sqlite3", "./pkget.db")
	if err != nil {
		log.Fatal(err)
	}
	s.db = db

	db.AutoMigrate(&Package{})
	db.Model(&Package{}).Related(&Package{}, "Package")
}

func (s sqlite) closeConnect() {
	s.db.Close()
}

func (s sqlite) install(pk Package, deps []Package) {
	s.db.Create(&pk)
	if deps != nil {
		for dep := range deps {
			s.db.Create(dep)
		}
	}
}

func (s sqlite) InstallPackage(name string) {

}
