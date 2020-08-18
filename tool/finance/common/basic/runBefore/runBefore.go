package runBefore

import (
	"finance/common/basic"
	"finance/db"
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
