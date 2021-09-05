package ent

import (
	"context"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/turnon/nervous/ent/event"
)

var xclient *Client

func init() {
	dbFile := os.Getenv("NERVOUS_DB")
	if dbFile == "" {
		panic("NERVOUS_DB undefined")
	}

	c, err := Open("sqlite3", "file:"+dbFile+"?cache=shared&_fk=1")
	if err != nil {
		panic(err)
	}
	if err := c.Schema.Create(context.Background()); err != nil {
		c.Close()
		panic(err)
	}
	xclient = c
}

func LoadEvents(date string, months int) []*Event {
	t, _ := time.Parse("2006-01-02", date)
	start_at := t.AddDate(0, -1, 0)
	end_at := t.AddDate(0, months, 0)

	return xclient.Debug().Event.Query().
		Where(event.StartAtLTE(end_at), event.EndAtGTE(start_at)).
		Order(Asc(event.FieldStartAt), Asc(event.FieldEndAt)).
		AllX(context.Background())
}
