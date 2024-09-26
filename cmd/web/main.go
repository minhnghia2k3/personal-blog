package main

import (
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/minhnghia2k3/personal-blog/internal"
	"github.com/minhnghia2k3/personal-blog/internal/config"
	"github.com/minhnghia2k3/personal-blog/internal/database"
	"github.com/minhnghia2k3/personal-blog/internal/handlers"
	"github.com/minhnghia2k3/personal-blog/internal/helpers"
	"github.com/minhnghia2k3/personal-blog/internal/logger"
	"github.com/minhnghia2k3/personal-blog/internal/repositories"
	"github.com/minhnghia2k3/personal-blog/internal/routes"
	"github.com/minhnghia2k3/personal-blog/internal/services"
	"log"
)

func main() {
	var err error

	// Load env
	err = godotenv.Load()
	helpers.MustCatch(err)

	// Initialize logger
	l := logger.New()
	l.DefaultLog()

	// Load app config
	cfg := config.Load()
	app := internal.NewApplication(cfg)

	// Load db connection pool
	db, err := database.ConnectDB()
	helpers.MustCatch(err)
	defer db.Close()

	// Initialize Categories module
	categoryRepo := repositories.NewPostgresCategoryRepository(db)
	categoryService := services.NewCategoryService(categoryRepo)
	categoryHandler := handlers.NewCategoryHandlers(categoryService)

	// Initialize repositories, services, and handlers
	articleRepo := repositories.NewPostgresArticleRepository(db)
	articleService := services.NewArticleService(articleRepo, categoryRepo)
	articleHandler := handlers.NewArticleHandler(articleService, categoryService)

	// Initialize Image handlers
	imageHandler := handlers.NewImageHandler()

	// Initialize router
	router := routes.Routes(articleHandler, categoryHandler, imageHandler)

	// Serve server
	log.Fatal(app.Serve(router))
}
