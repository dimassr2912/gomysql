package go_mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql" // Untuk memanggil init
	"testing"
)

func TestEmpty(t *testing.T) {

}

// Membuat koneksi ke db
func TestOpenConnection(t *testing.T) {
	// membuat object sql.DB menggunakan function sql.Open
	db, err := sql.Open("mysql", "root:ramadhan28@tcp(localhost:3306)/belajar_mysql")
	if err != nil {
		panic(err)
	}
	defer db.Close() // Tutup setelah digunakan
}

/*
 * Menghubungkan dengan database
 *
 */
