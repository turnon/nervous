package ent

import "testing"

func TestLoadEvents(t *testing.T) {
	dbHandler := DbHandler{}
	t.Log(dbHandler.LoadEvents("2021-09-03", 2))
}
