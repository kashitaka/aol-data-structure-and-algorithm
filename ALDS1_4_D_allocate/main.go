package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var sc = bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	sc.Scan()
	numPackage, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	numTrack, _ := strconv.Atoi(sc.Text())
	weights := make([]int, numPackage)
	for i := 0; i < numPackage; i++ {
		sc.Scan()
		v, _ := strconv.Atoi(sc.Text())
		weights[i] = v
	}
	solve(numPackage, numTrack, weights)
}

func solve(packCnt, trackCnt int, weights []int) {
	left := 1
	right := 100000 * 10000 // 100000個おもさ10000を1台で運ぶ
	res := binarySearch(left, right, packCnt, trackCnt, weights)
	fmt.Println(res)
}

func binarySearch(left, right, target, trucks int, weights []int) int {
	if left >= right {
		return right
	}
	mid := (left + right) / 2
	availableBoxCnt := BoxCnt(mid, trucks, weights)
	if target <= availableBoxCnt {
		right = mid
		return binarySearch(left, right, target, trucks, weights)
	} else {
		left = mid + 1
		return binarySearch(left, right, target, trucks, weights)
	}
}

func BoxCnt(p, trucks int, weights []int) int {
	res := 0
	trackWeight := 0
	i := 0
	for trucks > 0 && i < len(weights) {
		v := weights[i]
		if trackWeight+v <= p {
			// この箱は乗せる
			res++
			i++
			trackWeight += v
		} else {
			// truckは出荷する
			trucks--
			trackWeight = 0
		}
	}
	return res
}
