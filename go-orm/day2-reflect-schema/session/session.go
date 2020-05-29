package session

import (
	"database/sql"
	"github.com/bob798/go-learning-demo/go-orm/day2-reflect-schema/dialect"
	"github.com/bob798/go-learning-demo/go-orm/day2-reflect-schema/schema"
	"strings"
)

type Session struct {
	db       *sql.DB
	dialect  dialect.Dialect
	refTable *schema.Schema
	sql      strings.Builder
	sqlVars  []interface{}
}

func New(db *sql.DB, dialect dialect.Dialect) *Session {
	return &Session{
		db:      db,
		dialect: dialect,
	}
}
