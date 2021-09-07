package ent

import (
	"context"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/turnon/nervous/db"
	"github.com/turnon/nervous/ent/event"
	"github.com/turnon/nervous/ent/tag"
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

type DbHandler struct {
}

func (*DbHandler) LoadEvents(date string, months int) []*db.Event {
	t, _ := time.Parse("2006-01-02", date)
	start_at := t.AddDate(0, -1, 0)
	end_at := t.AddDate(0, months, 0)

	events := xclient.Debug().Event.Query().
		Where(event.StartAtLTE(end_at), event.EndAtGTE(start_at)).
		Order(Asc(event.FieldStartAt), Asc(event.FieldEndAt)).
		AllX(context.Background())

	var outEvent []*db.Event
	for _, e := range events {
		outEvent = append(outEvent, &db.Event{
			Title: e.Name,
			Start: e.StartAt.Format("2006-01-02"),
			End:   e.EndAt.Format("2006-01-02"),
		})
	}

	return outEvent
}

func (*DbHandler) NewEvents(dates []string, tagName string, eventName string, days int) error {
	tx, err := xclient.Debug().Tx(context.Background())
	if err != nil {
		return err
	}
	defer tx.Commit()

	if days == 0 {
		days = 1
	}

	tagObj, err := tx.Tag.Query().Where(tag.Name(tagName)).Only(context.Background())
	if _, ok := err.(*NotFoundError); ok {
		tagObj = tx.Tag.Create().SetName(tagName).SaveX(context.Background())
	}

	for _, date := range dates {
		startAt := parseTime(date)
		endAt := parseTime(date).AddDate(0, 0, days)
		tx.Event.Create().
			SetStartAt(startAt).
			SetEndAt(endAt).
			SetName(eventName).
			SetTag(tagObj).
			Save(context.Background())
	}

	return nil
}

func parseTime(str string) time.Time {
	rt, err := time.Parse("2006-1-2", str)
	if err != nil {
		panic(err)
	}
	return rt
}
