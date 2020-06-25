package calc_test

import (
	"testing"

	"wesionary.team/dipeshdulal/golang-test/calc"
)

func TestMathAdd(t *testing.T) {
	res, _ := calc.Add(1, 2, 3)

	if res != 6 {
		t.Errorf(`Add(1,2,3) FAILED, expected %v but got %v`, 6, res)
	} else {
		t.Logf(`Add(1,2,3) PASSED, expected %v and got value %v`, 6, res)
	}

}

// TestDataItem struct
type TestDataItem struct {
	Inputs   []int
	Result   int
	HasError bool
}

func TestMultipleMathAdd(t *testing.T) {
	dataItems := []TestDataItem{
		{[]int{1, 2, 3}, 6, false},
		{[]int{99, 99}, 198, false},
		{[]int{1, 1, 5, 5, 6}, 18, false},
		{[]int{1}, 0, true},
	}

	for _, item := range dataItems {
		res, err := calc.Add(item.Inputs...)
		if item.HasError {
			if err == nil {
				t.Errorf(`Add() with args %v : FAILED, expected an error but got value %v`, item.Inputs, res)
			} else {
				t.Logf(`Add() with args %v : PASSED, expected an error and got an error %v`, item.Inputs, err)
			}
		} else {
			if res != item.Result {
				t.Errorf("Add() with args %v : FAILED, expected %v but got value %v", item.Inputs, item.Result, res)
			} else {
				t.Logf("Add() with args %v : PASSED, expected %v and got value %v", item.Inputs, item.Result, res)
			}
		}
	}
}
