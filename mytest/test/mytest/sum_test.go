package mytest

import "testing"
import "mytest"

func TestSum(t *testing.T) {
	actual := mytest.MySum(10, 11, 12)
	expected := 33

	if actual != expected {
		t.Errorf("got %v\nant: %v", actual, expected)
	}
}

func TestMyFoldSum(t *testing.T) {
	actual := mytest.MyFoldSum([]int{10, 11, 12})
	expected := 33

	if actual != expected {
		t.Errorf("got %v\nant: %v", actual, expected)
	}
}
