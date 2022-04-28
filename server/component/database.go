package component

import (
	"database/sql"

	_ "github.com/gogf/gf/contrib/drivers/sqlite/v2"
	_ "github.com/mattn/go-sqlite3"

	"context"
	"os"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/gfile"
)

var (
	databaseInstance gdb.DB
)

func InitDatabase() {
	var err error
	createSQLITEIfNotExist()
	setDatabaseConfig()
	databaseInstance, err = gdb.NewByGroup("default")
	ctx := context.Background()
	if err != nil {
		Logger().Error(ctx, err)
		panic(err)
	}
}

func setDatabaseConfig() {
	var (
		config *gcfg.Config
	)
	config = gcfg.Instance()
	pwd, _ := os.Getwd()
	config.GetAdapter().(*gcfg.AdapterFile).SetFileName(pwd + "/server/config/config.json")
	ctx := context.Background()
	debug, _ := config.Get(ctx, "database.debug")
	dbType, _ := config.Get(ctx, "database.type")
	gdb.SetConfig(gdb.Config{
		"default": gdb.ConfigGroup{
			gdb.ConfigNode{
				Debug:  debug.Bool(),
				Type:   dbType.String(),
				Link:   "sqlite:" + pwd + "/db.sqlite3",
				Weight: 100,
			},
		},
	})
}

func createSQLITEIfNotExist() {
	if gfile.Exists("./db.sqlite3") {
		return
	}
	db, err := sql.Open("sqlite3", "./db.sqlite3")
	if err != nil {
		panic(err)
	}
    sql_table := gfile.GetContents("./server/config/model.sql")

    _, err = db.Exec(sql_table)
		if err != nil {
		panic(err)
	}
}

func GetDatabase() gdb.DB {
	return databaseInstance
}
