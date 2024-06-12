package handler

import (
	"auth/postgres"
	"database/sql"
)

type Handler struct{
	Db *sql.DB
	UserDb *postgres.UserRepo
}

func NewHandler(db *sql.DB,user *postgres.UserRepo) *Handler {
	return &Handler{
		Db: db,
		UserDb: user,
	}
}

