package go_mysql

import (
	"context"
	"fmt"
	"testing"
)

func TestQueryParameter(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	username := "admin"
	password := "admin"

	ctx := context.Background()
	sqlQuery := "SELECT username FROM user WHERE username = ? AND password = ?" // Menggunakan ?

	rows, err := db.QueryContext(ctx, sqlQuery, username, password) // Karena variadic, bisa dimasukkan sebanyak-banyaknya
	if err != nil {
		panic(err)
	}

	if rows.Next() {
		var username string
		err := rows.Scan(&username)

		if err != nil {
			panic(err)
		}
		fmt.Println("Sukses")
	} else {
		fmt.Println("Gagal")
	}
}

func TestExecParameter(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	username := "dimas"
	password := "dimas"

	ctx := context.Background()

	sqlQuery := "INSERT INTO user(username, password) VALUES(?, ?)"
	_, err := db.ExecContext(ctx, sqlQuery, username, password)

	if err != nil {
		panic(err)
	}

}
