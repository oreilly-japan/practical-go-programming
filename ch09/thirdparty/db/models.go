// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"time"
)

type Author struct {
	ID        int32
	Name      string
	CreatedAt time.Time
}

type Book struct {
	ID        int32
	Title     string
	Price     int32
	AuthorID  int32
	CreatedAt time.Time
}
