package dataservice

import (
	"context"
	"database/sql"
)


func InsertStudent(ctx context.Context, db *sql.DB, name string, age int, grade string) error {
	_, err := db.ExecContext(
		ctx,
		"INSERT INTO students (name, age, grade) VALUES (?, ?, ?)",
		name, age, grade,
	)
	return err
}
	