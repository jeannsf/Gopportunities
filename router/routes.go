package router

import ( 
	"github.com/gin-gonic/gin"
	"github.com/jeannsf/Gopportunities/handler"
	"github.com/jeannsf/Gopportunities/docs" 
	"github.com/swaggo/gin-swagger" 
 	"github.com/swaggo/files" 
)

func initializeRoutes(router *gin.Engine) {

	// Initialize Handler
	handler.InitializeHandler()
	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath
	v1 := router.Group(basePath)

	{
		v1.GET("/opening", handler.ShowOpeningHandler)

		v1.POST("/opening", handler.CreateOpeningHandler)

		v1.DELETE("/opening", handler.DeleteOpeningHandler)

		v1.PUT("/opening", handler.UpdateOpeningHandler)

		v1.GET("/openings", handler.ListOpeningsHandler)
	}

	// Initialize Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

}