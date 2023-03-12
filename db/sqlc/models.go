// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2

package db

import (
	"database/sql"
)

type Todo struct {
	ID        int64        `json:"id"`
	Title     string       `json:"title"`
	Completed bool         `json:"completed"`
	CreateAt  sql.NullTime `json:"create_at"`
	UpdateAt  sql.NullTime `json:"update_at"`
	DeleteAt  sql.NullTime `json:"delete_at"`
}