package go_mysql

import (
	"database/sql"
	"time"
)

func GetConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:ramadhan28@tcp(localhost:3306)/belajar_mysql?parseTime=true")
	if err != nil {
		panic(err)
	}
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}

/*
 * db.SetMaxIdleConns(): Jumlah koneksi minimal
 *	Saat pertama kali aplikasi digunakan, akan memanggil minimal koneksi agar jika trafic naik tidak menunggu
 * db.SetMaxOpenConns(): Jumlah koneksi maksimal
 *	Jika maksimal sudah digunakan, sisanya harus menunggu
 * db.SetMaxIdleTime() : Lama koneksi tidak digunakan akan dihapus
 *	Jika minimal sudah di set, tetapi digunakan ada lebih dari minimal kemudian tidak digunakan, akan di close koneksi sesuai paramater yang ditentukan
 * db.SetMaxLifeTime(): Berapa lama koneksi boleh digunakan
 *	Setelah beberapa waktu akan dibuat baru
 */
