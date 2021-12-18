package bprepo2json

import (
	"io/ioutil"
	"strings"

	log "github.com/sirupsen/logrus"
)

// Repository is a struct
type Repository struct {
	Name     string    `json:"name"`
	Packages []Package `json:"packages"`
}

// Package is a struct
type Package struct {
	Name    string      `json:"name"`
	Modules []Module    `json:"modules"`
	Repo    *Repository `json:"-"`
}

// Module is a struct
type Module struct {
	Name string   `json:"name"`
	Type string   `json:"type"`
	Pkg  *Package `json:"-"`
}

// NewRepositories is to new global repositories
func NewRepositories(repoListFile string) map[string]*Repository {
	var content string
	buf, err := ioutil.ReadFile(repoListFile)
	if err != nil {
		log.Fatalln("NewRepositories error:", err.Error())
	}
	content = string(buf)

	repoList := strings.Split(content, "\n")
	repos := map[string]*Repository{}
	for _, repoName := range repoList {
		if len(repoName) <= 0 {
			continue
		}
		repo := Repository{}
		repo.Name = repoName
		repo.Packages = []Package{}
		repos[repo.Name] = &repo
	}
	return repos
}
