package main_test

import "testing"

func Test(t *testing.T) {
	if 1 == 2 {
		t.FailNow()
	}
}
