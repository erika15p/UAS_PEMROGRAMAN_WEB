package models

import "time"

type Kehadiran struct {
	ID        uint   `gorm:"primaryKey"`
	Tanggal   string `form:"tanggal"`
	Nama      string `form:"nama"`
	NPM       string `form:"npm"`
	Prodi     string `form:"prodi"`
	Status    string `form:"status"`
	CreatedAt time.Time
}

type Keuangan struct {
	ID        uint      `gorm:"primaryKey"`
	Tanggal   time.Time `form:"-"`
	Deskripsi string    `form:"deskripsi"`
	Tipe      string    `form:"tipe"`
	Jumlah    float64   `form:"jumlah"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
