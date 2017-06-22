package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"sort"
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
	top := flag.Int("top", 5, "how many top starred repos will be displayed")
	flag.Parse()

	result := aggregate(getStarredRepositories(getFollowingList(*me)))
	fmt.Printf("Top %d repositories starred by the people you follow: \n", *top)
	for _, res := range result[0:*top] {
		fmt.Printf("Repository %s starred by %d of the people you follow\n", res.RepoName, res.Starred)
	}
}

func getFollowingList(me string) <-chan *user { // returns a list of users I am following
	req, err := http.NewRequest("GET", fmt.Sprintf(fmtFollowingURL, me), nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", *token))
	if err != nil {
		panic(err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	d := json.NewDecoder(resp.Body)

	users := []*user{}

	if err := d.Decode(&users); err != nil {
		panic(err)
	}

	out := make(chan *user, len(users))
	go func() {
		for _, user := range users {
			out <- user
		}
		close(out)
	}()

	return out
}

func getStarredRepositories(in <-chan *user) <-chan repo { // returns a list of repositories being starred by a list of users
	out := make(chan repo)
	go func() {
		wg := sync.WaitGroup{}
		for u := range in {
			wg.Add(1)
			go func(u *user) {
				for repo := range getStarredRepositoriesByUser(u) {
					out <- repo
				}
				wg.Done()
			}(u)
		}
		wg.Wait()
		close(out)
	}()

	return out
}

func getStarredRepositoriesByUser(u *user) <-chan repo {
	req, err := http.NewRequest("GET", fmt.Sprintf(fmtStargazerURL, u.Login), nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", *token))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	d := json.NewDecoder(resp.Body)
	repos := []*repo{}
	if err := d.Decode(&repos); err != nil {
		panic(err)
	}

	out := make(chan repo, len(repos))
	go func() {
		for _, repo := range repos {
			out <- *repo
		}
		close(out)
	}()

	return out
}

func aggregate(repos <-chan repo) []stat {
	result := map[string]int{}
	for repo := range repos {
		result[repo.FullName]++
	}
	stats := make([]stat, 0)
	for repo, stars := range result {
		stats = append(stats, stat{repo, stars})
	}
	sort.Slice(stats, func(i, j int) bool { return stats[i].Starred > stats[j].Starred })
	return stats
}
