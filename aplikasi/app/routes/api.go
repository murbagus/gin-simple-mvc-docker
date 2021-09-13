package routes

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/murbagus/gin-simple-mvc/app/appcontext"
	"github.com/murbagus/gin-simple-mvc/app/controllers"

	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/swaggo/gin-swagger/swaggerFiles"

	_ "github.com/murbagus/gin-simple-mvc/docs/api"
)

// @title Test Docs Route API
// @version 1.0
// @description Ini adalah dokumentasi API percobaan.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api
func CreateRouteApi(app *gin.Engine, context *appcontext.Context) {
	routePrefix := "/api"

	// Mengambil route setiap controller
	route := app.Group(routePrefix)

	// Generate dokumentasi api dengan swagger
	docsBaseUrl := fmt.Sprintf("http://%s:%s", os.Getenv("APP_URL"), os.Getenv("APP_PORT"))
	docsUrl := ginSwagger.URL(docsBaseUrl + routePrefix + "/swagger/doc.json")
	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, docsUrl))

	controllers.CreateController(route, context, "/ping", &controllers.PingController{})
	controllers.CreateController(route, context, "/pengguna", &controllers.PenggunaController{})
}
