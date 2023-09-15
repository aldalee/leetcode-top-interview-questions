// https://leetcode.cn/problems/excel-sheet-column-number
// Excel 表列序号
package main

func titleToNumber(columnTitle string) (num int) {
	n, pow := len(columnTitle), 1
	for i := n - 1; i >= 0; i-- {
		k := columnTitle[i] - 'A' + 1
		num += int(k) * pow
		pow *= 26
	}
	return
}
