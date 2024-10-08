package helpers

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/minhnghia2k3/personal-blog/internal/dto"
	"github.com/minhnghia2k3/personal-blog/internal/models"
	"log/slog"
	"net/http"
	"strconv"
)

// MustCatch catches and panic exception
func MustCatch(err error) {
	if err != nil {
		slog.Error("Internal Server Error", "status", http.StatusInternalServerError, "error", err)
		panic(err)
	}
}

func HttpCatch(w http.ResponseWriter, status int, err error) {
	if err != nil {
		slog.Debug("HTTP Error:", "error", err)
		http.Error(w, http.StatusText(status), status)
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

type validationError struct {
	Field       string      `json:"field"`
	Tag         string      `json:"tag"`
	ActualValue interface{} `json:"actual_value"`
}

func (v validationError) Error() string {
	return fmt.Sprintf("Field '%s' failed on '%s' validation, got '%v'", v.Field, v.Tag, v.ActualValue)
}

func ValidateStruct(obj interface{}) (errs []error) {
	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(obj)
	if err != nil {
		for _, vErr := range err.(validator.ValidationErrors) {
			e := validationError{
				Field:       vErr.StructField(),
				Tag:         vErr.Tag(),
				ActualValue: vErr.Value(),
			}

			errs = append(errs, e)
		}
		return errs
	}

	return nil
}

// ResponseErrors utility to return validation errors as JSON
func ResponseErrors(w http.ResponseWriter, errs []error) {
	w.WriteHeader(http.StatusBadRequest)
	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"errors": errs,
	})
}

// FormIntValue parse form value to integer type.
func FormIntValue(r *http.Request, name string) (int, error) {
	val, err := strconv.Atoi(r.FormValue("min_read"))
	if err != nil {
		return -1, err
	}
	return val, nil
}

type Response struct {
	StatusCode int
	Msg        string
}

func Respond(w http.ResponseWriter, r Response) {
	_, err := json.Marshal(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
		return
	}

	w.WriteHeader(r.StatusCode)
	fmt.Fprintf(w, r.Msg)
}
