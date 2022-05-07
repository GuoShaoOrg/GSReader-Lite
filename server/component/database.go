package component

import (
	"context"
	"database/sql"

	"github.com/gogf/gf/v2/os/gfile"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	databaseInstance *gorm.DB
)

func InitDatabase() {
	var err error
	ctx := context.Background()
	createSQLITEIfNotExist(ctx)
	setDatabaseConfig(ctx)
	if err != nil {
		Logger().Error(ctx, err)
		panic(err)
	}
}

func setDatabaseConfig(ctx context.Context) {
	var err error

	dbConfig := gorm.Config{}
	databaseInstance, err = gorm.Open(sqlite.Open("./db.sqlite3"), &dbConfig)
	if err != nil {
		Logger().Error(ctx, err)
	}
}

func createSQLITEIfNotExist(ctx context.Context) {
	if gfile.Exists("./db.sqlite3") {
		return
	}
	db, err := sql.Open("sqlite3", "./db.sqlite3")
	if err != nil {
		Logger().Error(ctx, err)
		panic(err)
	}
    sql_table := gfile.GetContents("./server/config/schema.sql")

    _, err = db.Exec(sql_table)
		if err != nil {
		Logger().Error(ctx, err)
		panic(err)
	}
}

func GetDatabase() *gorm.DB {
	return databaseInstance
}
