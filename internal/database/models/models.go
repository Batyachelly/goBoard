package models

import "time"

const (
	Deleted = iota
	Active
)

type Board struct {
	ID    uint64 `json:"id"`
	Title string `json:"title"`

	Threads MessageList `json:"threads,omitempty"`
}

type Message struct {
	ID       uint64    `json:"id"`
	BoardID  uint64    `json:"-"`
	ThreadID uint64    `json:"-"`
	Title    string    `json:"title"`
	Text     string    `json:"text"`
	Content  string    `json:"content"`
	Created  time.Time `json:"created"`
}

type BoardList []Board

type MessageList []Message
