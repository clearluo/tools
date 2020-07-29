package main

import (
	_ "finance/db"
	"finance/db/history"
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
	"time"
)

func main() {
	//Etf(0, 505, 0.00, 10, 0.33)
	//Snowball(2000000, 0.10, 28)
	//InstallmentCal(4500, 12, 385.35)
	//AnnualYield(5000000, 10000000000, 25)
	//YearRate()
	//Retire(1000000, 0.15, 0.04)
	History(1000000, 600036, 20020409, true)
}
func History(startMoney float64, code int, startTime int, isJoin bool) {
	// 600519 贵州茅台-20010827
	// 601318 中国平安-20070301
	// 000002 万科A-19910129
	// 600036 招商银行-20020409
	rows, err := history.GetHistoryByCodeAndTime(code, startTime)
	if err != nil || len(rows) < 1 {
		fmt.Println(err, " or rows is null")
		return
	}
	fmt.Printf("自%v开始买入%vW的%v股票，每年分红当天以收盘价分红投入，历史收益明细\n", startTime, startMoney/10000, rows[0].Name)

	var sumMoney float64 = startMoney // 总市值
	var sumShare float64              // 总股数
	var freeShare float64             // 转送股数
	var freeMoney float64             // 分红金额
	var buyShare float64              // 购买股数
	var blance float64 = startMoney   // 余额
	var sumAddMoney float64           // 参与配售投入金额
	var addCount int                  // 配售次数
	for _, row := range rows {
		if row.Typ == 1 {
			if !isJoin {
				fmt.Printf("%v不参与增发配售\n", row.DayTime)
				continue
			}
			tmpSumMoney := sumMoney
			addShare := sumShare * row.SharePer
			costMoney := addShare * row.BuyPrice
			sumAddMoney += costMoney
			sumShare += addShare
			sumMoney = sumShare*row.Price + blance
			fmt.Printf("%-9v 价格:%-8.2f 配股:%-12.2f 花费:%8.2fW 增发价格:%-8.2f 持股:%-10.0f 市值涨幅:%7.2f%%  总市值:%.0fW\n", fmt.Sprintf("%v*", row.DayTime), row.Price, addShare, 0-costMoney/10000, row.BuyPrice, sumShare, (sumMoney-tmpSumMoney)/tmpSumMoney*100, sumMoney/10000)
			addCount++
			continue
		}
		tmpSumMoney := sumMoney
		// 计算分红金额
		freeMoney = sumShare * row.MoneyPer
		// 计算赠送多少股
		freeShare = sumShare * row.SharePer
		// 计算可够买多少股
		tmpMoney := blance + freeMoney // 本次可购买的总金额
		buyShare = math.Floor(tmpMoney / (row.Price * 100))
		buyShare *= 100
		// 计算余额
		blance = tmpMoney - (buyShare * row.Price)
		// 计算总股数
		sumShare += buyShare + freeShare
		// 计算总市值
		sumMoney = sumShare*row.Price + blance
		fmt.Printf("%-9d 价格:%-8.2f 赠股:%-12.2f 分红:%8.2fW 分红买股:%-8.0f 持股:%-10.0f 市值涨幅:%7.2f%%  总市值:%.0fW\n", row.DayTime, row.Price, freeShare, freeMoney/10000, buyShare, sumShare, (sumMoney-tmpSumMoney)/tmpSumMoney*100, sumMoney/10000)
	}
	fmt.Printf("累计参与%d次配售，投入配售总额:%.2fW\n", addCount, sumAddMoney/10000)
	startYear := startTime / 10000
	endYear := time.Now().Year()
	AnnualYield(startMoney, sumMoney, endYear-startYear)
}

// Snowball 计算现在x元在股市未来n年后的价值
func Snowball(money float64, yearRate float64, yearCount int) float64 {
	sum := money
	for i := 0; i < yearCount; i++ {
		sum *= 1.0 + yearRate
	}
	fmt.Printf("%.0f元以年化%.1f%%增长在%d年后会变成%.2f万元\n", money, yearRate*100, yearCount, sum/10000)
	return sum
}

// Etf 每月定投股市收益计算
// startMoney: 初始投入多少
// monthMoney：初使每月定存多少
// yearInc：每月投入以多少比例增加
// (以年为单位递增，如：第一年每月投入1k，以0.05递增,那么第二年每月投入1.05k,第二年每月投入1.1025k)
// yearCount：定存年数30
// yearRate：期望平均年化复合收益率0.15
func Etf(startMoney float64, monthMoney float64, yearInc float64, yearCount int, yearRate float64) {
	var sumBase float64 = startMoney // 总本钱
	var sum float64 = sumBase        // 总收益
	var yearSum float64              // 一年中的利润
	for i := 1; i <= yearCount; i++ {
		yearSum = 0
		fmt.Printf("第%d年每月存入:%.2fW\n", i, monthMoney/10000)
		sumBase += monthMoney * 12
		for j := 1; j <= 12; j++ {
			sum += monthMoney
			yearSum += (sum * yearRate) / 12
			//fmt.Printf("第%d年,第%d个月:%.2fW\n", i, j, (sum+yearSum)/10000)
		}
		sum += yearSum
		monthMoney *= 1 + yearInc
		//fmt.Printf("第%d年后投入总金额:%.2fW\n", i, sumBase/10000)
		//fmt.Printf("第%d年后总金额:%.2fW\n", i, sum/10000)

	}
	fmt.Printf("总投入:%.2fW\n", sumBase/10000)
	fmt.Printf("总收益:%.2fW\n", sum/10000)
}

// InstallmentCal 根据贷款总额和每月还款金额计算实际贷款年化利率
// totalMoney:贷款总额(单位：元)
// monthCount:期数
// monthMoney:每期还款额(单位：元)
// 返回值：年化利率
// 每月还款额=[贷款本金×月利率×（1+月利率）^还款月数]÷[（1+月利率）^还款月数－1]
func InstallmentCal(totalMoney float64, monthCount int, monthMoney float64) float64 {
	// 采用最土的暴力破解法
	var rate = 0.0200 // 从2%开始算起，保留两位小数
	for ; rate < 1; rate += 0.0001 {
		monthRate := rate / 12
		monthX := (totalMoney * monthRate * math.Pow(1+monthRate, float64(monthCount))) /
			(math.Pow(1+monthRate, float64(monthCount)) - 1)
		if math.Abs(monthX-monthMoney) < 1 {
			tmp := rate * 100
			fmt.Printf("%.2f元,分%d期,每期还%.2f元,年化利率：%.2f%%\n", totalMoney, monthCount, monthMoney, tmp)
			return tmp
		}
	}
	fmt.Printf("年化利率：%.2f%%\n", 0)
	return 0
}

// AnnualYield 根据起始金额和结束金额和时间计算年化复合收益率
func AnnualYield(startMoney float64, endMoney float64, yearCount int) float64 {
	// 采用最土的暴力破解法
	var rate = 0.00200 // 从2%开始算起，保留两位小数
	for ; rate < 1; rate += 0.000001 {
		endMoneyTmp := startMoney
		for i := 0; i < yearCount; i++ {
			endMoneyTmp += endMoneyTmp * rate
		}
		if math.Abs(endMoneyTmp-endMoney)/endMoney < 0.0001 {
			tmp := rate * 100
			//fmt.Printf("endMoneyTmp:%.2f endMoney:%.2f\n", endMoneyTmp, endMoney)
			fmt.Printf("年化利率：%.2f%%\n", tmp)
			return tmp
		}
	}
	fmt.Printf("err")
	return 0
}

func YearRate() float64 {
	readContext, err := ioutil.ReadFile("data.txt")
	if err != nil {
		fmt.Println(err)
		return 0
	}
	lineSli := strings.Split(string(readContext), "\n")
	var principalTotal []float64
	var profitTotal float64
	for _, line := range lineSli {
		if len(line) < 3 || strings.HasPrefix(line, "#") {
			continue
		}
		line = strings.Trim(line, "\r")
		line = strings.Trim(line, " ")
		rowSli := strings.Split(line, ":")
		if len(rowSli) != 3 {
			fmt.Println("len(rowSli)!=3:", line)
			continue
		}
		rowSli[1] = strings.Trim(rowSli[1], " ")
		profitTmp, err := strconv.ParseFloat(rowSli[1], 64)
		if err != nil {
			fmt.Println("strconv.ParseFloat err:", rowSli[1])
			return 0
		}
		profitTotal += profitTmp
		principalSli := strings.Split(rowSli[2], ",")
		for _, v := range principalSli {
			v = strings.Trim(v, " ")
			principalTmp, err := strconv.ParseFloat(v, 64)
			if err != nil {
				fmt.Println("strconv.ParseFloat err:", v)
				return 0
			}
			principalTotal = append(principalTotal, principalTmp)
		}
	}
	var rate float64
	for rate = -0.9; rate < 1; rate += 0.00001 {
		monthRate := rate / 12
		sumProfit := 0.0
		for _, v := range principalTotal {
			sumProfit += v * monthRate
		}
		tmp := math.Abs(sumProfit - profitTotal)
		if tmp < 100 {
			fmt.Printf("total: %.2fW yearRate: %.2f%%\n", profitTotal/10000, rate*100)
			return rate
		}
	}
	fmt.Println("cal err")
	return 0
}

func Retire(base float64, rate float64, useRate float64) {
	fmt.Printf("初始本金: %.2fW, 投资年化收益率: %.2f%%, 每年腐败比例%.2f%%\n", base/10000, rate*100, useRate*100)
	for i := 0; i < 20; i++ {
		base *= 1 + rate
		use := base * useRate
		base -= use
		fmt.Printf("第%d年提取: %.2fW, 剩余: %.2fW\n", i+1, use/10000, base/10000)
	}
}
