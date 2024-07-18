package go_mysql

import (
	"context"
	"fmt"
	"testing"
)

func TestAutoIncrement(t *testing.T) {
	db := GetConnection()
	ctx := context.Background()

	email := "dimas@gmail.com"
	comment := "Test"
	sqlQuery := "INSERT INTO comments(email, comment) VALUES(?, ?)"
	result, err := db.ExecContext(ctx, sqlQuery, email, comment)

	if err != nil {
		panic(err)
	}

	insertId, err := result.LastInsertId()

	if err != nil {
		panic(err)
	}
	fmt.Println(insertId)
}
