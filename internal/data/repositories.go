package data

import (
	"fmt"
	"log"
	"time"

	"github.com/imroc/req/v3"
)

type RepositoryPage struct {
	PageSize     int              `json:"page_size"`
	PageId       int              `json:"page"`
	NextUrl      string           `json:"next"`
	PreviousUrl  string           `json:"previous"`
	Repositories []RepositoryData `json:"summaries"`
}

type RepositoryData struct {
	Labels              labels
}

type labels struct {
	DockerOfficial    bool
	VerifiedPublisher bool
	OpenSourceProgram bool
	Community         bool
}

type Publisher struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Architecture struct {
	Name  string `json:"name"`
	Label string `json:"label"`
}

func (data RepositoryData) GetRepoNameWithOwner() string {
	return data.Name
}

func (data RepositoryData) GetUrl() string {
	if data.Publisher.Id == "docker" {
		return fmt.Sprintf("https://hub.docker.com/_/%s", data.Slug)
	}
	return fmt.Sprintf("https://hub.docker.com/r/%s", data.Slug)
}

func (data RepositoryData) GetLastUpdate() time.Time {
	return data.Updated_at
}

func (data *RepositoryData) setLabels() {
	data.Labels = labels{
		DockerOfficial:    (data.Source == "store"),
		VerifiedPublisher: (data.Source == "verified_publisher"),
		OpenSourceProgram: (data.Source == "open_source"),
		Community:         (data.Source == "community"),
	}
}

func FetchRepositories() ([]RepositoryData, error) {
	client := req.C().
		SetTimeout(5 * time.Second)

	var repositoryPage RepositoryPage
	resp, err := client.R().
		SetHeader("Search-Version", "v3").
		SetResult(&repositoryPage).
		EnableDump().
		Get("https://hub.docker.com/api/content/v1/products/search?page_size=50")
	if err != nil {
		log.Println("error:", err)
		log.Println("raw content:")
		log.Println(resp.Dump())
		return nil, err
	}

	if resp.IsSuccess() {
		for i := range repositoryPage.Repositories {
			repositoryPage.Repositories[i].setLabels()
		}
		return repositoryPage.Repositories, err
	}

	log.Println("unknown status", resp.Status)
	log.Println("raw content:")
	log.Println(resp.Dump()) // Record raw content when server returned unknown status code.

	return repositoryPage.Repositories, err
}
