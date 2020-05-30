package geeorm

import (
	"database/sql"

	"github.com/bob798/go-learning-demo/go-orm/day2-reflect-schema/dialect"
	"github.com/bob798/go-learning-demo/go-orm/day2-reflect-schema/log"
	"github.com/bob798/go-learning-demo/go-orm/day2-reflect-schema/session"
)

type Engine struct {
	db      *sql.DB
	dialect dialect.Dialect
}

func NewEngine(driver, source string) (e *Engine, err error) {
	db, err := sql.Open(driver, source)
	if err != nil {
		log.Error(err)
		return
	}

	if err = db.Ping(); err != nil {
		log.Error(err)
		return
	}
	dial, ok := dialect.GetDialect(driver)
	if !ok {
		log.Errorf("dialect %s Not Found", driver)
		return
	}
	e = &Engine{db: db, dialect: dial}
	log.Info("Connect database success")
	return
}

func (engine *Engine) NewSession() *session.Session {
	return session.New(engine.db, engine.dialect)
}
