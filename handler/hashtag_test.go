package handler_test

import (
	"testing"

	"github.com/cage1016/alfred-change-case/handler"
)

func Test_HashTag(t *testing.T) {

	cases := []struct {
		in, want string
	}{
		{"1 2 3", "#1 #2 #3"},
		{"#1 #2 #3", "#1 #2 #3"},
		{"1 2    3", "#1 #2 #3"},
		{"1 2    #3", "#1 #2 #3"},
	}

	for _, c := range cases {
		got := handler.HashTag(c.in)
		if got != c.want {
			t.Errorf("HashTag(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
