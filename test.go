package util

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/tealeg/xlsx"
)

func do_check_table(item map[string]string) {
	filename := item["_filename"]
	mtable := item["_table"]

	if mtable == "tpl_plot" || mtable == "" {
		return
	}
	fmt.Printf("=========================================\n")
	fmt.Printf("开始解析数据【%s】 =>【%s】\n", filename, mtable)
	xlFile, err := xlsx.OpenFile(filename)
	if err != nil {
		fmt.Println(err)
		panic("error excel")
	}

	sheet := xlFile.Sheets[0]
	// 数据检查
	for idx, row := range sheet.Rows {
		cell := row.Cells
		// 首行过滤
		if idx < 1 {
			continue
		}
		// 空行结束
		if len(cell) < 1 {
			break
		}
		// 过滤 非有效行
		if cell[0].String() != "*" {
			continue
		}

		mapObj := make(map[string]interface{}, 0)

		for attr, rule := range item {
			if attr[0] == 95 { // _忽略
				continue
			}
			values := strings.Split(rule, "_")

			if len(values) < 2 {
				fmt.Printf("【%s】属性列【%s】不符合规范\n", mtable, attr)
				panic("配置异常")
			}
			if values[0] == "~" {
			} else {
				column, errCol := strconv.Atoi(values[1])
				if errCol != nil {
					column = col2int(values[1])
				}
				if column < 1 {
					fmt.Printf("【%s】属性列【%s】=【%s】不符合规范\n", mtable, attr, rule)
					myPanic("配置异常")
				}
				check_cell(cell, idx, column-1, values[0])
				if len(cell) >= column {
					mapObj[attr] = cell[column-1]
				}
			}
		}
		checkTabValue(mtable, mapObj)
	}
}
