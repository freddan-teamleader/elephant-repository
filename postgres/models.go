// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package postgres

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Acl struct {
	Uuid        pgtype.UUID
	Uri         string
	Permissions pgtype.Array[string]
}

type Document struct {
	Uuid           pgtype.UUID
	Uri            string
	Created        pgtype.Timestamptz
	CreatorUri     string
	Modified       pgtype.Timestamptz
	CurrentVersion int64
	Deleted        bool
}

type DocumentLink struct {
	FromDocument pgtype.UUID
	Version      int64
	ToDocument   pgtype.UUID
	Rel          pgtype.Text
}

type DocumentStatus struct {
	Uuid       pgtype.UUID
	Name       string
	ID         int64
	Version    int64
	Hash       []byte
	Created    pgtype.Timestamptz
	CreatorUri string
	Meta       []byte
}

type DocumentVersion struct {
	Uuid         pgtype.UUID
	Version      int64
	Hash         []byte
	Title        string
	Type         string
	Language     string
	Created      pgtype.Timestamptz
	CreatorUri   string
	Meta         []byte
	DocumentData []byte
	Archived     bool
}

type SchemaVersion struct {
	Version int32
}

type StatusHead struct {
	Uuid pgtype.UUID
	Name string
	ID   int64
}