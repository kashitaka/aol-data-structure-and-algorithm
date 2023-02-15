package main

import "testing"

func Test(t *testing.T) {
	solve([]*card{
		{"D", 3},
		{"H", 2},
		{"D", 1},
		{"S", 3},
		{"D", 2},
		{"C", 1},
	}, []*card{
		{"D", 3},
		{"H", 2},
		{"D", 1},
		{"S", 3},
		{"D", 2},
		{"C", 1},
	})
}
