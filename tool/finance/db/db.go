package db

import (
	"finance/common/basic"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

var (
	appDb *xorm.Engine
)

func InitMySql() {
	var err error
	dataSource := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&loc=Local",
		basic.MysqlApp.User, basic.MysqlApp.Password, basic.MysqlApp.Host, basic.MysqlApp.Port, basic.MysqlApp.Database)
	appDb, err = xorm.NewEngine("mysql", dataSource)
	if err != nil {
		err = fmt.Errorf("xorm.NewEngine err:%v", err)
		fmt.Println(err)
		panic(err)
	}

	appDb.DB().SetMaxOpenConns(100)
	appDb.DB().SetMaxIdleConns(30)
	appDb.DB().SetConnMaxLifetime(time.Second * 30)
	//appDb.ShowSQL(true) // debug 模式，打印执行的 sql
	//appDb.Logger().SetLevel(xormlog.LOG_DEBUG)
	if err := appDb.DB().Ping(); err != nil {
		err = fmt.Errorf("xorm ping err:%v", err)
		fmt.Println(err)
		fmt.Println(err)
		//panic(err)
	}

}
func GetAppDb() *xorm.Engine {
	return appDb
}
