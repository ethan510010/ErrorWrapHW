package db

import "database/sql"

func Query(query string) (interface{}, error) {
	return nil, sql.ErrNoRows
}