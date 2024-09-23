package helpers

import (
	"fmt"
	"github.com/minhnghia2k3/personal-blog/internal/models"
	"log"
	"net/http"
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
