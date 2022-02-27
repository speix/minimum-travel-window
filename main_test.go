package main

import (
	"reflect"
	"testing"
)

func TestNormalSeries(t *testing.T) {
	var test = []string{"New York", "Paris", "New York", "Paris", "Tokyo", "Paris", "Athens", "Tokyo"}
	expectedResult := &Result{
		days:  5,
		path:  []string{"New York", "Paris", "Tokyo", "Paris", "Athens"},
		start: 2,
		stop:  6,
	}

	result := solution(test)
	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("Fail: wanted %v got %v", expectedResult, result)
	}
}

func TestEdgeCase(t *testing.T) {
	var test = []string{"New York", "Paris", "New York", "Paris", "New York", "Paris"}
	expectedResult := &Result{
		days:  2,
		path:  []string{"New York", "Paris"},
		start: 0,
		stop:  1,
	}

	result := solution(test)
	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("Fail: wanted %v got %v", expectedResult, result)
	}
}
