package main

import (
	"testing"
)

func TestParse(t *testing.T) {
	res := ParseFile(".vne-config")

	if (len(res) != 2) {
		t.Errorf("size of parse result is %d and not 2", len(res))
	}

	if res[0] != "[envs]" || res[1] != "test" {
		t.Errorf("incorrect parsing: %s!=[envs] or %s!=test", res[0], res[1])
	}

}