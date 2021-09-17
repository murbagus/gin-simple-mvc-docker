package routes

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/swaggo/gin-swagger/swaggerFiles"

	_ "github.com/murbagus/gin-simple-mvc/docs/api"
)

// SetUpSwagger akan membuatkan url dokumentasi API
// untuk kelompok route tertentu
func SetUpSwagger(route *gin.RouterGroup, routePrefix string) {
	docsBaseUrl := fmt.Sprintf("http://%s:%s", os.Getenv("APP_URL"), os.Getenv("APP_PORT"))
	docsUrl := ginSwagger.URL(docsBaseUrl + routePrefix + "/swagger/doc.json")
	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, docsUrl))
}
