package ent

import (
	"context"
	"os"
	"strconv"
	"testing"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/turnon/nervous/ent/event"
	"github.com/turnon/nervous/ent/tag"
)

func TestConnection(t *testing.T) {
	pid := strconv.Itoa(os.Getpid())

	// test connection
	client, err := Open("sqlite3", "file:ent.test.db?cache=shared&_fk=1", Log(t.Log))
	if err != nil {
		t.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	if err := client.Schema.Create(context.Background()); err != nil {
		t.Fatalf("failed creating schema resources: %v", err)
	}

	tx, _ := client.Debug().Tx(context.Background())
	defer tx.Commit()

	// clear
	tx.Event.Delete().Exec(context.Background())
	tx.Tag.Delete().Exec(context.Background())

	// test create
	tagLiving, err := tx.Tag.Create().SetName("living:" + pid).Save(context.Background())
	if err != nil {
		t.Errorf("failed creating tag: %w", err)
	}
	_, err = tx.Tag.Create().SetName("working:" + pid).Save(context.Background())
	if err != nil {
		t.Errorf("failed creating tag: %w", err)
	}

	happenings := [][]time.Time{
		{parseTime("2021-08-10"), parseTime("2021-08-16")},
		{parseTime("2021-09-12"), parseTime("2021-09-13")},
		{parseTime("2021-09-23"), parseTime("2021-09-25")},
		{parseTime("2021-10-02"), parseTime("2021-10-04")},
	}

	for _, happening := range happenings {
		ev, err := tx.Event.
			Create().
			SetName("go shopping:" + pid).
			SetStartAt(happening[0]).
			SetEndAt(happening[1]).
			AddTags(tagLiving).
			Save(context.Background())
		if err != nil {
			t.Errorf("failed creating event: %w", err)
		}
		t.Log("event was created: ", ev)
	}

	// test query
	sep_start_at := parseTime("2021-09-01")
	sep_end_at := parseTime("2021-09-30")

	evs, err := tx.Event.
		Query().
		Where(event.EndAtGTE(sep_start_at), event.StartAtLTE(sep_end_at)).
		Order(Asc(event.FieldStartAt), Asc(event.FieldEndAt)).
		All(context.Background())
	if err != nil {
		t.Errorf("failed querying event: %w", err)
	}
	if len(evs) != 2 {
		t.Error("event returned: ", evs)
	}
	t.Log("event returned: ", evs)

	// test update
	firstEvent := evs[0]
	firstEvent, err = firstEvent.
		Update().
		SetName("go to sleep:" + pid).
		Save(context.Background())
	if err != nil {
		t.Errorf("failed update event: %w", err)
	}
	if firstEvent.Name != "go to sleep:"+pid {
		t.Error("failed update event: ", firstEvent)
	}
	t.Log("event returned: ", firstEvent)

	// test query by edge
	oneTag, err := tx.Tag.Query().Where(tag.HasEventsWith(event.Name("go to sleep:" + pid))).Only(context.Background())
	if err != nil {
		t.Errorf("failed find tag: %w", err)
	}
	if oneTag.Name != "living:"+pid {
		t.Error("failed find tag: ", oneTag)
	}
	t.Log("tag returned: ", oneTag)
}

func parseTime(str string) time.Time {
	rt, err := time.Parse("2006-1-2", str)
	if err != nil {
		panic(err)
	}
	return rt
}
