package main

// import (
// 	"net/http"
// 	"fmt"
// )

import (
	"context"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)
// const GITHUB_API_HOST = "https://api.github.com"


// func getIssue(repo string, issueId int, token string) (string, error) {
// 	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s/issues/%d", GITHUB_API_HOST, repo, issueId), nil)
// 	if err != nil {
// 		return nil, err
// 	}

// 	req.Header.Add("Authorization", fmt.Sprintf("token %s", token))

// 	client := &http.Client{}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return req, err
// }

// func getTargets(repo string, issueId int, token string) {
	
// }

func getIssue(user string, repo string, issueId int, token string) (*string, error) {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)
	issue, _, err := client.Issues.Get(ctx, user, repo, issueId)
	return issue.Body, err
}