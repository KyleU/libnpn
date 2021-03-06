package npndatabase

import (
	"fmt"

	"emperror.dev/errors"
	"github.com/jmoiron/sqlx"
)

// Runs a SQL query, returning a resultset or error
func (s *Service) Query(q string, tx *sqlx.Tx, values ...interface{}) (*sqlx.Rows, error) {
	if s.debug {
		logQuery(s, "running raw query", q, values)
	}
	if tx == nil {
		return s.db.Queryx(q, values...)
	}
	return tx.Queryx(q, values...)
}

// Runs a SQL query and parses into the "dest" argument for all rows, returning an optional error
func (s *Service) Select(dest interface{}, q string, tx *sqlx.Tx, values ...interface{}) error {
	if s.debug {
		logQuery(s, fmt.Sprintf("selecting rows of type [%T]", dest), q, values)
	}
	if tx == nil {
		return s.db.Select(dest, q, values...)
	}
	return tx.Select(dest, q, values...)
}

// Runs a SQL query for a single row and parses into the "dest" argument, returning an optional error
func (s *Service) Get(dto interface{}, q string, tx *sqlx.Tx, values ...interface{}) error {
	if s.debug {
		logQuery(s, fmt.Sprintf("getting single row of type [%T]", dto), q, values)
	}
	if tx == nil {
		return s.db.Get(dto, q, values...)
	}
	return tx.Get(dto, q, values...)
}

type singleIntResult struct {
	X *int64 `db:"x"`
}

// Runs a SQL query for a single integer return value (like counts), returning that int and an optional error
func (s *Service) SingleInt(q string, tx *sqlx.Tx, values ...interface{}) (int64, error) {
	x := &singleIntResult{}
	var err error
	if tx == nil {
		err = s.db.Get(x, q, values...)
	} else {
		err = tx.Get(x, q, values...)
	}
	if err != nil {
		return -1, errors.Wrap(err, "returned value is not an integer")
	}
	if x.X == nil {
		return 0, nil
	}
	return *x.X, nil
}
