package main

import "testing"

func TestRun(t *testing.T) {
	t.Log("Test Run")
	_, _ = run(nil, DefineEvent{})
}
