package http

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/Batyachelly/goBoard/internal/database/models"
	"github.com/Batyachelly/goBoard/internal/transport/http/data"

	"github.com/gorilla/mux"
)

// Get boards
// @Summary      Get boards list
// @Description  Get boards list
// @Tags         main
// @Produce      json
// @Success      200  {object}  data.GetBoardsResponse
// @Router       /board [get]
func (s *Server) GetBoards(w http.ResponseWriter, r *http.Request) {
	modelBoards, err := s.usecase.GetBoardList(r.Context())
	if err != nil {
		s.log.Error("GET boards: %w", err)

		s.responseJSON(w, http.StatusInternalServerError, nil)
	}

	dataBoards := make(data.GetBoardsResponse, 0, len(modelBoards))

	for _, modelBoard := range modelBoards {
		dataBoards = append(dataBoards, data.GetBoardResponse{
			ID:    modelBoard.ID,
			Title: modelBoard.Title,
		})
	}

	s.responseJSON(w, http.StatusOK, dataBoards)
}

// Get board
// @Summary      Get board
// @Description  Get board by id
// @Tags         main
// @Produce      json
// @Param        board_id   path int  true  "board ID"
// @Success      200  {object}  data.GetBoardResponse
// @Router       /board/{board_id} [get]
func (s *Server) GetBoard(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	boardID, err := strconv.ParseUint(vars["board_id"], 10, 64)
	if err != nil {
		s.responseJSON(w, http.StatusBadRequest, nil)
	}

	modelBoard, err := s.usecase.GetBoard(r.Context(), boardID)
	if err != nil {
		s.log.Error("GET board: %w", err)

		s.responseJSON(w, http.StatusInternalServerError, nil)
	}

	dataBoard := new(data.GetBoardResponse)

	dataBoard.ID = modelBoard.ID
	dataBoard.Title = modelBoard.Title
	dataBoard.Threads = make([]data.Message, 0, len(modelBoard.Threads))

	for _, thread := range modelBoard.Threads {
		dataBoard.Threads = append(dataBoard.Threads, data.Message{
			ID:      thread.ID,
			Title:   thread.Title,
			Text:    thread.Text,
			Content: thread.Content,
			Created: thread.Created,
		})
	}

	s.responseJSON(w, http.StatusOK, dataBoard)
}

// Get thread
// @Summary      Get thread
// @Description  Get thread by board ID and thread ID
// @Tags         main
// @Produce      json
// @Param        board_id   path int  true  "board ID"
// @Param        thread_id  path int  true  "thread ID"
// @Success      200  {object}  data.GetThread
// @Router       /board/{board_id}/thread/{thread_id} [get]
func (s *Server) GetThread(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	boardID, err := strconv.ParseUint(vars["board_id"], 10, 64)
	if err != nil {
		s.responseJSON(w, http.StatusBadRequest, nil)
	}

	threadID, err := strconv.ParseUint(vars["thread_id"], 10, 64)
	if err != nil {
		s.responseJSON(w, http.StatusBadRequest, nil)
	}

	modelMessages, err := s.usecase.GetThread(r.Context(), boardID, threadID)
	if err != nil {
		s.log.Error("GET thread: %w", err)

		s.responseJSON(w, http.StatusInternalServerError, nil)
	}

	respMessages := make(data.GetThread, 0, len(modelMessages))

	for _, modelMessage := range modelMessages {
		respMessages = append(respMessages, data.Message{
			ID:      modelMessage.ID,
			Title:   modelMessage.Title,
			Text:    modelMessage.Text,
			Content: modelMessage.Content,
			Created: modelMessage.Created,
		})
	}

	s.responseJSON(w, http.StatusOK, respMessages)
}

// Post thread
// @Summary      Post thread
// @Description  Create new thread by board ID
// @Tags         main
// @Accept       json
// @Produce      json
// @Param        board_id   path int  true  "board ID"
// @Param        thread body data.PostThreadRequest true "Thread create request"
// @Success      200  {object}  data.PostThreadResponse
// @Router       /board/{board_id}/thread [post]
func (s *Server) PostThread(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		s.responseJSON(w, http.StatusBadRequest, nil)
	}

	boardID, err := strconv.ParseUint(vars["board_id"], 10, 64)
	if err != nil {
		s.responseJSON(w, http.StatusBadRequest, nil)
	}

	thread := new(data.PostThreadRequest)

	if err := json.Unmarshal(body, thread); err != nil {
		s.responseJSON(w, http.StatusBadRequest, nil)
	}

	threadID, err := s.usecase.PostThread(r.Context(), &models.Message{
		BoardID: boardID,
		Title:   thread.Title,
		Text:    thread.Text,
		Content: thread.Content,
	})
	if err != nil {
		s.log.Error("POST thread: %w", err)

		s.responseJSON(w, http.StatusInternalServerError, nil)
	}

	s.responseJSON(w, http.StatusOK, &data.PostThreadResponse{ThreadID: threadID})
}

// Post message
// @Summary      Post message
// @Description  Create new message by board ID and thread ID
// @Tags         main
// @Accept       json
// @Produce      json
// @Param        board_id   path int  true  "board ID"
// @Param        thread_id  path int  true  "thread ID"
// @Param        thread body data.PostMessageRequest true "Message create request"
// @Success      200  {object}  data.PostMessageResponse
// @Router       /board/{board_id}/thread/{thread_id} [post]
func (s *Server) PostMessage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		s.responseJSON(w, http.StatusBadRequest, nil)
	}

	comment := new(data.PostMessageRequest)

	if err := json.Unmarshal(body, comment); err != nil {
		s.responseJSON(w, http.StatusBadRequest, nil)
	}

	boardID, err := strconv.ParseUint(vars["board_id"], 10, 64)
	if err != nil {
		s.responseJSON(w, http.StatusBadRequest, nil)
	}

	threadID, err := strconv.ParseUint(vars["thread_id"], 10, 64)
	if err != nil {
		s.responseJSON(w, http.StatusBadRequest, nil)
	}

	messageID, err := s.usecase.PostMessage(r.Context(), &models.Message{
		BoardID:  boardID,
		ThreadID: threadID,
		Title:    comment.Title,
		Text:     comment.Text,
		Content:  comment.Content,
	})
	if err != nil {
		s.log.Error("POST message: %w", err)

		s.responseJSON(w, http.StatusInternalServerError, nil)
	}

	s.responseJSON(w, http.StatusOK, &data.PostMessageResponse{MessageID: messageID})
}

func (s Server) responseJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")

	if status != http.StatusOK {
		w.WriteHeader(status)
	}

	if err := json.NewEncoder(w).Encode(data); err != nil {
		s.log.Error("%w", err.Error)
	}
}
