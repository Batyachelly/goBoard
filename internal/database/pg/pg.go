package pg

import (
	"context"
	"fmt"

	"github.com/Batyachelly/goBoard/internal/database/models"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/jackc/tern/migrate"
)

type DatabaseService struct {
	pool         *pgxpool.Pool
	versionTable string
}

type Config struct {
	Host         string
	Port         int
	User         string
	Password     string
	DB           string
	SSLMode      string
	VersionTable string
}

func NewDatabaseService(cfg Config) (*DatabaseService, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DB, cfg.SSLMode)

	poolCfg, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, fmt.Errorf("pg parse config: %w", err)
	}

	pool, err := pgxpool.ConnectConfig(context.Background(), poolCfg)
	if err != nil {
		return nil, fmt.Errorf("pg connect config: %w", err)
	}

	return &DatabaseService{
		pool:         pool,
		versionTable: cfg.VersionTable,
	}, nil
}

func (ds *DatabaseService) Migrate() error {
	ctx := context.Background()

	conn, err := ds.pool.Acquire(ctx)
	if err != nil {
		return fmt.Errorf("pg acquire connection for migration: %w", err)
	}

	defer conn.Conn().Close(ctx)

	m, err := migrate.NewMigratorEx(ctx, conn.Conn(), ds.versionTable, nil)
	if err != nil {
		return fmt.Errorf("pg create migrator: %w", err)
	}

	if err := m.Migrate(ctx); err != nil {
		return fmt.Errorf("pg try to migrate: %w", err)
	}

	return nil
}

func (ds *DatabaseService) GetBoardList(ctx context.Context) (models.BoardList, error) {
	rows, err := ds.pool.Query(ctx, "select id, title from board where status>0")
	if err != nil {
		return nil, fmt.Errorf("pg select boards: %w", err)
	}
	defer rows.Close()

	boards := models.BoardList{}

	for rows.Next() {
		b := models.Board{}

		if err := rows.Scan(&b.ID, &b.Title); err != nil {
			return nil, fmt.Errorf("pg scan boards: %w", err)
		}

		boards = append(boards, b)
	}

	return boards, nil
}

func (ds *DatabaseService) GetBoard(ctx context.Context, boardID uint64) (*models.Board, error) {
	tx, err := ds.pool.Begin(ctx)
	if err != nil {
		return nil, fmt.Errorf("pg start tx for get board: %w", err)
	}

	defer tx.Rollback(ctx) //nolint:errcheck

	board := new(models.Board)

	board.ID = boardID

	{
		row := ds.pool.QueryRow(ctx, "select id, title from board where status>0 and id=$1 limit 1", boardID)

		if err := row.Scan(&board.ID, &board.Title); err != nil {
			return nil, fmt.Errorf("pg select board: %w", err)
		}
	}

	{
		rows, err := tx.Query(ctx, "select id, title, text, content, created from message where status>0 and board_id=$1", boardID)
		if err != nil {
			return nil, fmt.Errorf("pg select board threads: %w", err)
		}
		defer rows.Close()

		for rows.Next() {
			m := models.Message{}

			if err := rows.Scan(&m.ID, &m.Title, &m.Text, &m.Content, &m.Created); err != nil {
				return nil, fmt.Errorf("pg scan messages: %w", err)
			}

			board.Threads = append(board.Threads, m)
		}
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, fmt.Errorf("pg commit tx for get board: %w", err)
	}

	return board, nil
}

func (ds *DatabaseService) GetThread(ctx context.Context, boardID, threadID uint64) (models.MessageList, error) {
	rows, err := ds.pool.Query(ctx, "select id, title, text, content, created from message where status>0 and board_id=$1 and thread_id=$2", boardID, threadID)
	if err != nil {
		return nil, fmt.Errorf("pg select comments: %w", err)
	}

	defer rows.Close()

	messages := models.MessageList{}

	for rows.Next() {
		m := models.Message{}

		if err := rows.Scan(&m.ID, &m.Title, &m.Text, &m.Content, &m.Created); err != nil {
			return nil, fmt.Errorf("pg scan message: %w", err)
		}

		messages = append(messages, m)
	}

	return messages, nil
}

func (ds *DatabaseService) PostThread(ctx context.Context, thread *models.Message) (uint64, uint64, error) {
	row := ds.pool.QueryRow(ctx, "insert into message (status, board_id, title, text, content) values (1, $1, $2, $3, $4) returning id, thread_id",
		thread.BoardID, thread.Title, thread.Text, thread.Content)

	var id, threadID uint64

	if err := row.Scan(&id, &threadID); err != nil {
		return 0, 0, fmt.Errorf("pg insert thread: %w", err)
	}

	return id, threadID, nil
}

func (ds *DatabaseService) PostMessage(ctx context.Context, message *models.Message) (uint64, error) {
	row := ds.pool.QueryRow(ctx, "insert into message (status, board_id, thread_id, title, text, content) values (1, $1, $2, $3, $4, $5) returning id",
		message.BoardID, message.ThreadID, message.Title, message.Text, message.Content)

	var id uint64

	if err := row.Scan(&id); err != nil {
		return 0, fmt.Errorf("pg insert message: %w", err)
	}

	return id, nil
}
