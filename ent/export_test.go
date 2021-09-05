package ent

import "testing"

func TestLoadEvents(t *testing.T) {
	t.Log(LoadEvents("2021-09-03", 2))
}
