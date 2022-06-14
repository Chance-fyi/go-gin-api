package database

import "gorm.io/gorm"

type db struct {
	defaultConnect string
	connections    map[string]*gorm.DB
}

var DB = db{
	defaultConnect: "",
	connections:    map[string]*gorm.DB{},
}

func (db *db) SetDefaultConnect(name string) {
	db.defaultConnect = name
}

func (db *db) Connect(name string) *gorm.DB {
	return db.connections[name]
}

func (db *db) CreateConnection(name string, dialector gorm.Dialector) {
	open, err := gorm.Open(dialector)
	if err != nil {
		panic(err)
	}
	db.connections[name] = open
}
