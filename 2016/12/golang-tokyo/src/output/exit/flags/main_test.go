package main

import (
	"strings"
	"testing"
)

func TestRun(t *testing.T) {
	cases := []struct {
		Command string
		Status  int
	}{
		{"exec testdata/test.json arg1 arg2", 0},
		{"exec testdata/invalid.json arg1 arg2", 1},
	}

	for _, tc := range cases {
		args := strings.Split(tc.Command, " ")
		if got := Run(args); got != tc.Status {
			t.Fatalf("Run exit %d, want = %d", got, tc.Status)
		}
	}
}
