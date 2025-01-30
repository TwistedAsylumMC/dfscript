package dfscript

import (
	"database/sql"
	"fmt"
	"github.com/dop251/goja"
	"github.com/google/uuid"
	"os"
	"time"

	_ "github.com/glebarez/go-sqlite"
	_ "github.com/go-sql-driver/mysql"
)

type SqlDB struct {
	db *sql.DB
}

func (s SqlDB) Begin() (SqlTx, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return SqlTx{}, err
	}
	return SqlTx{tx: tx}, nil
}

func (s SqlDB) Close() error {
	return s.db.Close()
}

func (s SqlDB) Exec(query string, args ...any) (sql.Result, error) {
	res, err := s.db.Exec(query, args...)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s SqlDB) Ping() error {
	return s.db.Ping()
}

func (s SqlDB) Prepare(query string) (SqlStmt, error) {
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return SqlStmt{}, err
	}
	return SqlStmt{stmt: stmt}, nil
}

func (s SqlDB) Query(query string, args ...any) ([]map[string]any, error) {
	rows, err := s.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	return parseSqlRows(rows)
}

func (s SqlDB) QueryRow(query string, args ...any) (map[string]any, error) {
	rows, err := s.Query(query, args...)
	if err != nil {
		return nil, err
	} else if len(rows) == 0 {
		return nil, nil
	}
	return rows[0], nil
}

func (s SqlDB) SetConnMaxIdleTime(d time.Duration) {
	s.db.SetConnMaxIdleTime(d)
}

func (s SqlDB) SetConnMaxLifetime(d time.Duration) {
	s.db.SetConnMaxLifetime(d)
}

func (s SqlDB) SetMaxIdleConns(n int) {
	s.db.SetMaxIdleConns(n)
}

func (s SqlDB) SetMaxOpenConns(n int) {
	s.db.SetMaxOpenConns(n)
}

func (s SqlDB) Stats() sql.DBStats {
	return s.db.Stats()
}

type SqlStmt struct {
	stmt *sql.Stmt
}

func (s SqlStmt) Close() error {
	return s.stmt.Close()
}

func (s SqlStmt) Exec(args ...any) (sql.Result, error) {
	return s.stmt.Exec(args...)
}

func (s SqlStmt) Query(args ...any) ([]map[string]any, error) {
	rows, err := s.stmt.Query(args...)
	if err != nil {
		return nil, err
	}
	return parseSqlRows(rows)
}

func (s SqlStmt) QueryRow(args ...any) map[string]any {
	rows, err := s.Query(args...)
	if err != nil {
		return nil
	}
	if len(rows) == 0 {
		return nil
	}
	return rows[0]
}

type SqlTx struct {
	tx *sql.Tx
}

func (t SqlTx) Commit() error {
	return t.tx.Commit()
}

func (t SqlTx) Exec(query string, args ...any) (sql.Result, error) {
	return t.tx.Exec(query, args...)
}

func (t SqlTx) Prepare(query string) (SqlStmt, error) {
	stmt, err := t.tx.Prepare(query)
	if err != nil {
		return SqlStmt{}, err
	}
	return SqlStmt{stmt: stmt}, nil
}

func (t SqlTx) Query(query string, args ...any) ([]map[string]any, error) {
	rows, err := t.tx.Query(query, args)
	if err != nil {
		return nil, err
	}
	return parseSqlRows(rows)
}

func (t SqlTx) QueryRow(query string, args ...any) (map[string]any, error) {
	rows, err := t.Query(query, args...)
	if err != nil {
		return nil, err
	} else if len(rows) == 0 {
		return nil, nil
	}
	return rows[0], nil
}

func (t SqlTx) Rollback() error {
	return t.tx.Rollback()
}

func (t SqlTx) Stmt(stmt SqlStmt) SqlStmt {
	return SqlStmt{stmt: t.tx.Stmt(stmt.stmt)}
}

func parseSqlRows(rows *sql.Rows) ([]map[string]any, error) {
	cols, err := rows.Columns()
	if err != nil {
		return nil, err
	}
	var res []map[string]any
	for rows.Next() {
		result := make([]any, len(cols))
		scanArgs := make([]any, len(cols))
		for i := range scanArgs {
			scanArgs[i] = &result[i]
		}
		if err := rows.Scan(scanArgs...); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		m := make(map[string]any)
		for i, col := range cols {
			m[col] = result[i]
		}
		res = append(res, m)
	}
	return res, nil
}

func (r *Runtime) setupStd() error {
	err := newObject().
		Method("open", func(c goja.FunctionCall) goja.Value {
			driver, ok := c.Argument(0).Export().(string)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not a string, got %T", c.Argument(0).Export()))
			}
			dsn, ok := c.Argument(1).Export().(string)
			if !ok {
				panic(r.vm.NewTypeError("argument 1 is not a string, got %T", c.Argument(1).Export()))
			}
			db, err := sql.Open(driver, dsn)
			if err != nil {
				panic(r.vm.ToValue(err))
			}
			return r.vm.ToValue(SqlDB{db: db})
		}).
		Apply(r.vm, "sql")
	if err != nil {
		return err
	}

	err = newObject().
		Const("nanosecond", time.Nanosecond).
		Const("microsecond", time.Microsecond).
		Const("millisecond", time.Millisecond).
		Const("second", time.Second).
		Const("minute", time.Minute).
		Const("hour", time.Hour).
		Apply(r.vm, "time")
	if err != nil {
		return err
	}

	err = newObject().
		Method("readFile", func(c goja.FunctionCall) goja.Value {
			path := c.Argument(0).String()
			data, err := os.ReadFile(path)
			if err != nil {
				panic(r.vm.ToValue(err))
			}
			return r.vm.ToValue(string(data))
		}).
		Method("writeFile", func(c goja.FunctionCall) goja.Value {
			path := c.Argument(0).String()
			data := c.Argument(1).String()
			err := os.WriteFile(path, []byte(data), 0644)
			if err != nil {
				panic(r.vm.ToValue(err))
			}
			return nil
		}).
		Apply(r.vm, "os")
	if err != nil {
		return err
	}

	err = newObject().
		Method("create", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(uuid.New())
		}).
		Method("parse", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(uuid.MustParse(c.Argument(0).String()))
		}).
		Apply(r.vm, "uuid")
	return nil
}
