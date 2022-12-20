package array

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

func TestReverseArray(t *testing.T) {
	var testCases = []struct {
		in, out []int32
	}{
		{in: []int32{1, 2, 3}, out: []int32{3, 2, 1}},
		{in: []int32{1, 4, 3, 2}, out: []int32{2, 3, 4, 1}},
	}

	for i, test := range testCases {
		name := fmt.Sprintf("test [%d]", i+1)
		t.Run(name, func(t *testing.T) {
			res := reverseArray(test.in)
			if !reflect.DeepEqual(res, test.out) {
				t.Errorf("%s ex: %+v got: %+v\n", name, test.out, res)
			}
		})
	}
}

func TestHourglassSum(t *testing.T) {
	var testCases = []struct {
		in  [][]int32
		out int32
	}{
		{in: [][]int32{
			{-9, -9, -9, 1, 1, 1},
			{0, -9, 0, 4, 3, 2},
			{-9, -9, -9, 1, 2, 3},
			{0, 0, 8, 6, 6, 0},
			{0, 0, 0, -2, 0, 0},
			{0, 0, 1, 2, 4, 0},
		},
			out: 28,
		},
		{in: [][]int32{
			{1, 1, 1, 0, 0, 0},
			{0, 1, 0, 0, 0, 0},
			{1, 1, 1, 0, 0, 0},
			{0, 0, 2, 4, 4, 0},
			{0, 0, 0, 2, 0, 0},
			{0, 0, 1, 2, 4, 0},
		},
			out: 19,
		},
		{in: [][]int32{
			{0, -4, -6, 0, -7, -6},
			{-1, -2, -6, -8, -3, -1},
			{-8, -4, -2, -8, -8, -6},
			{-3, -1, -2, -5, -7, -4},
			{-3, -5, -3, -6, -6, -6},
			{-3, -6, 0, -8, -6, -7},
		},
			out: -19,
		},
	}

	for i, test := range testCases {
		name := fmt.Sprintf("test [%d]", i+1)
		t.Run(name, func(t *testing.T) {
			res := hourglassSum(test.in)
			if !reflect.DeepEqual(res, test.out) {
				t.Errorf("%s ex: %+v got: %+v\n", name, test.out, res)
			}
		})
	}
}

func TestDynamicArray(t *testing.T) {
	var testCases = []struct {
		n       int32
		queries [][]int32
		out     []int32
	}{
		{
			n: 2,
			queries: [][]int32{
				{1, 0, 5},
				{1, 1, 7},
				{1, 0, 3},
				{2, 1, 0},
				{2, 1, 1},
			},
			out: []int32{7, 3},
		},
	}

	for i, test := range testCases {
		name := fmt.Sprintf("test [%d]", i+1)
		t.Run(name, func(t *testing.T) {
			res := dynamicArray(test.n, test.queries)
			if !reflect.DeepEqual(res, test.out) {
				t.Errorf("%s ex: %+v got: %+v\n", name, test.out, res)
			}
		})
	}
}

func TestMatchingStrings(t *testing.T) {
	var testCases = []struct {
		stringList, queries []string
		out                 []int32
	}{
		{
			stringList: []string{"aba", "baba", "aba", "xzxb"},
			queries:    []string{"aba", "xzxb", "ab"},
			out:        []int32{2, 1, 0},
		},
		{
			stringList: []string{"def", "de", "fgh"},
			queries:    []string{"de", "lmn", "fgh"},
			out:        []int32{1, 0, 1},
		},
		{
			stringList: []string{"abcde", "sdaklfj", "asdjf",
				"na", "basdn", "sdaklfj", "asdjf",
				"na", "asdjf", "na", "basdn", "sdaklfj", "asdjf"},
			queries: []string{"abcde", "sdaklfj", "asdjf", "na", "basdn"},
			out:     []int32{1, 3, 4, 3, 2},
		},
	}

	for i, test := range testCases {
		name := fmt.Sprintf("test [%d]", i+1)
		t.Run(name, func(t *testing.T) {
			res := matchingStrings(test.stringList, test.queries)
			if !reflect.DeepEqual(res, test.out) {
				t.Errorf("%s ex: %+v got: %+v\n", name, test.out, res)
			}
		})
	}
}

func TestArrayManipulation(t *testing.T) {
	var testCases = []struct {
		n       int32
		out     int64
		queries [][]int32
	}{
		{n: 10, out: 10, queries: [][]int32{{1, 5, 3}, {4, 8, 7}, {6, 9, 1}}},
		{n: 5, out: 200, queries: [][]int32{{1, 2, 100}, {2, 5, 100}, {3, 4, 100}}},
	}

	for i, test := range testCases[:1] {
		name := fmt.Sprintf("test [%d]", i+1)
		t.Run(name, func(t *testing.T) {
			res := arrayManipulation(test.n, test.queries)
			if !reflect.DeepEqual(res, test.out) {
				t.Errorf("%s ex: %+v got: %+v\n", name, test.out, res)
			}
		})
	}

	t.Run("time limit exceeded test", func(t *testing.T) {
		ex := int64(2497169732) // test 7
		res := arrayManipulation(mustBuildQueriesFromFile(t))
		if !reflect.DeepEqual(res, ex) {
			t.Errorf("time limit exceeded test failed ex: %+v got: %+v\n", ex, res)
		}
	})

}

func mustBuildQueriesFromFile(t *testing.T) (int32, [][]int32) {
	t.Helper()
	data, err := ioutil.ReadFile("testdata/arrayManipulation_test")
	if err != nil {
		panic(err)
	}

	readSeeker := bytes.NewReader(data)
	firstRow, err := bufio.NewReader(readSeeker).ReadString('\n')
	if err != nil {
		panic(err)
	}
	sizes := strings.Split(strings.TrimRight(firstRow, "\r\n"), " ")
	if len(sizes) != 2 || len(sizes) == 0 {
		panic(err)
	}

	n, err := strconv.Atoi(sizes[0])
	if err != nil {
		panic(err)
	}
	m, err := strconv.Atoi(sizes[1])
	if err != nil {
		panic(err)
	}
	_, err = readSeeker.Seek(int64(len(firstRow)), io.SeekStart)
	if err != nil {
		panic(err)
	}

	var queries = make([][]int32, m)
	iter := 0
	scanner := bufio.NewScanner(readSeeker)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		txt := strings.TrimRight(scanner.Text(), "\r\n")
		strNumbers := strings.Split(txt, " ")
		for _, strNumber := range strNumbers {
			v, err := strconv.ParseInt(strNumber, 10, 64)
			if err != nil {
				panic(err)
			}
			queries[iter] = append(queries[iter], int32(v))
		}
		iter++
	}

	if err = scanner.Err(); err != nil {
		panic(err)
	}
	return int32(n), queries
}
