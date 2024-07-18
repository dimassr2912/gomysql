package go_mysql

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"
)

func TestExecSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "INSERT INTO customer(id, name, email, balance, rating, birth_date, married) VALUES('dimas', 'Dimas', 'dimas@gmail.com', 100000000, 5.0, '1998-12-29', false)"
	_, err := db.ExecContext(ctx, script)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new customer")
}

func TestQuerySql(t *testing.T) {
	db := GetConnection()

	ctx := context.Background()
	rows, err := db.QueryContext(ctx, "SELECT id, name, email, balance, rating, birth_date, married, created_at FROM customer") // Untuk melakukan query
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() { // Untuk melakukan iterasi terhadap hasil query
		var (
			id, name   string
			email      sql.NullString
			balance    int32
			rating     float32
			married    bool
			birth_date sql.NullTime
			created_at time.Time
		)
		err = rows.Scan(&id, &name, &email, &balance, &rating, &birth_date, &married, &created_at) // rows.Scan(): Untuk membaca data
		if err != nil {
			panic(err)
		}
		fmt.Println("Id:", id)
		fmt.Println("Name:", name)
		if email.Valid {
			fmt.Println("Email:", email.String)
		}
		fmt.Println("Balance:", balance)
		fmt.Println("Rating:", rating)
		if birth_date.Valid {
			fmt.Println("Birth Date:", birth_date.Time)
		}
		fmt.Println("Married:", married)
		fmt.Println("Create at:", created_at)
	}
}
