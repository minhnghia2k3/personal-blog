package helpers

import (
	"fmt"
	"github.com/minhnghia2k3/personal-blog/internal/dto"
	"github.com/minhnghia2k3/personal-blog/internal/models"
	"log"
	"net/http"
	"strconv"
)

// Catch catches error occurred
func Catch(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

func HttpCatch(w http.ResponseWriter, err error) {
	if err != nil {
		log.Println("HTTP Error:", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func ContainsCategory(categories []*models.Category, category *models.Category) bool {
	for _, c := range categories {
		if c.ID == category.ID { // Adjust this comparison based on your struct's fields
			return true
		}
	}
	return false
}

// GetPaginationValues helpers will parse URL query and return dto.Pagination data.
func GetPaginationValues(r *http.Request) dto.Pagination {
	q := r.URL.Query()

	// Default limit and page values
	defaultLimit := 10
	defaultPage := 1

	// Parse limit and page query parameters
	limitStr := q.Get("limit")
	limit := defaultLimit
	if limitStr != "" {
		limit, _ = strconv.Atoi(limitStr)
	}

	pageStr := q.Get("page")
	page := defaultPage
	if pageStr != "" {
		page, _ = strconv.Atoi(pageStr)
	}

	search := q.Get("search")

	return dto.Pagination{
		Page:   page,
		Limit:  limit,
		Search: search,
	}

}
