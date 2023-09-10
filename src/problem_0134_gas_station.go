// https://leetcode.cn/problems/gas-station
// 加油站
package main

func canCompleteCircuit(gas, cost []int) (start int) {
	n := len(gas)
	for start < n {
		sumOfGas := 0
		cnt := 0
		for ; cnt < n; cnt++ {
			cur := (start + cnt) % n
			sumOfGas += gas[cur] - cost[cur]
			if sumOfGas < 0 {
				break
			}
		}
		if cnt == n {
			return
		}
		start += cnt + 1
	}
	return -1
}
