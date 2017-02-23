/*
 * Exercise 4.14: Create a web server that queries GitHub once and then allows navigation of the list of bug reports,
 * milestones, and users.
 */

// TODO
package main

import (
	"net/http"
	"log"
	"fmt"
	"encoding/json"
	"net/url"
)

var GITHUB_URL string = "https://api.github.com"

type UsersResult struct {
	Users 	[]*User
}

type User struct {
	Id	int
	Login	string
}

type RepositoriesResult struct {
	Repositories	[]*Repositories
}
type Repositories struct {
	id 	int
	name 	string
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/repositories", repositories)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	res, _ := getUsers()

	html := "<html>"
	html += "<head>"
	html += "</head>"
	html += "<body>"
	html += "<h1>Users</h1>"
	for i := range *res {
		html += "<p>"
		html += "Repositories for: "
		html += "<a href=\"/repositories?u=" + (*res)[i].Login + "\">"
		html += (*res)[i].Login + "</a>"
		html += "</p>"
	}
	html += "</body>"
	html += "</html>"
	fmt.Fprintf(w, "%s\n", html)
}

func repositories(w http.ResponseWriter, r *http.Request) {
	user := r.URL.Query()["u"]
	if user != nil {
		res, _ := getRepositories(user[0])
		log.Printf("%v", res)

		html := "<html>"
		html += "<head>"
		html += "</head>"
		html += "<body>"
		html += "<h1>Users</h1>"
		//for i := range *res {
		//	html += "<p>"
		//	html += "User: " + user[0] + " Repository: " + (*res)[i].name
		//	html += "</p>"
		//}
		html += "</body>"
		html += "</html>"
		fmt.Fprintf(w, "%s\n", html)
	} else {
		html := "<html>"
		html += "<head>"
		html += "</head>"
		html += "<body>"
		html += "No such user."
		html += "</body>"
		html += "</html>"
		fmt.Fprintf(w, "%s\n", html)
	}
}

func getRepositories(user string)(*RepositoriesResult, error) {
	q := url.QueryEscape(user)
	q = user
	url := GITHUB_URL + "/users/" + q + "/repos"
	log.Printf("%s", url)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Error: %v", err)
		return nil, fmt.Errorf("query failed: %s", resp.Status)
	}

	var repos RepositoriesResult

	log.Printf("Body: %v", resp.Body)

	if err := json.NewDecoder(resp.Body).Decode(&repos); err != nil {
		log.Printf("Error: %v", err)
		return nil, err
	}
	log.Printf("Repos 0: %s", repos.Repositories[0].name)
	log.Printf("Repos 0: %d", repos.Repositories[0].id)

	return &repos, nil
}

func getUsers() (*[]User, error){
	resp, err := http.Get(GITHUB_URL + "/users")
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Error: %v", err)
		return nil, fmt.Errorf("query failed: %s", resp.Status)
	}

	var users []User

	if err := json.NewDecoder(resp.Body).Decode(&users); err != nil {
		log.Printf("Error: %v", err)
		return nil, err
	}

	return &users, nil
}
