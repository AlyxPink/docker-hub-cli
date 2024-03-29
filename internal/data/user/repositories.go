package data_user

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/imroc/req/v3"
)

type RepositoryPage struct {
	Count        int          `json:"count"`
	NextUrl      string       `json:"next"`
	PreviousUrl  string       `json:"previous"`
	Repositories []Repository `json:"results"`
}

type Repository struct {
	Name           string    `json:"name"`
	Namespace      string    `json:"namespace"`
	RepositoryType string    `json:"repository_type"`
	Status         int       `json:"status"`
	IsPrivate      bool      `json:"is_private"`
	StarCount      int       `json:"star_count"`
	PullCount      int       `json:"pull_count"`
	UpdatedAt      time.Time `json:"last_updated"`
	CreatedAt      time.Time `json:"date_registered"`
	Affiliation    string    `json:"affiliation"`
}

func (data Repository) GetUrl() string {
	return fmt.Sprintf("https://hub.docker.com/repository/docker/%s/%s", data.Namespace, data.Name)
}

func FetchRepositories() ([]Repository, error) {
	client := req.C().
		SetTimeout(5 * time.Second)

	var repositoryPage RepositoryPage
	if os.Getenv("DOCKER_USERNAME") == "" {
		return nil, nil
	}
	repo_url := fmt.Sprintf("https://hub.docker.com/v2/repositories/%s", os.Getenv("DOCKER_USERNAME"))
	resp, err := client.R().
		SetHeader("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("DOCKER_BEARER"))).
		SetResult(&repositoryPage).
		EnableDump().
		SetQueryParam("page_size", "100").
		SetQueryParam("page", "1").
		SetQueryParam("ordering", "last_updated").
		Get(repo_url)
	if err != nil {
		log.Println("error:", err)
		log.Println("raw content:")
		log.Println(resp.Dump())
		return nil, err
	}

	if resp.IsSuccess() {
		return repositoryPage.Repositories, err
	}

	log.Println("unknown status", resp.Status)
	log.Println("raw content:")
	log.Println(resp.Dump()) // Record raw content when server returned unknown status code.

	return repositoryPage.Repositories, err
}
