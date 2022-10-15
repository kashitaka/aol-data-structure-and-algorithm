package main

import "testing"

func Test(t *testing.T) {
	in := []string{"5 100", "p1 150", "p2 80", "p3 200", "p4 350", "p5 20"}
	solve(in)
}
