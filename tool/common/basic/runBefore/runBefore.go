package runBefore

import (
	"tool/common/basic"
	"tool/common/db"
)

func InitRun() {
	// 初始化Msyql
	db.InitMySql()
	// 初始化定时任务
	if basic.App.RunCron {
		go func() {
			//crontab.CronMain()
		}()
	}
}
