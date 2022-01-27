package usecase_test

import (
	"context"
	"database/sql"
	"errors"
	"testing"
	"time"

	"github.com/Batyachelly/goBoard/generated/mocks"
	"github.com/Batyachelly/goBoard/internal/database"
	"github.com/Batyachelly/goBoard/internal/database/models"
	"github.com/Batyachelly/goBoard/internal/usecase"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUsecase_GetBoardList(t *testing.T) {
	t.Parallel()

	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		ds      database.Databaser
		want    models.BoardList
		wantErr error
	}{
		{
			name: "1",
			args: args{
				ctx: context.Background(),
			},
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
		{
			name: "2 error, thread not found",
			args: args{
				ctx: context.Background(),
			},
			ds: func() database.Databaser {
				ds := &mocks.Databaser{}
				ds.On("GetBoardList", mock.Anything).Once().Return(nil, sql.ErrNoRows)

				return ds
			}(),
			wantErr: sql.ErrNoRows,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			s := usecase.NewUsecase(tt.ds)
			got, err := s.GetBoardList(tt.args.ctx)
			if (err != nil || tt.wantErr != nil) && !errors.Is(err, tt.wantErr) {
				t.Errorf("Usecase.GetBoardList() error = %v, wantErr %v", err, tt.wantErr)

				return
			}

			if !assert.Equal(t, got, tt.want) {
				t.Errorf("Usecase.GetBoardList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUsecase_GetBoard(t *testing.T) {
	t.Parallel()

	type args struct {
		ctx     context.Context
		boardID uint64
	}
	tests := []struct {
		name    string
		args    args
		ds      database.Databaser
		want    *models.Board
		wantErr error
	}{
		{
			name: "1",
			args: args{
				ctx:     context.Background(),
				boardID: 101,
			},
			ds: func() database.Databaser {
				ds := &mocks.Databaser{}
				ds.On("GetBoard", mock.Anything, uint64(101)).Once().Return(&models.Board{
					ID:    101,
					Title: "TestBoard",
					Threads: models.MessageList{
						{
							ID:      1,
							Title:   "Title1",
							Text:    "Text1",
							Content: "Content1",
						},
						{
							ID:      2,
							Title:   "Title2",
							Text:    "Text2",
							Content: "Content2",
						},
					},
				}, nil)

				return ds
			}(),
			want: &models.Board{
				ID:    101,
				Title: "TestBoard",
				Threads: models.MessageList{
					{
						ID:      1,
						Title:   "Title1",
						Text:    "Text1",
						Content: "Content1",
					},
					{
						ID:      2,
						Title:   "Title2",
						Text:    "Text2",
						Content: "Content2",
					},
				},
			},
		},
		{
			name: "2 error, thread not found",
			args: args{
				ctx:     context.Background(),
				boardID: 101,
			},
			ds: func() database.Databaser {
				ds := &mocks.Databaser{}
				ds.On("GetBoard", mock.Anything, uint64(101)).Once().Return(nil, sql.ErrNoRows)

				return ds
			}(),
			wantErr: sql.ErrNoRows,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			s := usecase.NewUsecase(tt.ds)
			got, err := s.GetBoard(tt.args.ctx, tt.args.boardID)
			if (err != nil || tt.wantErr != nil) && !errors.Is(err, tt.wantErr) {
				t.Errorf("Usecase.GetBoard() error = %v, wantErr %v", err, tt.wantErr)

				return
			}

			if !assert.Equal(t, got, tt.want) {
				t.Errorf("Usecase.GetBoard() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUsecase_GetThread(t *testing.T) {
	t.Parallel()

	type args struct {
		ctx      context.Context
		boardID  uint64
		threadID uint64
	}
	tests := []struct {
		name    string
		args    args
		ds      database.Databaser
		want    models.MessageList
		wantErr error
	}{
		{
			name: "1",
			args: args{
				ctx:      context.Background(),
				boardID:  101,
				threadID: 202,
			},
			ds: func() database.Databaser {
				ds := &mocks.Databaser{}
				ds.On("GetThread", mock.Anything, uint64(101), uint64(202)).Once().Return(models.MessageList{
					{
						ID:       1,
						BoardID:  101,
						ThreadID: 202,
						Title:    "Title1",
						Text:     "Text1",
						Content:  "Content1",
						Created:  time.Time{},
					},
					{
						ID:       2,
						BoardID:  101,
						ThreadID: 202,
						Title:    "Title2",
						Text:     "Text2",
						Content:  "Content2",
					},
				}, nil)

				return ds
			}(),
			want: models.MessageList{
				{
					ID:       1,
					BoardID:  101,
					ThreadID: 202,
					Title:    "Title1",
					Text:     "Text1",
					Content:  "Content1",
					Created:  time.Time{},
				},
				{
					ID:       2,
					BoardID:  101,
					ThreadID: 202,
					Title:    "Title2",
					Text:     "Text2",
					Content:  "Content2",
				},
			},
		},
		{
			name: "2 error, thread not found",
			args: args{
				ctx:      context.Background(),
				boardID:  101,
				threadID: 202,
			},
			ds: func() database.Databaser {
				ds := &mocks.Databaser{}
				ds.On("GetThread", mock.Anything, uint64(101), uint64(202)).Once().Return(nil, sql.ErrNoRows)

				return ds
			}(),
			wantErr: sql.ErrNoRows,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			s := usecase.NewUsecase(tt.ds)
			got, err := s.GetThread(tt.args.ctx, tt.args.boardID, tt.args.threadID)
			if (err != nil || tt.wantErr != nil) && !errors.Is(err, tt.wantErr) {
				t.Errorf("Usecase.GetThread() error = %v, wantErr %v", err, tt.wantErr)

				return
			}

			if !assert.Equal(t, got, tt.want) {
				t.Errorf("Usecase.GetThread() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUsecase_PostThread(t *testing.T) {
	t.Parallel()

	type args struct {
		ctx    context.Context
		thread *models.Message
	}
	tests := []struct {
		name    string
		args    args
		ds      database.Databaser
		want    uint64
		wantErr error
	}{
		{
			name: "1",
			args: args{
				ctx: context.Background(),
				thread: &models.Message{
					BoardID: 101,
					Title:   "Title",
					Text:    "Text",
					Content: "Content",
				},
			},
			ds: func() database.Databaser {
				ds := &mocks.Databaser{}
				ds.On("GetBoard", mock.Anything, uint64(101)).Once().Return(&models.Board{}, nil)
				ds.On("PostThread", mock.Anything, &models.Message{
					BoardID: 101,
					Title:   "Title",
					Text:    "Text",
					Content: "Content",
				}).Once().Return(uint64(1), uint64(303), nil)

				return ds
			}(),
			want: 303,
		},
		{
			name: "2 error, thread not found",
			args: args{
				ctx:    context.Background(),
				thread: &models.Message{BoardID: 101},
			},
			ds: func() database.Databaser {
				ds := &mocks.Databaser{}
				ds.On("GetBoard", mock.Anything, uint64(101)).Once().Return(nil, sql.ErrNoRows)

				return ds
			}(),
			wantErr: sql.ErrNoRows,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			s := usecase.NewUsecase(tt.ds)
			got, err := s.PostThread(tt.args.ctx, tt.args.thread)
			if (err != nil || tt.wantErr != nil) && !errors.Is(err, tt.wantErr) {
				t.Errorf("Usecase.PostThread() error = %v, wantErr %v", err, tt.wantErr)

				return
			}

			if got != tt.want {
				t.Errorf("Usecase.PostThread() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUsecase_PostComment(t *testing.T) {
	t.Parallel()

	type args struct {
		ctx     context.Context
		comment *models.Message
	}
	tests := []struct {
		name    string
		args    args
		ds      database.Databaser
		want    uint64
		wantErr error
	}{
		{
			name: "1",
			args: args{
				ctx: context.Background(),
				comment: &models.Message{
					BoardID:  101,
					ThreadID: 202,
					Title:    "Title",
					Text:     "Text",
					Content:  "Content",
				},
			},
			ds: func() database.Databaser {
				ds := &mocks.Databaser{}
				ds.On("GetThread", mock.Anything, uint64(101), uint64(202)).Once().Return(models.MessageList{}, nil)
				ds.On("PostMessage", mock.Anything, &models.Message{
					BoardID:  101,
					ThreadID: 202,
					Title:    "Title",
					Text:     "Text",
					Content:  "Content",
				}).Once().Return(uint64(303), nil)

				return ds
			}(),
			want: 303,
		},
		{
			name: "2 error, thread not found",
			args: args{
				ctx: context.Background(),
				comment: &models.Message{
					BoardID:  101,
					ThreadID: 202,
				},
			},
			ds: func() database.Databaser {
				ds := &mocks.Databaser{}
				ds.On("GetThread", mock.Anything, uint64(101), uint64(202)).Once().Return(nil, sql.ErrNoRows)

				return ds
			}(),
			wantErr: sql.ErrNoRows,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			s := usecase.NewUsecase(tt.ds)
			got, err := s.PostMessage(tt.args.ctx, tt.args.comment)
			if (err != nil || tt.wantErr != nil) && !errors.Is(err, tt.wantErr) {
				t.Errorf("Usecase.PostMessage() error = %v, wantErr %v", err, tt.wantErr)

				return
			}

			if got != tt.want {
				t.Errorf("Usecase.PostMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}
