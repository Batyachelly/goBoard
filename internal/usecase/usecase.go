package usecase

import (
	"context"
	"fmt"

	"github.com/Batyachelly/goBoard/internal/database"
	"github.com/Batyachelly/goBoard/internal/database/models"
)

type Usecase struct {
	ds database.Databaser
}

func NewUsecase(ds database.Databaser) *Usecase {
	return &Usecase{
		ds: ds,
	}
}

func (s *Usecase) GetBoardList(ctx context.Context) (models.BoardList, error) {
	boardList, err := s.ds.GetBoardList(ctx)
	if err != nil {
		return nil, fmt.Errorf("usecase get boards: %w", err)
	}

	return boardList, nil
}

func (s *Usecase) GetBoard(ctx context.Context, boardID uint64) (*models.Board, error) {
	board, err := s.ds.GetBoard(ctx, boardID)
	if err != nil {
		return nil, fmt.Errorf("usecase get board: %w", err)
	}

	return board, nil
}

func (s *Usecase) GetThread(ctx context.Context, boardID, threadID uint64) (models.MessageList, error) {
	thread, err := s.ds.GetThread(ctx, boardID, threadID)
	if err != nil {
		return nil, fmt.Errorf("usecase get thread: %w", err)
	}

	return thread, nil
}

func (s *Usecase) PostThread(ctx context.Context, thread *models.Message) (uint64, error) {
	if _, err := s.ds.GetBoard(ctx, thread.BoardID); err != nil {
		return 0, fmt.Errorf("usecase is board exists: %w", err)
	}

	_, threadID, err := s.ds.PostThread(ctx, thread)
	if err != nil {
		return 0, fmt.Errorf("usecase post thread: %w", err)
	}

	return threadID, nil
}

func (s *Usecase) PostMessage(ctx context.Context, message *models.Message) (uint64, error) {
	if _, err := s.ds.GetThread(ctx, message.BoardID, message.ThreadID); err != nil {
		return 0, fmt.Errorf("usecase is thread exists: %w", err)
	}

	messageID, err := s.ds.PostMessage(ctx, message)
	if err != nil {
		return 0, fmt.Errorf("usecase post comment: %w", err)
	}

	return messageID, nil
}
