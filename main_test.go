package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestGetPagination(t *testing.T) {
	testCases := []struct {
		boundaries  int
		currentPage int
		totalPages  int
		around      int
		expected    string
	}{
		{
			boundaries:  1,
			currentPage: 4,
			totalPages:  5,
			around:      0,
			expected:    "1 ... 4 5 ",
		},
		{
			boundaries:  2,
			currentPage: 4,
			totalPages:  10,
			around:      2,
			expected:    "1 2 3 4 5 6 ... 9 10 ",
		},
	}

	for _, test := range testCases {
		result := getPagination(test.boundaries, test.currentPage, test.totalPages, test.around)
		if result != test.expected {
			t.Error("expected something else")
		}
	}
}

func TestViewHandler(t *testing.T) {
	testCases := []struct {
		boundaries  int
		currentPage int
		totalPages  int
		around      int
		perPage     int
		expected    string
		status      int
	}{
		{
			boundaries:  1,
			currentPage: 4,
			totalPages:  5,
			around:      0,
			perPage:     1,
			expected:    "1 ... 4 5 ",
			status:      200,
		},
		{
			boundaries:  2,
			currentPage: 4,
			totalPages:  10,
			around:      2,
			perPage:     5,
			expected:    "1 2 3 4 5 6 ... 9 10 ",
			status:      200,
		},
		{
			boundaries:  -1,
			currentPage: 4,
			totalPages:  5,
			around:      0,
			perPage:     1,
			expected:    "1 ... 4 5 ",
			status:      400,
		},
		{
			boundaries:  6,
			currentPage: 4,
			totalPages:  5,
			around:      0,
			perPage:     1,
			expected:    "1 ... 4 5 ",
			status:      400,
		},
		{
			boundaries:  1,
			currentPage: 4,
			totalPages:  5,
			around:      10,
			perPage:     1,
			expected:    "1 ... 4 5 ",
			status:      400,
		},
	}

	for _, test := range testCases {
		bdr := strconv.Itoa(test.boundaries)
		tpages := strconv.Itoa(test.totalPages)
		arnd := strconv.Itoa(test.around)
		cpages := strconv.Itoa(test.currentPage)
		pp := strconv.Itoa(test.perPage)

		req, err := http.NewRequest("GET", "localhost:8080?boundaries="+bdr+"&total_pages="+tpages+"&around="+arnd+"&current_page="+cpages+"&per_page="+pp, nil)

		if err != nil {
			t.Error("bad request")
			return
		}

		rec := httptest.NewRecorder()
		viewHandler(rec, req)

		res := rec.Result()
		data, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Error("not found")
			return
		}

		if res.StatusCode != test.status {
			t.Error("something went wrong")
			return
		}

		if res.StatusCode == 400 {
			return
		}

		fmt.Println(string(data))

		var result Result
		err = json.Unmarshal(data, &result)

		if err != nil {
			t.Error("error when trying to deserialize object")
			return
		}

		if result.Pagination != test.expected {
			t.Error("wrong pagination")
			return
		}

		if len(result.Items) != test.perPage {
			t.Error("something went wrong")
			return
		}

		if test.status == 400 {
			t.Error("Erro Case")
		}

	}
}
