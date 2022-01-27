package http_test

import (
	"encoding/json"
	"io"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/Batyachelly/goBoard/generated/mocks"
	"github.com/Batyachelly/goBoard/internal/database"
	"github.com/Batyachelly/goBoard/internal/database/models"
	"github.com/Batyachelly/goBoard/internal/logger"
	"github.com/Batyachelly/goBoard/internal/transport/http"
	"github.com/Batyachelly/goBoard/internal/usecase"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestServer_GetBoards(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		ds   database.Databaser
		want models.BoardList
	}{
		{
			name: "1",
			ds: func() database.Databaser {
				ds := &mocks.Databaser{}
				ds.On("GetBoardList", mock.Anything).Once().Return(models.BoardList{
					{
						ID:    1,
						Title: "Title1",
					},
					{
						ID:    2,
						Title: "Title2",
					},
				}, nil)

				return ds
			}(),
			want: models.BoardList{
				{
					ID:    1,
					Title: "Title1",
				},
				{
					ID:    2,
					Title: "Title2",
				},
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			req := httptest.NewRequest("GET", "/board", nil)
			w := httptest.NewRecorder()

			s := http.NewServer(http.Config{
				Log: logger.TestLogger{},
			}, usecase.NewUsecase(tt.ds))

			s.GetBoards(w, req)

			body, _ := io.ReadAll(w.Result().Body)

			wantJSON, _ := json.Marshal(tt.want)

			require.JSONEq(t, string(wantJSON), string(body))
		})
	}
}

func TestServer_GetBoard(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		boardID int64
		ds      database.Databaser
		want    *models.Board
	}{
		{
			name:    "1",
			boardID: 2,
			ds: func() database.Databaser {
				ds := &mocks.Databaser{}
				ds.On("GetBoard", mock.Anything, uint64(2)).Once().Return(&models.Board{
					ID:    2,
					Title: "Title",
					Threads: models.MessageList{
						{
							ID:       1,
							BoardID:  2,
							ThreadID: 1,
							Title:    "Title1",
							Text:     "Text1",
							Content:  "Context1",
							Created:  time.Time{}.Add(time.Hour),
						},
						{
							ID:       2,
							BoardID:  2,
							ThreadID: 1,
							Title:    "Title2",
							Text:     "Text2",
							Content:  "Context2",
							Created:  time.Time{}.Add(2 * time.Hour),
						},
					},
				}, nil)

				return ds
			}(),
			want: &models.Board{
				ID:    2,
				Title: "Title",
				Threads: models.MessageList{
					{
						ID:       1,
						BoardID:  2,
						ThreadID: 1,
						Title:    "Title1",
						Text:     "Text1",
						Content:  "Context1",
						Created:  time.Time{}.Add(time.Hour),
					},
					{
						ID:       2,
						BoardID:  2,
						ThreadID: 1,
						Title:    "Title2",
						Text:     "Text2",
						Content:  "Context2",
						Created:  time.Time{}.Add(2 * time.Hour),
					},
				},
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			req := httptest.NewRequest("GET", "/board/"+strconv.FormatInt(tt.boardID, 10), nil)
			req = mux.SetURLVars(req, map[string]string{"board_id": strconv.FormatInt(tt.boardID, 10)})

			w := httptest.NewRecorder()

			s := http.NewServer(http.Config{
				Log: logger.TestLogger{},
			}, usecase.NewUsecase(tt.ds))

			s.GetBoard(w, req)

			body, _ := io.ReadAll(w.Result().Body)

			wantJSON, _ := json.Marshal(tt.want)

			require.JSONEq(t, string(wantJSON), string(body))
		})
	}
}
