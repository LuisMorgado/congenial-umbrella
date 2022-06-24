package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

type pagination struct {
	currentPage int
	totalPages  int
	boundaries  int
	around      int
	perPage     int
}

type Item struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Metadata struct {
	Page    int `json:"page"`
	PerPage int `json:"per_page"`
	Total   int `json:"total"`
}

type Result struct {
	Metadata   Metadata `json:"metadata"`
	Pagination string   `json:"pagination"`
	Items      []Item   `json:"items"`
}

//this is the method who runs all the logic.
//get the query params and populate the methods that we need and where is created the object to return
func viewHandler(w http.ResponseWriter, r *http.Request) {
	result := Result{}
	queryParams := r.URL.Query()
	myPagination, err := validateInput(queryParams)

	if err != nil {
		errorValidator(w, result)
		return
	}

	result.Metadata.Page = myPagination.currentPage
	result.Metadata.PerPage = myPagination.perPage
	result.Metadata.Total = myPagination.totalPages

	if err != nil {
		errorValidator(w, result)
		return
	}

	count := myPagination.currentPage
	for i := 0; i < myPagination.perPage; i++ {
		result.Items, count = fizzBuzz(result.Items, count)
	}

	result.Pagination = getPagination(myPagination.boundaries, myPagination.currentPage, myPagination.totalPages, myPagination.around)

	ok(w, result)
}

//This method validates all the query params and the currect inputs
func validateInput(queryParams url.Values) (pagination, error) {
	boundaries, err := strconv.Atoi(queryParams.Get("boundaries"))
	if err != nil {
		return pagination{}, err
	}

	currentPage, err := strconv.Atoi(queryParams.Get("current_page"))
	if err != nil {
		return pagination{}, err
	}

	totalPages, err := strconv.Atoi(queryParams.Get("total_pages"))
	if err != nil {
		return pagination{}, err
	}

	around, err := strconv.Atoi(queryParams.Get("around"))
	if err != nil {
		return pagination{}, err
	}

	perPage, err := strconv.Atoi(queryParams.Get("per_page"))
	if err != nil {
		return pagination{}, err
	}

	if boundaries < 0 || currentPage < 0 || totalPages < 0 || around < 0 {
		return pagination{}, fmt.Errorf("invalid input")
	}

	if currentPage > totalPages || boundaries > totalPages || around > totalPages {
		return pagination{}, fmt.Errorf("invalid input")
	}

	return pagination{
		boundaries:  boundaries,
		currentPage: currentPage,
		totalPages:  totalPages,
		around:      around,
		perPage:     perPage,
	}, nil
}

//Returns a ok: statusCode 200
func ok(w http.ResponseWriter, result Result) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(result)
}

//Returns a badRequest error: statusCode 400
func errorValidator(w http.ResponseWriter, result Result) {
	w.WriteHeader(http.StatusBadRequest)

	json.NewEncoder(w).Encode(result)
}
