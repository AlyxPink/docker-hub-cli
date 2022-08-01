package data_user

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/imroc/req/v3"
)

type OrganizationPage struct {
	Count         int            `json:"count"`
	NextUrl       string         `json:"next"`
	PreviousUrl   string         `json:"previous"`
	Organizations []Organization `json:"results"`
}

type Organization struct {
	Id            string    `json:"id"`
	Orgname       string    `json:"orgname"`
	Name          string    `json:"full_name"`
	Location      string    `json:"location"`
	Company       string    `json:"company"`
	ProfileUrl    string    `json:"profile_url"`
	DateJoined    time.Time `json:"date_joined"`
	GravatarUrl   string    `json:"gravatar_url"`
	GravatarEmail string    `json:"gravatar_email"`
	Type          string    `json:"type"`
	Badge         string    `json:"badge"`
}

func (data Organization) GetUrl() string {
	return fmt.Sprintf("https://hub.docker.com/u/%s", data.Orgname)
}

func FetchOrganizations() ([]Organization, error) {
	client := req.C().
		SetTimeout(5 * time.Second)

	var organizationPage OrganizationPage
	if os.Getenv("DOCKER_USERNAME") == "" {
		return nil, nil
	}

	orgs_url := "https://hub.docker.com/v2/user/orgs/"
	resp, err := client.R().
		SetHeader("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("DOCKER_BEARER"))).
		SetResult(&organizationPage).
		EnableDump().
		SetQueryParam("page_size", "100").
		SetQueryParam("page", "1").
		Get(orgs_url)
	if err != nil {
		log.Println("error:", err)
		log.Println("raw content:")
		log.Println(resp.Dump())
		return nil, err
	}

	if resp.IsSuccess() {
		return organizationPage.Organizations, err
	}

	log.Println("unknown status", resp.Status)
	log.Println("raw content:")
	log.Println(resp.Dump()) // Record raw content when server returned unknown status code.

	return organizationPage.Organizations, err
}
