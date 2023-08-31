package repository

import "github.com/jackc/pgx/v5/pgxpool"

type Database interface {
	Connection() *pgxpool.Pool
}
