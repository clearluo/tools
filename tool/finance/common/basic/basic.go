package basic

import (
	"finance/common/log"
	"os"
	"path/filepath"

	ini "gopkg.in/ini.v1"
)

// MysqlType 存储数据库相关配置
type MysqlType struct {
	User     string
	Password string
	Host     string
	Port     string
	Database string
}

type RedisType struct {
	Host     string
	Password string
	Port     string
}

type AppType struct {
	Port    string
	RpcPort string
	WsPort  string
	Secret  string
	RunCron bool
	BinName string
}
type PathType struct {
	LogDir string
}

var (
	MysqlApp MysqlType
	App      AppType
	Path     PathType
	Redis    RedisType
)

func init() {
	initIni()
	initDir()
	initLog()
}

func initDir() {
	os.MkdirAll(Path.LogDir, os.ModePerm)
}
func initLog() {
	fileName := filepath.Join(Path.LogDir, "console.log")
	logConfig := log.LogConfig{
		Filename:        fileName,
		RetainFileCount: 2048,
	}
	logConn := log.SetLogger(logConfig)
	_ = logConn
	log.SetLevel(log.DebugLevel)
}

func initIni() {
	cfg, err := ini.Load("config/config.ini")
	if err != nil {
		panic(err)
	}
	MysqlApp.User = cfg.Section("mysqlApp").Key("user").String()
	MysqlApp.Password = cfg.Section("mysqlApp").Key("password").String()
	MysqlApp.Host = cfg.Section("mysqlApp").Key("host").String()
	MysqlApp.Port = cfg.Section("mysqlApp").Key("port").String()
	MysqlApp.Database = cfg.Section("mysqlApp").Key("database").String()
	checkIni()
}

func checkIni() {
	//if len(Path.LogDir) < 3 {
	//	err := fmt.Errorf("logsdir err in conf.ini")
	//	fmt.Println(err)
	//	panic(err)
	//}
}
