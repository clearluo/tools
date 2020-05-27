package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func main() {
	//Etf(0, 0.2, 0.05, 30, 0.20)
	//InstallmentCal(15000, 12, 1329.2)
	//Snowball(8000, 0.98, 31)
	//AnnualYield(14200, 16200, 5)
	YearRate()
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
// startMoney: 初始投入多少W
// monthMoney：初使每月定存多少W
// yearInc：每月投入以多少比例增加
// (以年为单位递增，如：第一年每月投入1k，以0.05递增,那么第二年每月投入1.05k,第二年每月投入1.1025k)
// yearCount：定存年数30
// yearRate：期望平均年化复合收益率0.15
// 返回值：总收入W
func Etf(startMoney float64, monthMoney float64, yearInc float64, yearCount int, yearRate float64) float64 {
	var sumBase float64 = startMoney // 总本钱
	var sum float64 = sumBase        // 总收益
	for i := 1; i <= yearCount; i++ {
		fmt.Printf("第%d年每月存入:%fW\n", i, monthMoney)
		sumBase += monthMoney * 12
		for j := 1; j <= 12; j++ {
			sum += monthMoney
			sum += (sum * yearRate) / 12
			fmt.Printf("第%d年,第%d个月:%.2fW\n", i, j, sum)
		}
		monthMoney *= 1 + yearInc
		//fmt.Printf("第%d年后投入总金额:%.2fW\n", i, sumBase)
		//fmt.Printf("第%d年后总金额:%.2fW\n", i, sum)

	}
	fmt.Printf("总投入:%.2fW\n", sumBase)
	fmt.Printf("总收益:%.2fW\n", sum)
	return sum
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
			fmt.Println("len(rowSli)!=3")
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
			fmt.Printf("total: %v yearRate: %.2f%%\n", profitTotal, rate*100)
			return rate
		}
	}
	fmt.Println("cal err")
	return 0
}
