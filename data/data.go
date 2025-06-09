package data

import (
	"database/sql"
	"log"
)

type Database struct {
	Handle *sql.DB
}

func Init() Database {
	handle, err := sql.Open("sqlite", "file.db")
	if err != nil {
		log.Fatalf("ERR: Failed to Initialize Database: %s\n", err.Error())
	}

	return Database {
		Handle: handle,
	}
}

func (d *Database) ApplyMigrations(body string) {
	tx, err := d.Handle.Begin();
	if err != nil {
		log.Fatalf("ERR: Failed to start a transaction: %s\n", err.Error())
	}

	_, err = tx.Exec(body)
	if err != nil {
		log.Fatalf("ERR: Failed to apply migrations: %s\n", err.Error())
	}

	_ = tx.Commit();
}

func (d *Database) ExecuteTransaction(query string) (sql.Result, error){
	tx, err := d.Handle.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Commit()

	result, err := tx.Exec(query)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (d *Database) ExecutePreparedTx(query string, args ...any) (sql.Result, error) {
	tx, err := d.Handle.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Commit()

	stmt, err := tx.Prepare(query)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(args...)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (d *Database) QueryPreparedTx(query string, args ...any) (*sql.Rows, error) {
	tx, err := d.Handle.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Commit()

	stmt, err := tx.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	result, err := stmt.Query(args)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (d *Database) QueryRowPreparedTx(query string, args ...any) (*sql.Row, error) {
	tx, err := d.Handle.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Commit()

	stmt, err := tx.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	result := stmt.QueryRow(args)
	return result, nil
}
