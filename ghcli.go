package main

import (
	"encoding/json"
	"net/http"

	"github.com/tsuzu/modi/pkg/ghclitransport"
)

var (
	githubClient = http.Client{Transport: &ghclitransport.Transport{}}
)

func listOrgs() ([]string, error) {
	resp, err := githubClient.Get("/user/orgs")

	if err != nil {
		return nil, err
	}

	type user struct {
		Login string `json:"login"`
	}
	defer resp.Body.Close()

	var users []user
	if err := json.NewDecoder(resp.Body).Decode(&users); err != nil {
		return nil, err
	}

	orgs := make([]string, 0, len(users))
	for _, u := range users {
		orgs = append(orgs, u.Login)
	}

	return orgs, nil
}

func user() (string, error) {
	resp, err := githubClient.Get("/user")

	if err != nil {
		return "", err
	}

	type user struct {
		Login string `json:"login"`
	}
	defer resp.Body.Close()

	var u user
	if err := json.NewDecoder(resp.Body).Decode(&u); err != nil {
		return "", err
	}

	return u.Login, nil
}

func listUserOrgs() ([]string, error) {
	orgs, err := listOrgs()

	if err != nil {
		return nil, err
	}

	u, err := user()

	if err != nil {
		return nil, err
	}

	return append(append([]string{}, u), orgs...), nil
}
