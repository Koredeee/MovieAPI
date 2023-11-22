package routes

import (
	"MovieAPI/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// set db to gin context
	r.Use(func(ctx *gin.Context) {
		ctx.Set("db", db)
	})

	r.GET("/age-rating-categories", controllers.GetAllRating)
	r.POST("/age-rating-categories", controllers.CreateRating)
	r.GET("/age-rating-categories/:id", controllers.GetRatingById)
	r.GET("/age-rating-categories/:id/movies", controllers.GetMoviesByAgeRatingCategoryId)
	r.PATCH("/age-rating-categories/:id", controllers.UpdateRating)
	r.DELETE("/age-rating-categories/:id", controllers.DeleteRating)

	r.GET("/movies", controllers.GetAllMovie)
	r.POST("/movies", controllers.CreateMovie)
	r.GET("/movies/:id", controllers.GetMovieById)
	r.PATCH("/movies/:id", controllers.UpdateMovie)
	r.DELETE("/movies/:id", controllers.DeleteMovie)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
