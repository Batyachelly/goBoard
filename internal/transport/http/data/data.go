package data

import "time"

type GetBoardResponse struct {
	ID    uint64 `json:"id"`
	Title string `json:"title"`

	Threads GetThread `json:"threads,omitempty"`
}

type Message struct {
	ID      uint64    `json:"id"`
	Title   string    `json:"title"`
	Text    string    `json:"text"`
	Content string    `json:"content"`
	Created time.Time `json:"created"`
}

type GetBoardsResponse []GetBoardResponse

type GetThread []Message

type PostThreadRequest struct {
	Title   string `json:"title"`
	Text    string `json:"text"`
	Content string `json:"content"`
}

type PostThreadResponse struct {
	ThreadID uint64 `json:"threadId"`
}

type PostMessageRequest struct {
	Title   string `json:"title"`
	Text    string `json:"text"`
	Content string `json:"content"`
}

type PostMessageResponse struct {
	MessageID uint64 `json:"messageId"`
}
