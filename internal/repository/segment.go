package repository

import (
	"context"
	"fmt"
	"github.com/MyLi2tlePony/AvitoInternshipGolang2023/internal/entity"
	"github.com/jackc/pgx/v5/pgxpool"
	"os"
	"time"
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

func (s *SegmentRepository) CreateCSV(fileName string, year, month int) error {
	sql := `SELECT users.user_id, segment_name, operation, operation_time
		FROM segments_users_history
		JOIN users ON segments_users_history.user_id = users.id
		WHERE EXTRACT(YEAR FROM operation_time) = $1 AND EXTRACT(MONTH FROM operation_time) = $2;`

	query, err := s.db.Query(s.ctx, sql, year, month)
	if err != nil {
		return err
	}

	file, err := os.Create(fileName)
	if err != nil {
		return err
	}

	for query.Next() {
		var userID int
		var segment, operation string
		var timestamp time.Time

		err := query.Scan(&userID, &segment, &operation, &timestamp)
		if err != nil {
			return err
		}

		_, err = file.WriteString(fmt.Sprintf("%d;%s;%s;%s\n", userID, segment, operation, timestamp.Format("20060102150405")))
		if err != nil {
			return err
		}
	}

	err = file.Close()
	if err != nil {
		return err
	}

	return nil
}
