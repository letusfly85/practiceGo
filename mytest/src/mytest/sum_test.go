package mytest

import "testing"
import "mytest"

func TestSum(t *testing.T) {
	actual := mytest.MySum(10)
	expected := 20

	if actual != expected {
		t.Errorf("got %v\nant: %v", actual, expected)
	}
}
