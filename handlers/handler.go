package handlers

import (
	"badminton-app/database"
	"badminton-app/models"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jung-kurt/gofpdf"
)

func ShowHome(c *gin.Context) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	err := tmpl.ExecuteTemplate(c.Writer, "index.html", gin.H{
		"title": "Beranda",
	})
	if err != nil {
		c.String(http.StatusInternalServerError, "Template error: %v", err)
	}
}

// KEHADIRAN
func ListKehadiran(c *gin.Context) {
	var data []models.Kehadiran
	database.DB.Order("tanggal desc").Find(&data)
	isLoggedIn, exists := c.Get("IsLoggedIn")
	if !exists {
		isLoggedIn = false
	}
	c.HTML(http.StatusOK, "kehadiran.html", gin.H{"data": data, "IsLoggedIn": isLoggedIn})
}

func AddKehadiran(c *gin.Context) {
	var input models.Kehadiran

	if err := c.ShouldBind(&input); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("Form tidak valid: %v", err))
		return
	}

	parsedDate, err := time.Parse("2006-01-02", input.Tanggal)
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("Format tanggal salah: %v", err))
		return
	}

	data := models.Kehadiran{
		Tanggal:   parsedDate.Format("2006-01-02"),
		Nama:      input.Nama,
		NPM:       input.NPM,
		Prodi:     input.Prodi,
		Status:    input.Status,
		CreatedAt: time.Now(),
	}

	if err := database.DB.Create(&data).Error; err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("Gagal simpan ke DB: %v", err))
		return
	}

	c.Redirect(http.StatusSeeOther, "/kehadiran")
}

func UpdateKehadiran(c *gin.Context) {
	id := c.Param("id")
	fmt.Println("UPDATE handler masuk untuk ID:", id)

	var hadir models.Kehadiran
	if err := database.DB.First(&hadir, id).Error; err != nil {
		c.String(http.StatusNotFound, "Data tidak ditemukan")
		return
	}

	tanggal := c.PostForm("tanggal")
	nama := c.PostForm("nama")
	npm := c.PostForm("npm")
	prodi := c.PostForm("prodi")
	status := c.PostForm("status")

	parsedTanggal, err := time.Parse("2006-01-02", tanggal)
	if err != nil {
		c.String(http.StatusBadRequest, "Format tanggal tidak valid")
		return
	}

	hadir.Tanggal = parsedTanggal.Format("2006-01-02")
	hadir.Nama = nama
	hadir.NPM = npm
	hadir.Prodi = prodi
	hadir.Status = status

	if err := database.DB.Save(&hadir).Error; err != nil {
		c.String(http.StatusInternalServerError, "Gagal update kehadiran")
		return
	}

	fmt.Println("Data berhasil diupdate:", hadir)
	c.Redirect(http.StatusSeeOther, "/kehadiran")
}

func DeleteKehadiran(c *gin.Context) {
	id := c.Param("id")
	database.DB.Delete(&models.Kehadiran{}, id)
	c.Redirect(http.StatusSeeOther, "/kehadiran")
}

// KEUANGAN
func ListKeuangan(c *gin.Context) {
	var data []models.Keuangan
	database.DB.Order("tanggal desc").Find(&data)
	isLoggedIn, exists := c.Get("IsLoggedIn")
	if !exists {
		isLoggedIn = false
	}
	c.HTML(http.StatusOK, "keuangan.html", gin.H{"data": data, "IsLoggedIn": isLoggedIn})
}

func AddKeuangan(c *gin.Context) {
	var form struct {
		Tanggal   string  `form:"tanggal"`
		Deskripsi string  `form:"deskripsi"`
		Tipe      string  `form:"tipe"`
		Jumlah    float64 `form:"jumlah"`
	}

	if err := c.ShouldBind(&form); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("Form tidak valid: %v", err))
		return
	}

	parsedDate, err := time.Parse("2006-01-02", form.Tanggal)
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("Format tanggal salah: %v", err))
		return
	}

	data := models.Keuangan{
		Tanggal:   parsedDate,
		Deskripsi: form.Deskripsi,
		Tipe:      form.Tipe,
		Jumlah:    form.Jumlah,
	}

	if err := database.DB.Create(&data).Error; err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("Gagal simpan ke DB: %v", err))
		return
	}

	c.Redirect(http.StatusSeeOther, "/keuangan")
}

func UpdateKeuangan(c *gin.Context) {
	id := c.Param("id")
	var kas models.Keuangan

	if err := database.DB.First(&kas, id).Error; err != nil {
		c.String(http.StatusNotFound, "Data tidak ditemukan")
		return
	}

	tanggal := c.PostForm("tanggal")
	deskripsi := c.PostForm("deskripsi")
	tipe := c.PostForm("tipe")
	jumlahStr := c.PostForm("jumlah")

	parsedTanggal, err := time.Parse("2006-01-02", tanggal)
	if err != nil {
		c.String(http.StatusBadRequest, "Format tanggal tidak valid")
		return
	}

	jumlah, err := strconv.ParseFloat(jumlahStr, 64)
	if err != nil {
		c.String(http.StatusBadRequest, "Jumlah tidak valid")
		return
	}

	kas.Tanggal = parsedTanggal
	kas.Deskripsi = deskripsi
	kas.Tipe = tipe
	kas.Jumlah = jumlah

	if err := database.DB.Save(&kas).Error; err != nil {
		c.String(http.StatusInternalServerError, "Gagal menyimpan perubahan")
		return
	}

	c.Status(http.StatusOK)
}

func DeleteKeuangan(c *gin.Context) {
	id := c.Param("id")
	database.DB.Delete(&models.Keuangan{}, id)
	c.Redirect(http.StatusSeeOther, "/keuangan")
}

// LAPORAN
func ShowLaporan(c *gin.Context) {
	isLoggedIn, exists := c.Get("IsLoggedIn")
	if !exists {
		isLoggedIn = false
	}
	c.HTML(http.StatusOK, "laporan.html", gin.H{"IsLoggedIn": isLoggedIn})
}

func GenerateLaporan(c *gin.Context) {
	bulan := c.PostForm("bulan")

	var kehadiranCount int64
	var totalMasuk float64
	var totalKeluar float64

	database.DB.Model(&models.Kehadiran{}).
		Where("DATE_FORMAT(tanggal, '%Y-%m') = ?", bulan).
		Count(&kehadiranCount)

	database.DB.Model(&models.Keuangan{}).
		Where("DATE_FORMAT(tanggal, '%Y-%m') = ? AND tipe = ?", bulan, "Pemasukan").
		Select("SUM(jumlah)").Scan(&totalMasuk)

	database.DB.Model(&models.Keuangan{}).
		Where("DATE_FORMAT(tanggal, '%Y-%m') = ? AND tipe = ?", bulan, "Pengeluaran").
		Select("SUM(jumlah)").Scan(&totalKeluar)

	isLoggedIn, exists := c.Get("IsLoggedIn")
	if !exists {
		isLoggedIn = false
	}

	c.HTML(http.StatusOK, "laporan.html", gin.H{
		"Bulan":       bulan,
		"Kehadiran":   kehadiranCount,
		"TotalMasuk":  totalMasuk,
		"TotalKeluar": totalKeluar,
		"IsLoggedIn":  isLoggedIn,
	})
}

func ExportPDF(c *gin.Context) {
	bulan := c.Query("bulan")

	if bulan == "" {
		c.String(http.StatusBadRequest, "Bulan harus diisi")
		return
	}

	var keuangan []models.Keuangan
	start, _ := time.Parse("2006-01", bulan)
	end := start.AddDate(0, 1, 0)

	database.DB.Where("tanggal >= ? AND tanggal < ?", start, end).Find(&keuangan)

	var totalMasuk, totalKeluar float64
	for _, k := range keuangan {
		if k.Tipe == "Pemasukan" {
			totalMasuk += k.Jumlah
		} else {
			totalKeluar += k.Jumlah
		}
	}

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	pdf.Ln(30)
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(0, 10, "Laporan Keuangan UKM Badminton")
	pdf.Ln(10)

	pdf.SetFont("Arial", "", 12)
	loc, _ := time.LoadLocation("Asia/Jakarta")
	monthName := start.In(loc).Format("January 2006")
	pdf.Cell(0, 10, "Periode: "+monthName)
	pdf.Ln(10)

	pdf.SetFont("Arial", "B", 11)
	pdf.SetFillColor(230, 230, 230)
	pdf.CellFormat(30, 8, "Tanggal", "1", 0, "C", true, 0, "")
	pdf.CellFormat(80, 8, "Deskripsi", "1", 0, "C", true, 0, "")
	pdf.CellFormat(40, 8, "Tipe", "1", 0, "C", true, 0, "")
	pdf.CellFormat(30, 8, "Jumlah", "1", 1, "C", true, 0, "")

	pdf.SetFont("Arial", "", 11)
	for _, k := range keuangan {
		pdf.CellFormat(30, 8, k.Tanggal.Format("02-01-2006"), "1", 0, "C", false, 0, "")
		pdf.CellFormat(80, 8, k.Deskripsi, "1", 0, "L", false, 0, "")
		pdf.CellFormat(40, 8, k.Tipe, "1", 0, "C", false, 0, "")
		pdf.CellFormat(30, 8, fmt.Sprintf("Rp %.0f", k.Jumlah), "1", 1, "R", false, 0, "")
	}

	pdf.Ln(5)
	pdf.SetFont("Arial", "B", 11)
	pdf.CellFormat(150, 8, "Total Pemasukan", "0", 0, "R", false, 0, "")
	pdf.CellFormat(30, 8, fmt.Sprintf("Rp %.0f", totalMasuk), "0", 1, "R", false, 0, "")

	pdf.CellFormat(150, 8, "Total Pengeluaran", "0", 0, "R", false, 0, "")
	pdf.CellFormat(30, 8, fmt.Sprintf("Rp %.0f", totalKeluar), "0", 1, "R", false, 0, "")

	c.Header("Content-Type", "application/pdf")
	c.Header("Content-Disposition", "attachment; filename=laporan_keuangan_"+bulan+".pdf")
	_ = pdf.Output(c.Writer)
}
