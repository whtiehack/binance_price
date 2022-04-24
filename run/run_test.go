package run

import "testing"

func TestRun(t *testing.T) {
	t.Log("Test Run")
	_, _ = Run(nil, DefineEvent{})
}
