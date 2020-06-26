package goroutines_test

import (
	"testing"
	"time"

	"wesionary.team/dipeshdulal/golang-test/goroutines"
)

func TestAdd2ValuesAfter2Seconds(t *testing.T) {
	before := time.Now()
	val := goroutines.Add2ValuesAfter2Seconds(2, 4)
	after := time.Now()
	t.Logf("1. Value after adding 2 values is %v took %v", val, after.Sub(before))

	before = time.Now()
	go goroutines.Add2ValuesAfter2Seconds(2, 4)
	after = time.Now()
	t.Logf("1. Value after adding 2 values is %v took %v", val, after.Sub(before))

}
