package memoization_test

import (
	"testing"

	"wesionary.team/dipeshdulal/golang-test/reflection/memoization"
)

func TestCacherInput(t *testing.T) {
	mFnc := func(d int) int {
		return d * 2
	}
	memoization.Cacher(mFnc, 200)
	t.Logf("Success: takes input functions")

	d := 20
	if _, err := memoization.Cacher(d, 300); err != nil {
		t.Logf("Success: takes only input functions")
	}
}
