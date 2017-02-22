/*
 * Exercise 4.14: Create a web server that queries GitHub once and then allows navigation of the list of bug reports,
 * milestones, and users.
 */

// TODO
package main

import (
	"net/http"
	"sync"
	"log"
	"fmt"
	"encoding/json"
)



var mu sync.Mutex
var count int

var GITHUB_URL string = "https://api.github.com"

type UsersResult struct {
	Users 	[]*User
}

type User struct {
	Id	int
	Login	string
}

//
//{
//"login": "bmizerany",
//"id": 46,
//"avatar_url": "https://avatars.githubusercontent.com/u/46?v=3",
//"gravatar_id": "",
//"url": "https://api.github.com/users/bmizerany",
//"html_url": "https://github.com/bmizerany",
//"followers_url": "https://api.github.com/users/bmizerany/followers",
//"following_url": "https://api.github.com/users/bmizerany/following{/other_user}",
//"gists_url": "https://api.github.com/users/bmizerany/gists{/gist_id}",
//"starred_url": "https://api.github.com/users/bmizerany/starred{/owner}{/repo}",
//"subscriptions_url": "https://api.github.com/users/bmizerany/subscriptions",
//"organizations_url": "https://api.github.com/users/bmizerany/orgs",
//"repos_url": "https://api.github.com/users/bmizerany/repos",
//"events_url": "https://api.github.com/users/bmizerany/events{/privacy}",
//"received_events_url": "https://api.github.com/users/bmizerany/received_events",
//"type": "User",
//"site_admin": false
//}
//]

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	res, err := getUsers()
	log.Printf("resp: %v", res)
	log.Printf("err: %v", err)
	fmt.Fprintf(w, "URL.Path = %v\n", res)
}

func counter(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Count %d\n", count)
}

func getUsers() (*[]User, error){
	//q := url.QueryEscape("/users")
	resp, err := http.Get(GITHUB_URL + "/users")
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("query failed: %s", resp.Status)
	}

	var users []User

	log.Printf("Body: %s", resp.Body)

	if err := json.NewDecoder(resp.Body).Decode(&users); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &users, nil
}
