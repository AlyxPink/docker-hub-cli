package data

import (
	"time"
)

type RepositoryData struct {
	Name        string
	Url         string
	Labels      Labels
	Publisher   Publisher
	Stats       Stats
	Description string
	LastUpdate  time.Time
}

type Labels struct {
	DockerOfficial    bool
	VerifiedPublisher bool
	OpenSourceProgram bool
}

type Publisher struct {
	Name string
}

type Stats struct {
	Downloads string
	Stars     string
}

func (data RepositoryData) GetRepoNameWithOwner() string {
	return data.Name
}

func (data RepositoryData) GetUrl() string {
	return data.Url
}

func (data RepositoryData) GetLastUpdate() time.Time {
	return data.LastUpdate
}

func fakeRepositories() []RepositoryData {
	repositories := make([]RepositoryData, 0)

	repositories = append(repositories, RepositoryData{
		Name:        "alpine",
		Url:         "https://hub.docker.com/_/alpine",
		Labels:      Labels{DockerOfficial: true, VerifiedPublisher: false, OpenSourceProgram: true},
		Publisher:   Publisher{Name: "Docker"},
		Stats:       Stats{Downloads: "1B+", Stars: "9.0K"},
		Description: "A minimal Docker image based on Alpine Linux with a complete package index and only 5 MB in size!",
		LastUpdate:  time.Now().Add(-(30 * 24 * time.Hour)),
	})

	repositories = append(repositories, RepositoryData{
		Name:        "busybox",
		Url:         "https://hub.docker.com/_/busybox",
		Labels:      Labels{DockerOfficial: true, VerifiedPublisher: false, OpenSourceProgram: false},
		Publisher:   Publisher{Name: "Docker"},
		Stats:       Stats{Downloads: "1B+", Stars: "2.7K"},
		Description: "Busybox base image.",
		LastUpdate:  time.Now().Add(-(30 * 24 * time.Hour)),
	})

	repositories = append(repositories, RepositoryData{
		Name:        "grafana",
		Url:         "https://hub.docker.com/r/grafana/grafana",
		Labels:      Labels{DockerOfficial: false, VerifiedPublisher: true, OpenSourceProgram: false},
		Publisher:   Publisher{Name: "Grafana Labs"},
		Stats:       Stats{Downloads: "1B+", Stars: "2.3K"},
		Description: "The official Grafana docker container ",
		LastUpdate:  time.Now().Add(-(24 * time.Minute)),
	})

	repositories = append(repositories, RepositoryData{
		Name:        "ubuntu",
		Url:         "https://hub.docker.com/_/ubuntu",
		Labels:      Labels{DockerOfficial: true, VerifiedPublisher: false, OpenSourceProgram: false},
		Publisher:   Publisher{Name: "Docker"},
		Stats:       Stats{Downloads: "1B+", Stars: "10K+"},
		Description: "Ubuntu is a Debian-based Linux operating system based on free software.",
		LastUpdate:  time.Now().Add(-(30 * 24 * time.Hour)),
	})

	repositories = append(repositories, RepositoryData{
		Name:        "fluent-bit",
		Url:         "https://hub.docker.com/r/fluent/fluent-bit",
		Labels:      Labels{DockerOfficial: true, VerifiedPublisher: false, OpenSourceProgram: false},
		Publisher:   Publisher{Name: "Fluent organization: Fluentd project"},
		Stats:       Stats{Downloads: "1B+", Stars: "78"},
		Description: "Fluent Bit, lightweight logs and metrics collector and forwarder ",
		LastUpdate:  time.Now().Add(-(16 * 24 * time.Hour)),
	})

	repositories = append(repositories, RepositoryData{
		Name:        "nginx",
		Url:         "https://hub.docker.com/_/nginx",
		Labels:      Labels{DockerOfficial: true, VerifiedPublisher: false, OpenSourceProgram: false},
		Publisher:   Publisher{Name: "Docker"},
		Stats:       Stats{Downloads: "1B+", Stars: "10K+"},
		Description: "Official build of Nginx.",
		LastUpdate:  time.Now().Add(-(13 * 24 * time.Hour)),
	})

	return repositories
}

func FetchRepositories(query string) ([]RepositoryData, error) {
	return fakeRepositories(), nil
}
