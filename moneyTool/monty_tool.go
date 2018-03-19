package main

import (
	"fmt"
)

func main() {
	etf300()
}

/*
 每月定存etf300指数基金计划
*/
func etf300() {
	var monthMondy float64 = 0.1 // 每月定存多少W
	var yearCount int = 30       // 定存年数
	var yearRate float64 = 0.1   // 平均年化复合收益率
	var sum float64 = 0          // 总收益
	for i := 1; i <= yearCount; i++ {
		//fmt.Printf("第%d年每月存入:%f\n", i, monthMondy)
		for j := 0; j < 12; j++ {
			sum += monthMondy
			sum = (sum*yearRate)/12 + sum
		}
		fmt.Printf("第%d年后总金额:%.2fW\n", i, sum)
	}
}
