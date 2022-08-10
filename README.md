# Project Akhir DTS PROA Golang Batch 3 Tahun 2022

Project ini menggunakan Go versi 1.90 dan database MySQL

## Stuktur folder project

1.  `/cmd`  
    Tempat entry point aplikasi, main.go ada di folder ini

2.  `/internal/pkg`  
    Tempat package-package controller dan initiator aplikasi

3.  `/web/db`  
    Folder berisi informasi migrasi dan physical model database, file `.mwb` dapat dibuka menggunakan aplikasi MySQL Workbench

4.  `/web/templates`  
    Folder berisi file template untuk template engine Go

5.  `/web/public`  
    Folder berisi dokumen yang dapat diakses user aplikasi web seperti gambar, css, javascript, dll

## Persiapan Local Dev dan Database

1.  Install dependency project menggunakan perintah:

        go mod tidy

2.  Inisialisasikan database dengan menggunakan query migrasi di `/web/db/migration.sql`

3.  Sesuaikan konfigurasi database aplikasi di `/internal/pkg/commons/init.go`, ganti baris di bawah ini sesuai dengan konfigurasi database anda

        db, err := sql.Open("mysql", "root@(127.0.0.1:3306)/test_gotodo?parseTime=true")

4.  Disarankan mengnstall package `github.com/cosmtrek/air` untuk dapat menggunakan fasilitas hot reload

        go install github.com/cosmtrek/air@latest

    Setelah instalasi package server dan hot reload dapat dijalankan menggunakan perintah ini di folder project

        air

5.  Aplikasi bisa diakses di `http://127.0.0.1:8000/` atau sesuai konfigurasi di `/internal/pkg/commons/init.go`
