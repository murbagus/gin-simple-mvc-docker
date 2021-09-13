package main

import (
	"fmt"
	"os"

	"github.com/murbagus/gin-simple-mvc/app"
)

// Fungsi utama aplikasi
func main() {
	app := app.NewApp()

	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	app.Run(fmt.Sprintf(":%s", os.Getenv("APP_PORT")))
}
