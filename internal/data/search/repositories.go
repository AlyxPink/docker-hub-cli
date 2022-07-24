package data_search

import (
	"fmt"
	"log"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/imroc/req/v3"
	"github.com/victorbersy/docker-hub-cli/internal/ui/constants"
)

type RepositoryPage struct {
	PageSize     int              `json:"page_size"`
	PageId       int              `json:"page"`
	NextUrl      string           `json:"next"`
	PreviousUrl  string           `json:"previous"`
	Repositories []RepositoryData `json:"summaries"`
}

type Category struct {
	Name  string `json:"name"`
	Label string `json:"label"`
}

type OperatingSystem struct {
	Name  string `json:"name"`
	Label string `json:"label"`
}

type RepositoryData struct {
	Architectures       []Architecture    `json:"architectures"`
	Categories          []Category        `json:"categories"`
	CertificationStatus string            `json:"certification_status"`
	Created_at          time.Time         `json:"created_at"`
	Description         string            `json:"short_description"`
	FilterType          string            `json:"filter_type"`
	Name                string            `json:"name"`
	OperatingSystems    []OperatingSystem `json:"operating_systems"`
	Publisher           Publisher         `json:"publisher"`
	PullCount           string            `json:"pull_count"`
	Slug                string            `json:"slug"`
	Source              string            `json:"source"`
	StarCount           int               `json:"star_count"`
	Type                string            `json:"type"`
	UpdatedAt           time.Time         `json:"updated_at"`
	Labels              []Label
}

type Label struct {
	Name    string
	Glyph   string
	Color   lipgloss.AdaptiveColor
	Enabled bool
}

type Publisher struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Architecture struct {
	Name  string `json:"name"`
	Label string `json:"label"`
}

func (data RepositoryData) GetUrl() string {
	if data.Publisher.Id == "docker" {
		return fmt.Sprintf("https://hub.docker.com/_/%s", data.Slug)
	}
	return fmt.Sprintf("https://hub.docker.com/r/%s", data.Slug)
}

func (data *RepositoryData) setLabels() {
	data.Labels = append(data.Labels, Label{
		Name:    "Docker Official",
		Glyph:   constants.GlyphLabelDockerOfficial,
		Color:   constants.ColorLabelDockerOfficial,
		Enabled: (data.Source == "store"),
	})
	data.Labels = append(data.Labels, Label{
		Name:    "Verified Publisher",
		Glyph:   constants.GlyphLabelVerifiedPublisher,
		Color:   constants.ColorLabelVerifiedPublisher,
		Enabled: (data.Source == "verified_publisher"),
	})
	data.Labels = append(data.Labels, Label{
		Name:    "Open Source Program",
		Glyph:   constants.GlyphLabelOpenSourceProgram,
		Color:   constants.ColorLabelOpenSourceProgram,
		Enabled: (data.Source == "open_source"),
	})
	data.Labels = append(data.Labels, Label{
		Name:    "Community",
		Glyph:   constants.GlyphLabelCommunity,
		Color:   constants.ColorLabelCommunity,
		Enabled: (data.Source == "community"),
	})
}

func FetchRepositories() ([]RepositoryData, error) {
	client := req.C().
		SetTimeout(5 * time.Second)

	var repositoryPage RepositoryPage
	resp, err := client.R().
		SetHeader("Search-Version", "v3").
		SetResult(&repositoryPage).
		EnableDump().
		SetQueryParam("image_filter", "official,store,open_source").
		SetQueryParam("order", "desc").
		SetQueryParam("page_size", "100").
		SetQueryParam("page", "1").
		SetQueryParam("type", "image").
		Get("https://hub.docker.com/api/content/v1/products/search")
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
