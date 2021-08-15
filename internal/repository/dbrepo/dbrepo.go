package dbrepo

import (
	"database/sql"
	"github.com/tsawler/bookings-app/internal/config"
)

type postgresDBRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

func NewPostgresRepo(conn *sql.DB, a *config.AppConfig) *postgresDBRepo {
	return &postgresDBRepo{
		App: a,
		DB:  conn,
	}
}
