package repository

import (
	"context"
	"github.com/MyLi2tlePony/AvitoInternshipGolang2023/internal/entity"
	"github.com/jackc/pgx/pgtype"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	db  *pgxpool.Pool
	ctx context.Context
}

func NewUserRepository(ctx context.Context, db Database) *UserRepository {
	return &UserRepository{
		db:  db.Connection(),
		ctx: ctx,
	}
}

func (r *UserRepository) ChangeSegments(userID int, insert, delete []string) error {
	sql := `INSERT INTO users (user_id) VALUES ($1) ON CONFLICT DO NOTHING RETURNING id;`

	var id int
	err := r.db.QueryRow(r.ctx, sql, userID).Scan(&id)
	if err != nil && err != pgx.ErrNoRows {
		return err
	}

	if id == 0 {
		sql = `SELECT id FROM users WHERE user_id = $1;`

		err := r.db.QueryRow(r.ctx, sql, userID).Scan(&id)
		if err != nil {
			return err
		}
	}

	params := &pgtype.TextArray{}
	if err = params.Set(insert); err != nil {
		return err
	}

	sql = `SELECT id FROM segments WHERE name = ANY ($1);`

	query, err := r.db.Query(r.ctx, sql, params)
	if err != nil {
		return err
	}

	for query.Next() {
		var insertID int

		if err = query.Scan(&insertID); err != nil {
			return err
		}

		sql = `INSERT INTO segments_users (user_id, segment_id) VALUES ($1, $2);`

		_, err = r.db.Exec(r.ctx, sql, id, insertID)
		if err != nil {
			return err
		}
	}

	params = &pgtype.TextArray{}
	if err = params.Set(delete); err != nil {
		return err
	}

	sql = `SELECT id FROM segments WHERE name = ANY ($1);`

	query, err = r.db.Query(r.ctx, sql, params)
	if err != nil {
		return err
	}

	for query.Next() {
		var deleteID int

		if err = query.Scan(&deleteID); err != nil {
			return err
		}

		sql = `DELETE FROM segments_users WHERE user_id = $1 AND segment_id = $2;`

		_, err = r.db.Exec(r.ctx, sql, id, deleteID)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *UserRepository) GetSegments(userID int) ([]entity.Segment, error) {
	sql := `SELECT segments.name FROM segments_users 
		JOIN users ON segments_users.user_id = users.id
		JOIN segments ON segments.id = segments_users.segment_id
		WHERE users.user_id = $1`

	query, err := r.db.Query(r.ctx, sql, userID)
	if err != nil {
		return nil, err
	}

	var segments []entity.Segment
	for query.Next() {
		segment := entity.Segment{}

		err = query.Scan(&segment.Name)
		if err != nil {
			return nil, err
		}

		segments = append(segments, segment)
	}

	return segments, nil
}
