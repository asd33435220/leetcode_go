package main

func main() {
	//["DetectSquares", "add", "add", "add", "count", "count", "add", "count"]
	//	[[], [[3, 10]], [[11, 2]], [[3, 2]], [[11, 10]], [[14, 8]], [[11, 2]], [[11, 10]]]
	//	输出：
	//	[null, null, null, null, 1, 0, null, 2]
	detect := Constructor()
	detect.Add([]int{3, 10})
	detect.Add([]int{11, 2})
	detect.Add([]int{3, 2})
	detect.Count([]int{11, 10})
	detect.Count([]int{14, 8})
	detect.Add([]int{11, 2})
	detect.Count([]int{11, 10})
}

type DetectSquares map[int]map[int]int

func Constructor() DetectSquares {
	return DetectSquares{}
}

func (d DetectSquares) Add(point []int) {
	x, y := point[0], point[1]
	if d[x] == nil {
		d[x] = make(map[int]int)
	}
	d[x][y]++
}

func (d DetectSquares) Count(point []int) (ans int) {
	x, y := point[0], point[1]
	if d[x] == nil {
		return
	}
	xCnt := d[x]
	for col, colCnt := range d {
		if col != x {
			// 根据对称性，这里可以不用取绝对值
			dis := col - x
			ans += colCnt[y] * xCnt[y+dis] * colCnt[y+dis]
			ans += colCnt[y] * xCnt[y-dis] * colCnt[y-dis]
		}
	}
	return
}
