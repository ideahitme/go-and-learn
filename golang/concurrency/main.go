package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"strings"
	"sync"
)

const (
	fmtFollowingURL = "https://api.github.com/users/%s/following?per_page=5"
	fmtStargazerURL = "https://api.github.com/users/%s/starred?per_page=10"
)

var token *string

type (
	User struct {
		Handle string `json:"login"`
		URL    string `json:"url"`
	}
	Repo struct {
		Name string `json:"full_name"`
		Link string `json:"url"`
	}
)

func main() {
	me := flag.String("me", "ideahitme", "my github handle")
	token = flag.String("token", "", "token to authenticate with github API")
	flag.Parse()

	answer := map[string]uint16{} // repository name -> how many stars among
	done := make(chan struct{})
	repos := reposStarredBy(done, usersFollowedBy(*me)) //
	for repo := range repos {
		answer[repo.Link]++
		if len(answer) > 10 {
			close(done)
			break
		}
	}
	fmt.Println(answer)
}

func usersFollowedBy(me string) <-chan *User {
	users := make(chan *User)
	nextPage := fmt.Sprintf(fmtFollowingURL, me)
	go func() {
		wg := sync.WaitGroup{}
		var (
			result []*User
			err    error
		)
		for nextPage != "" {
			result, nextPage, err = getUsers(nextPage)
			if err != nil {
				panic(err)
			}

			wg.Add(1)
			go func() {
				for _, user := range result {
					users <- user
				}
				wg.Done()
			}()
		}
		wg.Wait()
		close(users)
	}()

	return users
}

func getUsers(url string) ([]*User, string, error) {
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", *token))
	if err != nil {
		return nil, "", err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, "", err
	}
	defer resp.Body.Close()

	result := []*User{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, "", err
	}
	return result, getNextPage(resp.Header.Get("Link")), nil
}

func reposStarredBy(done <-chan struct{}, in <-chan *User) <-chan *Repo { // returns a list of repositories being starred by a list of users
	repos := make(chan *Repo)

	go func() {
		wg := sync.WaitGroup{}
		for u := range in {
			wg.Add(1)
			go func(u *User) {
				defer wg.Done()
				for repo := range starredBy(u) {
					select {
					case repos <- repo:
					case <-done:
					}
				}
			}(u)
		}
		wg.Wait()
		close(repos)
	}()

	return repos
}

func starredBy(u *User) <-chan *Repo {
	out := make(chan *Repo)
	nextPage := fmt.Sprintf(fmtStargazerURL, u.Handle)

	go func() {
		wg := sync.WaitGroup{}
		var (
			repos []*Repo
			err   error
		)
		for nextPage != "" {
			repos, nextPage, err = getRepos(nextPage)

			if err != nil {
				panic(err)
			}

			wg.Add(1)
			go func() {
				for _, repo := range repos {
					out <- repo
				}
				wg.Done()
			}()
		}
		wg.Wait()
		close(out)
	}()

	return out
}

func getRepos(url string) ([]*Repo, string, error) {
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", *token))
	if err != nil {
		return nil, "", err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, "", err
	}
	defer resp.Body.Close()

	result := []*Repo{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, "", err
	}
	return result, getNextPage(resp.Header.Get("Link")), nil
}

func getNextPage(linkHeader string) (nextPage string) {
	if linkHeader == "" {
		return ""
	}
	defer func() {
		if err := recover(); err != nil {
			nextPage = ""
		}
	}()
	left := strings.Split(linkHeader, ",")[0]
	url := strings.Split(left, ";")[0]
	rel := strings.Split(left, ";")[1]
	if strings.TrimSpace(rel) != `rel="next"` {
		return ""
	}
	nextPage = strings.TrimLeft(url, "<")
	nextPage = strings.TrimRight(nextPage, ">")
	return nextPage
}
