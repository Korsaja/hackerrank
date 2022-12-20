package array

func reverseArray(a []int32) []int32 {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	return a
}

func hourglassSum(arr [][]int32) int32 {
	var size = len(arr) >> 1
	var res = int32(-1 << 31)
	for i := 0; i <= size; i++ {
		for j := 0; j <= size; j++ {
			top := arr[i][j] + arr[i][j+1] + arr[i][j+2]
			mid := arr[i+1][j+1]
			bot := arr[i+2][j] + arr[i+2][j+1] + arr[i+2][j+2]
			res = max(res, top+mid+bot)
		}
	}

	return res
}

func dynamicArray(n int32, queries [][]int32) []int32 {
	lastAnswer := int32(0)
	arrays := make([][]int32, n)
	answer := make([]int32, 0)

	getIndex := func(x int32) int32 { return (x ^ lastAnswer) % n }

	for _, query := range queries {
		typ, x, y := query[0], query[1], query[2]
		i := getIndex(x)
		if typ == 1 {
			arrays[i] = append(arrays[i], y)
		} else if typ == 2 {
			lastAnswer = arrays[i][y%int32(len(arrays[i]))]
			answer = append(answer, lastAnswer)
		}
	}
	return answer
}

func max(i int32, j int32) int32 {
	if i > j {
		return i
	}
	return j
}

// medium
func matchingStrings(stringList []string, queries []string) []int32 {
	var counter = make(map[string]int32)
	for _, str := range stringList {
		counter[str]++
	}

	out := make([]int32, len(queries))
	for i, query := range queries {
		c, ok := counter[query]
		if ok {
			out[i] = c
		}

	}

	return out
}

// hard
func arrayManipulation(n int32, queries [][]int32) int64 {
	var arr = make([]int64, n+1)
	for _, q := range queries {
		a, b, k := q[0]-1, q[1], q[2]
		arr[a] += int64(k)
		if b < int32(len(arr)) {
			arr[b] -= int64(k)
		}
	}

	var sum, maxSum int64
	for i := range arr {
		sum += arr[i]
		if sum > maxSum {
			maxSum = sum
		}
	}

	return maxSum
}
