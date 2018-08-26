package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("总收入：%.2fW\n",Etf300(0.1,30,0.15))
	fmt.Printf("年化利率：%.2f%%\n",InstallmentCal(14744, 12, 1271.67))
}

/*
 * 功能：每月定存etf300指数基金计划
 * yearRate：每月定存多少W
 * yearCount：定存年数30
 * yearRate：平均年化复合收益率0.15
 * 返回值：总收入W
 */
func Etf300(monthMondy float64,yearCount int,yearRate float64)float64 {
	var sum float64 = 0          // 总收益
	for i := 1; i <= yearCount; i++ {
		//fmt.Printf("第%d年每月存入:%f\n", i, monthMondy)
		for j := 0; j < 12; j++ {
			sum += monthMondy
			sum = (sum*yearRate)/12 + sum
		}
		fmt.Printf("第%d年后总金额:%.2fW\n", i, sum)
	}
	return sum
}

/*
 * 功能：根据贷款总额和每月还款金额计算实际贷款年化利率
 * totalMoney:贷款总额(单位：元)
 * monthCount:期数
 * monthMoney:每期还款额(单位：元)
 * 返回值：年化利率
 * 每月还款额=[贷款本金×月利率×（1+月利率）^还款月数]÷[（1+月利率）^还款月数－1]
 */
func InstallmentCal(totalMoney float64, monthCount int, monthMoney float64) float64 {
	// 采用最土的暴力破解法
	var rate float64 = 0.0200 // 从2%开始算起，保留两位小数
	for ; rate < 1; rate += 0.0001 {
		monthRate := rate / 12
		monthX := (totalMoney * monthRate * math.Pow(1+monthRate, float64(monthCount))) /
			(math.Pow(1+monthRate, float64(monthCount)) - 1)
		if math.Abs(monthX-monthMoney) < 1 {
			return rate * 100
		}
	}
	return 0
}