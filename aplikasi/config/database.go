package config

import (
	"fmt"
	"io/ioutil"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DatabaseConnectionSettings berisa konfigurasi koneksi database
// untuk koneksi database default gunakan ConnectionName "default"
type DatabaseConnectionSettings struct {
	ConnectionName string
	Driver         string
	User           string
	Password       string
	Domain         string
	Port           string
	DatabaseName   string
}

func GetDatabaseConnection() map[string]*gorm.DB {
	// Setting koneksi database
	// tambahkan konkesi database dalam array connectionSettings
	var connectionSettings = []DatabaseConnectionSettings{
		{
			ConnectionName: "default",
			Driver:         os.Getenv("DB_DRIVER"),
			User:           os.Getenv("DB_USER"),
			Password:       readPasswordFromFile(os.Getenv("DB_PASSWORD")),
			Domain:         os.Getenv("DB_DOMAIN"),
			Port:           os.Getenv("DB_PORT"),
			DatabaseName:   os.Getenv("DB_NAME"),
		},
	}

	return buildConnection(connectionSettings)
}

// readPasswordFromFile digunakan untuk membaca password dari file *.txt
// password yang diletakan dalam file bianya digunakan pada docker secrets
func readPasswordFromFile(file string) string {
	password, err := ioutil.ReadFile(file)
	if err != nil {
		panic("Terjadi kesalahah saat membaca file password database")
	}

	return string(password)
}

// buildConnection dipanggil dalam fungsi GetDatabaseConnection
// untuk membuat koneksi-koneksi database yang di definisikan di sana
// database dibuat berdasarkan setting driver yang diberikan
func buildConnection(connectionSettings []DatabaseConnectionSettings) map[string]*gorm.DB {
	connections := map[string]*gorm.DB{}

	for _, connectionSetting := range connectionSettings {
		// Periksa driver yang digunakan
		if connectionSetting.Driver == "mysql" {
			// Koneksi driver MySQL
			dsn := fmt.Sprintf(
				"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
				connectionSetting.User,
				connectionSetting.Password,
				connectionSetting.Domain,
				connectionSetting.Port,
				connectionSetting.DatabaseName,
			)
			if db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}); err != nil {
				panic("Terjadi kesalahan saat membuat koneksi database")
			} else {
				connections[connectionSetting.ConnectionName] = db
			}
		}
	}

	return connections
}
