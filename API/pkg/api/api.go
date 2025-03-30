package api

import (
	"hospitalApi/pkg/api/infrastructure"
	"hospitalApi/pkg/web/mid"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type APIConfig struct {
	DB      *gorm.DB
	Timeout time.Duration
}

func APIMux(cfg APIConfig) http.Handler {

	createHelper := infrastructure.CreateHelper()
	createRepository := infrastructure.CreateRepository(cfg.DB, createHelper)
	createService := infrastructure.CreateService(createRepository, createHelper)
	createHandler := infrastructure.CreateHandler(createService, createHelper)

	app := gin.Default()
	createRouting(app, createHandler)

	return app
}

func createRouting(app *gin.Engine, handler infrastructure.Handler) {
	app.Use(mid.Cors("*"))

	staffAPI(app, handler)
	patientAPI(app, handler)
}

func patientAPI(app *gin.Engine, handler infrastructure.Handler) {
	patient := app.Group("/patient")
	{
		patient.POST("/search", mid.AuthMiddleware(), handler.PatientHandler.Search)
		patient.GET("/search/:id", mid.AuthMiddleware(), handler.PatientHandler.SearchById)
	}
}

func staffAPI(app *gin.Engine, handler infrastructure.Handler) {
	staff := app.Group("/staff")
	{
		staff.GET("/login", handler.StaffHandler.Login)
		staff.POST("/create", handler.StaffHandler.Create)
	}
}
