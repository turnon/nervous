package db

type DbHandler interface {
	LoadEvents(date string, months int) []*Event
	NewEvents(dates []string, tagName string, eventName string, days int) error
}

type Event struct {
	Title string `json:"title"`
	Start string `json:"start"`
	End   string `json:"end"`
}
