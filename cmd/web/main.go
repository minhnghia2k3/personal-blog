package main

import (
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/minhnghia2k3/personal-blog/internal/database"
	"github.com/minhnghia2k3/personal-blog/internal/handlers"
	"github.com/minhnghia2k3/personal-blog/internal/helpers"
	"github.com/minhnghia2k3/personal-blog/internal/repositories"
	"github.com/minhnghia2k3/personal-blog/internal/routes"
	"github.com/minhnghia2k3/personal-blog/internal/services"
	"net/http"
)

func main() {
	var err error

	// Load env
	err = godotenv.Load()
	helpers.Catch(err)

	db, err := database.ConnectDB()
	helpers.Catch(err)
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

	fmt.Println("Server running on port :8080")
	err = http.ListenAndServe(":8080", router)
	helpers.Catch(err)
}
