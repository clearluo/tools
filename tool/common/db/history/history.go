package history

import "tool/common/db"

type History struct {
	Id       int     `xorm:"not null pk autoincr INT(11)"`
	Code     int     `xorm:"default 600519 INT(11)"`
	Name     string  `xorm:"default '' comment('名字') VARCHAR(50)"`
	DayTime  int     `xorm:"default 0 comment('20200715') index INT(11)"`
	Typ      int     `xorm:"default 0 INT(11)"`
	Price    float64 `xorm:"default 0 comment('当天收盘价') DOUBLE"`
	Dividend string  `xorm:"default '' comment('分红说明') VARCHAR(50)"`
	MoneyPer float64 `xorm:"default 0 comment('每股分红金额') DOUBLE"`
	SharePer float64 `xorm:"default 0 comment('每股赠送股数') DOUBLE"`
	BuyPrice float64 `xorm:"default 0 comment('增发配售价格') DOUBLE"`
}

func GetHistoryByCodeAndTime(code int, startTime int) ([]*History, error) {
	rows := []*History{}
	sql := `SELECT * FROM history WHERE code=? AND day_time>=? ORDER BY day_time,typ`
	err := db.GetAppDb().SQL(sql, code, startTime).Find(&rows)
	return rows, err
}
