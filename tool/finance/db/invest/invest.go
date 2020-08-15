package invest

import "finance/db"

type Invest struct {
	Id        int     `xorm:"not null pk autoincr INT(11)"`
	MonthTime int     `xorm:"not null default 0 comment('年月') index INT(11)"`
	Capital   float64 `xorm:"not null default 0 comment('当月本金') DOUBLE"`
	Profit    float64 `xorm:"not null default 0 DOUBLE"`
	Uid       int     `xorm:"not null default 0 index INT(11)"`
}

func GetInvestByUid(uid int) ([]*Invest, error) {
	rows := []*Invest{}
	sql := `SELECT * FROM invest WHERE uid=? ORDER BY month_time`
	err := db.GetAppDb().SQL(sql, uid).Find(&rows)
	return rows, err
}
