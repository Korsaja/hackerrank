package math

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFindPoint(t *testing.T) {
	var testCases = []struct {
		px, py, qx, qy int32
		out            []int32
	}{
		{px: 0, py: 0, qx: 1, qy: 1, out: []int32{2, 2}},
		{px: 1, py: 1, qx: 2, qy: 2, out: []int32{3, 3}},
	}

	for i, test := range testCases {
		name := fmt.Sprintf("test [%d]", i+1)
		t.Run(name, func(t *testing.T) {
			points := findPoint(test.px, test.py, test.qx, test.qy)
			if !reflect.DeepEqual(points, test.out) {
				t.Errorf("%s fail ex: %+v got: %+v", name, test.out, points)
			}
		})
	}

}

func TestMaximumDraws(t *testing.T) {
	var testCases = []struct {
		n, out int32
	}{
		{n: 1, out: 2},
		{n: 2, out: 3},
	}

	for i, test := range testCases {
		name := fmt.Sprintf("test [%d]", i+1)
		t.Run(name, func(t *testing.T) {
			res := maximumDraws(test.n)
			if !reflect.DeepEqual(res, test.out) {
				t.Errorf("%s fail ex: %+v got: %+v", name, test.out, res)
			}
		})
	}
}

func TestHandshake(t *testing.T) {
	var testCases = []struct {
		n, out int32
	}{
		{n: 1, out: 0},
		{n: 2, out: 1},
		{n: 6, out: 15},
	}

	for i, test := range testCases {
		name := fmt.Sprintf("test [%d]", i+1)
		t.Run(name, func(t *testing.T) {
			res := handshake(test.n)
			if !reflect.DeepEqual(res, test.out) {
				t.Errorf("%s fail ex: %+v got: %+v", name, test.out, res)
			}
		})
	}
}

func TestLowestTriangle(t *testing.T) {
	var testCases = []struct {
		base, area, out int32
	}{
		{base: 2, area: 2, out: 2},
		{base: 17, area: 100, out: 12},
	}

	for i, test := range testCases {
		name := fmt.Sprintf("test [%d]", i+1)
		t.Run(name, func(t *testing.T) {
			res := lowestTriangle(test.base, test.area)
			if !reflect.DeepEqual(res, test.out) {
				t.Errorf("%s fail ex: %+v got: %+v", name, test.out, res)
			}
		})
	}
}
