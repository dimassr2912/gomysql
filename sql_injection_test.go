package go_mysql

import (
	"context"
	"fmt"
	"testing"
)

func TestSqlIjection(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "admin'; #" // SQL Injection (Setelah # diabaikan, karena # adalah komentar)
	password := "admin"

	script := "SELECT username FROM user WHERE username = '" + username + "' AND password = '" + password + "' LIMIT 1 "
	fmt.Println(script)
	rows, err := db.QueryContext(ctx, script)

	if err != nil {
		panic(err)
	}

	if rows.Next() {
		var username string
		rows.Scan(&username)
		fmt.Println("Sukses")
	} else {
		fmt.Println("Gagal")
	}
}
