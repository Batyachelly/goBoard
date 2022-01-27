package database

import (
	"context"

	"github.com/Batyachelly/goBoard/internal/database/models"
)

//go:generate mockery --name=Databaser --output=./../../generated/mocks
type Databaser interface {
	Migrate() error
	GetBoardList(ctx context.Context) (models.BoardList, error)
	GetBoard(ctx context.Context, boardID uint64) (*models.Board, error)
	GetThread(ctx context.Context, boardID, threadID uint64) (models.MessageList, error)
	PostThread(ctx context.Context, thread *models.Message) (uint64, uint64, error)
	PostMessage(ctx context.Context, message *models.Message) (uint64, error)
}
