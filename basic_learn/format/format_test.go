package format

import (
	"testing"
	"time"
)

func TestFormatAny(t *testing.T) {
	var x int64 = 1
	var d time.Duration = 1 * time.Nanosecond
	c1 := make(chan int)
	var tests = []struct {
		input interface{}
		want  string
	}{
		{x, "1"},
		{d, "1"},
		{[]int64{x}, "???"},
		{[]time.Duration{d}, "???"},
		{c1, "???"},
	}

	for _, test := range tests {
		if got := Any(test.input); got != test.want {
			t.Errorf("TestformatAny(%q) = %v", test.input, got)
		}
	}
}
