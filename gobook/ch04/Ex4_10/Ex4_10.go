package main

import (
	"encoding/json"
	"net/url"
	"strings"
	"net/http"
	"fmt"
	"os"
	"log"
	"time"
)

/*
 * Exercise 4.10: Modify issues to report the results in age categories, say less than a month old,
 less than a year old, and more than a year old.
 */

const IssuesURL = "https://api.github.com/search/issues"

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string    // in Markdown format
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}


func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result IssuesSearchResult

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}

func main() {
	result, err := SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	groups := make(map[string]int)
	for _, item := range result.Items {
		if time.Now().Sub(item.CreatedAt).Hours() / 24 < 30 {
			groups["1 month ago"]++
		} else if time.Now().Sub(item.CreatedAt).Hours() / 24 < 365 {
			groups["1 - 12 months ago"]++
		} else {
			groups["1 year or more ago"]++
		}
	}

	for s,i := range groups {
		fmt.Printf("%s : %d\n", s, i)
	}
}
