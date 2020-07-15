package models

type History struct {
	Id       int     `xorm:"not null pk autoincr INT(11)"`
	Code     int     `xorm:"default 600519 INT(11)"`
	Name     int     `xorm:"default 0 INT(11)"`
	DayTime  int     `xorm:"default 0 comment('20200715') index INT(11)"`
	Price    float64 `xorm:"default 0 comment('当天收盘价') DOUBLE"`
	Dividend string  `xorm:"default '' comment('分红说明') VARCHAR(50)"`
	MoneyPer float64 `xorm:"default 0 comment('每股分红金额') DOUBLE"`
	SharePer float64 `xorm:"default 0 comment('每股赠送股数') DOUBLE"`
}
