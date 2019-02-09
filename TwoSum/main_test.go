package main

import (
	"reflect"
	"testing"
)

type TestCase struct {
	data  string
	count int
}

func TestAllAvailableFunctions(t *testing.T) {
	expected := []int{0, 1}
	res := twoSum([]int{2, 7, 11, 15}, 9)

	if !reflect.DeepEqual(res, expected) {
		t.Errorf("Expected %v got %v", expected, res)
	}
}
