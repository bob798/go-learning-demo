package geeorm

import (
	"database/sql"
	"github.com/bob798/go-learning-demo/go-orm/day1-database-sql/log"
	"github.com/bob798/go-learning-demo/go-orm/day1-database-sql/session"
)

type Engine struct {
	db *sql.DB
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
	e = &Engine{db: db}
	log.Info("Connect database success")
	return
}

func (engine *Engine) Close() {
	if err := engine.db.Close(); err != nil {
		log.Error("Failed to close database")
	}
	log.Info("Close database success")

}

func (engine *Engine) NewSession() *session.Session {
	return session.New(engine.db)
}