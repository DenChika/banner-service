package postgres

import (
	"banner_service/internal/repository"
	"fmt"
)

type Driver struct{}

func NewDriver() *Driver {
	return &Driver{}
}

func (pd Driver) GetConnectionString(cfg *repository.DbConfig) string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=%s",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name, cfg.Ssl)
}
