// app.go
package main

import (
	"database/sql"
	"net/http"

	"log"

	"github.com/SulaimonYNWA/GoTemple/handlers"
	"github.com/SulaimonYNWA/GoTemple/repo"
	"github.com/SulaimonYNWA/GoTemple/services"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

func SetupRouter(db *sql.DB) http.Handler {
	router := httprouter.New()

	// Logger
	logger := log.Default()

	// Dependencies
	schoolRepo := repo.NewSchoolRepo(db)
	courseRepo := repo.NewCourseRepo(db)
	instructorRepo := repo.NewInstructorRepo(db)
	schoolService := services.NewSchoolService(schoolRepo, courseRepo, instructorRepo)

	// Handlers
	schoolHandler := &handlers.SchoolHandler{
		Logger:  logger,
		Service: schoolService,
	}

	instructorsHandler := &handlers.InstructorsHandler{
		Logger:  logger,
		Service: schoolService,
	}

	dashboardHandler := &handlers.DashboardHandler{
		Logger:     logger,
		SchoolRepo: schoolRepo,
	}

	// Dashboard routes
	router.GET("/", dashboardHandler.HandleDashboard)
	router.GET("/dashboard", dashboardHandler.HandleDashboard)

	// Instructors
	router.GET("/instructors", instructorsHandler.HandleInstructorsPage)

	// School routes
	router.GET("/schools", schoolHandler.HandleList)         // HTML view
	router.GET("/api/schools", schoolHandler.HandleListJSON) // JSON API
	router.GET("/api/schools/:id/courses", schoolHandler.HandleCoursesBySchool)

	return router
}
