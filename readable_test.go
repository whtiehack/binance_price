package main

import "testing"

func TestHumanReadable(t *testing.T) {
	t.Log(parseHumanReadableQuality(2132))
	t.Log(parseHumanReadableQuality(21340))
	t.Log(parseHumanReadableQuality(213401))
	t.Log(parseHumanReadableQuality(2134010))
	t.Log(parseHumanReadableQuality(21340104))
	t.Log(parseHumanReadableQuality(2134010411))
	t.Log(parseHumanReadableQuality(42134010411))
}
