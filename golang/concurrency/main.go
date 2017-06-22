package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"sync"
)

const (
	fmtFollowingURL = "https://api.github.com/users/%s/following"
	fmtStargazerURL = "https://api.github.com/users/%s/starred"
)

var token *string

type (
	user struct {
		Login string `json:"login"`
		URL   string `json:"url"`
	}
	repo struct {
		FullName string `json:"full_name"`
	}
	stat struct {
		RepoName string
		Starred  int
	}
)

func main() {
	me := flag.String("me", "ideahitme", "my github handle")
	token = flag.String("token", "", "token to authenticate with github API")
	flag.Parse()

	done := make(chan struct{})
	repos := starredByUsers(done, followedBy(done, *me))
	counter := map[string]int{}
	result := map[string]int{}
	for repo := range repos {
		if len(result) == 5 {
			close(done)
			break
		}

		counter[repo.FullName]++
		if counter[repo.FullName] > 1 {
			result[repo.FullName] = counter[repo.FullName]
		}
	}
	for repo, stars := range result {
		if stars > 1 {
			fmt.Printf("repo %s has %d stars\n", repo, stars)
		}
	}
}

func followedBy(done <-chan struct{}, me string) <-chan *user { // returns a list of users I am following
	req, err := http.NewRequest("GET", fmt.Sprintf(fmtFollowingURL, me), nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", *token))
	if err != nil {
		panic(err)
	}
	resp, err := http.DefaultClient.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		panic(err)
	}

	users := []*user{}
	if err := json.NewDecoder(resp.Body).Decode(&users); err != nil {
		panic(err)
	}

	out := make(chan *user)
	go func() {
		defer close(out)
		for _, user := range users {
			select {
			case out <- user:
			case <-done:
				return
			}
		}
	}()

	return out
}

func starredByUsers(done <-chan struct{}, in <-chan *user) <-chan repo { // returns a list of repositories being starred by a list of users
	out := make(chan repo)

	go func() {
		wg := sync.WaitGroup{}
		for u := range in {
			wg.Add(1)
			go func(u *user) {
				defer wg.Done()
				for repo := range starredBy(done, u) {
					select {
					case out <- repo:
					case <-done:
						return
					}
				}
			}(u)
		}
		wg.Wait()
		close(out)
	}()

	return out
}

func starredBy(done <-chan struct{}, u *user) <-chan repo {
	req, err := http.NewRequest("GET", fmt.Sprintf(fmtStargazerURL, u.Login), nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", *token))

	resp, err := http.DefaultClient.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		panic(err)
	}

	repos := []*repo{}
	if err := json.NewDecoder(resp.Body).Decode(&repos); err != nil {
		panic(err)
	}

	out := make(chan repo)
	go func() {
		defer close(out)
		for _, repo := range repos {
			select {
			case out <- *repo:
			case <-done:
				return
			}
		}
	}()

	return out
}
