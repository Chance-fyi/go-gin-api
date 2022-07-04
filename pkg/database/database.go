package database

import (
	"fmt"
	"go-gin-api/pkg/console"
	"gorm.io/gorm"
)

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

func (db *db) GetDefaultConnect() string {
	return db.defaultConnect
}

func (db *db) Connect(name ...string) *gorm.DB {
	if len(name) > 0 {
		connection, ok := db.connections[name[0]]
		if !ok {
			console.Errorp(fmt.Sprintf("%s:connection does not exist", name[0]))
		}
		return connection
	}
	return db.connections[db.defaultConnect]
}

func (db *db) CreateConnection(name string, dialector gorm.Dialector) {
	open, err := gorm.Open(dialector)
	console.ExitIf(err)
	db.connections[name] = open
}
