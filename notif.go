package models

type Notification struct {
	Title    string      `json:"title"`
	Body     string      `json:"body"`
	Topic    string      `json:"topic"`
	Priority int         `json:"priority"`
	Data     interface{} `json:"data"`
}
