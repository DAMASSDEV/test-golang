package configs

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB() {
    var err error

    // Tambahkan parameter parseTime=true di string koneksi
    // untuk menangani konversi waktu MySQL ke Go time.Time
    DB, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/event_realm?parseTime=true")
    if err != nil {
        log.Fatal("Gagal membuka koneksi database:", err)
    }

    // Konfigurasi pool koneksi
    DB.SetMaxOpenConns(10)
    DB.SetMaxIdleConns(5)
    DB.SetConnMaxLifetime(time.Minute * 5)

    // Verifikasi koneksi dengan ping
    err = DB.Ping()
    if err != nil {
        log.Fatal("melakukan ping ke database:", err)
    }

    // Verifikasi tabel events
    _, err = DB.Query("SELECT 1 FROM events LIMIT 1")
    if err != nil {
        log.Fatal("Gagal menjalankan query pada tabel events:", err)
    }

    log.Println("Berhasil terhubung ke database")
}