package statements

import (
	"context"
	"database/sql"
	"fmt"
)

func GetAllTablesStmnt(ctx context.Context, db *sql.DB) (*sql.Stmt, error) {
	statement := fmt.Sprintf("select * from information_schema.tables")
	stmt, err := db.PrepareContext(ctx, statement)
	return stmt, err
}

func GetColumnsOfTable(ctx context.Context, db *sql.DB) (*sql.Stmt, error) {
	statement := fmt.Sprintf("select * from information_schema.columns where TABLE_SCHEMA=? and TABLE_NAME=?")
	stmt, err := db.PrepareContext(ctx, statement)
	return stmt, err
}
