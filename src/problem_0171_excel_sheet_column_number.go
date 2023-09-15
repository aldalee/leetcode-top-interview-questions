// https://leetcode.cn/problems/excel-sheet-column-number
// Excel 表列序号
package main

func titleToNumber(columnTitle string) (num int) {
	for i := 0; i < len(columnTitle); i++ {
		k := columnTitle[i] - 'A' + 1
		num = num*26 + int(k)
	}
	return
}
