// https://leetcode.cn/problems/max-points-on-a-line
// 直线上最多的点数
package main

type slope struct {
	x int // △x
	y int // △y
}

func maxPoints(points [][]int) (res int) {
	n := len(points)
	for i := 0; i < n; i++ {
		if res >= n-i || res > n/2 {
			break
		}
		kmap := map[slope]int{}
		for j := i + 1; j < n; j++ {
			x := points[j][0] - points[i][0] // △x
			y := points[j][1] - points[i][1] // △y
			g := gcd(x, y)
			kmap[slope{x / g, y / g}]++
		}
		maxn := 0
		for _, v := range kmap {
			maxn = max(maxn, v)
		}
		res = max(res, maxn+1)
	}
	return
}

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b, a = a%b, t
	}
	return a
}
