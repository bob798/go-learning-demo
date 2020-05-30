package session

import (
	"database/sql"
	"strings"

	"github.com/bob798/go-learning-demo/go-orm/day2-reflect-schema/log"

	"github.com/bob798/go-learning-demo/go-orm/day2-reflect-schema/dialect"
	"github.com/bob798/go-learning-demo/go-orm/day2-reflect-schema/schema"
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

func (s *Session) DB() *sql.DB {
	return s.DB()
}

func (s *Session) Clear() {
	s.sql.Reset()
	s.sqlVars = nil
}

func (s *Session) Raw(sql string, values ...interface{}) *Session {
	s.sql.WriteString(sql)
	s.sql.WriteString(" ")
	s.sqlVars = append(s.sqlVars, values...)
	return s

}

func (s *Session) Exec() (result sql.Result, err error) {
	log.Info(s.sql, s.sqlVars)
	if result, err = s.DB().Exec(s.sql.String(), s.sqlVars...); err != nil {
		log.Error(err)
	}

	return

}

func (s *Session) QueryRaw() *sql.Row {
	defer s.Clear()
	log.Info(s.sql.String(), s.sqlVars)
	return s.DB().QueryRow(s.sql.String(), s.sqlVars...)
}
