package boot

import (
	"errors"
	"fmt"
	"go-gin-api/pkg/config"
	Db "go-gin-api/pkg/database"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type database struct {
	Config dbConfig
}

type dbConfig struct {
	Default     string
	Connections map[string]struct {
		Type     string
		Hostname string
		Port     int
		Username string
		Password string
		Database string
		Charset  string
	}
}

var Database = database{}

func (db *database) Init() {
	config.UnmarshalKey("database", &Database.Config)
	Db.DB.SetDefaultConnect(Database.Config.Default)
	db.setupDb()
}

func (db *database) setupDb() {
	var dialector gorm.Dialector
	for name, conn := range Database.Config.Connections {
		switch conn.Type {
		case "mysql":
			dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=%v&parseTime=True&loc=Local",
				conn.Username,
				conn.Password,
				conn.Hostname,
				conn.Port,
				conn.Database,
				conn.Charset,
			)
			dialector = mysql.New(mysql.Config{
				DSN: dsn,
			})
		default:
			panic(errors.New("database connection not supported"))
		}
		Db.DB.CreateConnection(name, dialector)
	}
}
