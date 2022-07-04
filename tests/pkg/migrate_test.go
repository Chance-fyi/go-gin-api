package pkg

import (
	. "github.com/smartystreets/goconvey/convey"
	"go-gin-api/boot"
	"go-gin-api/pkg/g"
	"go-gin-api/pkg/migrate"
	"gorm.io/gorm"
	"testing"
)

func TestMigrate(t *testing.T) {
	boot.Boot.Init()
	type TestTable1 struct {
		Id   int
		Name string
	}
	type TestTable2 struct {
		Id   int
		Name string
	}
	Convey("database migration", t, func() {
		migrate.Add(migrate.MigrationFile{
			FileName: "test_create_test_table_1",
			Up: func(db *gorm.DB) error {
				return db.Migrator().AutoMigrate(&TestTable1{})
			},
			Down: func(db *gorm.DB) error {
				return db.Migrator().DropTable(&TestTable1{})
			},
		}, "default")
		migrate.Add(migrate.MigrationFile{
			FileName: "test_create_test_table_1",
			Up: func(db *gorm.DB) error {
				return db.Migrator().AutoMigrate(&TestTable1{})
			},
			Down: func(db *gorm.DB) error {
				return db.Migrator().DropTable(&TestTable1{})
			},
		}, "test")

		migrate.Run(migrate.Up)
		So(g.DB().Migrator().HasTable(&migrate.Migration{}), ShouldEqual, true)
		So(g.DB().Migrator().HasTable(&TestTable1{}), ShouldEqual, true)
		So(g.DB("test").Migrator().HasTable(&migrate.Migration{}), ShouldEqual, true)
		So(g.DB("test").Migrator().HasTable(&TestTable1{}), ShouldEqual, true)

		migrate.Add(migrate.MigrationFile{
			FileName: "test_create_test_table_2",
			Up: func(db *gorm.DB) error {
				return db.Migrator().AutoMigrate(&TestTable2{})
			},
			Down: func(db *gorm.DB) error {
				return db.Migrator().DropTable(&TestTable2{})
			},
		}, "default")

		migrate.Run(migrate.Up)
		So(g.DB().Migrator().HasTable(&TestTable1{}), ShouldEqual, true)
		So(g.DB().Migrator().HasTable(&TestTable2{}), ShouldEqual, true)
		So(g.DB("test").Migrator().HasTable(&TestTable1{}), ShouldEqual, true)

		migrate.Run(migrate.Reset)
		So(g.DB().Migrator().HasTable(&TestTable1{}), ShouldEqual, false)
		So(g.DB().Migrator().HasTable(&TestTable2{}), ShouldEqual, false)
		So(g.DB("test").Migrator().HasTable(&TestTable1{}), ShouldEqual, false)

		migrate.Run(migrate.Up)
		So(g.DB().Migrator().HasTable(&TestTable1{}), ShouldEqual, true)
		So(g.DB().Migrator().HasTable(&TestTable2{}), ShouldEqual, true)
		So(g.DB("test").Migrator().HasTable(&TestTable1{}), ShouldEqual, true)

		migrate.Add(migrate.MigrationFile{
			FileName: "test_create_test_table_2",
			Up: func(db *gorm.DB) error {
				return db.Migrator().AutoMigrate(&TestTable2{})
			},
			Down: func(db *gorm.DB) error {
				return db.Migrator().DropTable(&TestTable2{})
			},
		}, "test")

		migrate.Run(migrate.Up)
		So(g.DB().Migrator().HasTable(&TestTable1{}), ShouldEqual, true)
		So(g.DB().Migrator().HasTable(&TestTable2{}), ShouldEqual, true)
		So(g.DB("test").Migrator().HasTable(&TestTable1{}), ShouldEqual, true)
		So(g.DB("test").Migrator().HasTable(&TestTable2{}), ShouldEqual, true)

		migrate.Run(migrate.Rollback)
		So(g.DB().Migrator().HasTable(&TestTable1{}), ShouldEqual, false)
		So(g.DB().Migrator().HasTable(&TestTable2{}), ShouldEqual, false)
		So(g.DB("test").Migrator().HasTable(&TestTable1{}), ShouldEqual, true)
		So(g.DB("test").Migrator().HasTable(&TestTable2{}), ShouldEqual, false)

		migrate.Run(migrate.Reset)
	})
}
