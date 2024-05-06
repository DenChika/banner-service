package repository

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	bannersTable = "banners"
	usersTable   = "users"
)

type DbConfig struct {
	User     string
	Password string
	Host     string
	Port     int
	Name     string
	Ssl      string
	Driver   string
}

type Driver interface {
	GetConnectionString(cfg *DbConfig) string
}

func ConnectToDb(cfg *DbConfig, driver Driver) (*sqlx.DB, error) {
	db, err := sqlx.Open(
		cfg.Driver,
		driver.GetConnectionString(cfg))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
