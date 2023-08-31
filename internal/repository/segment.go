package repository

import (
	"context"

	"github.com/MyLi2tlePony/AvitoInternshipGolang2023/internal/entity"
	"github.com/jackc/pgx/v5/pgxpool"
)

type SegmentRepository struct {
	db  *pgxpool.Pool
	ctx context.Context
}

func NewSegmentRepository(ctx context.Context, db Database) *SegmentRepository {
	return &SegmentRepository{
		db:  db.Connection(),
		ctx: ctx,
	}
}

func (s *SegmentRepository) Create(segment *entity.Segment) error {
	sql := `INSERT INTO segments (name) VALUES ($1);`

	_, err := s.db.Exec(s.ctx, sql, segment.Name)
	if err != nil {
		return err
	}

	return nil
}

func (s *SegmentRepository) Delete(segment *entity.Segment) error {
	sql := `DELETE FROM segments WHERE name = $1;`

	_, err := s.db.Exec(s.ctx, sql, segment.Name)
	if err != nil {
		return err
	}

	return nil
}
