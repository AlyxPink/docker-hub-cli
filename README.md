# Docker Hub CLI

|Purpose        | Badges |
|---------------|--------|
| Documentation |[![Go Reference](https://pkg.go.dev/badge/github.com/victorbersy/docker-hub-cli.svg)](https://pkg.go.dev/github.com/victorbersy/docker-hub-cli)|
| Tests         |[![Go Report Card](https://goreportcard.com/badge/github.com/victorbersy/docker-hub-cli)](https://goreportcard.com/report/github.com/victorbersy/docker-hub-cli) ![build](https://github.com/VictorBersy/docker-hub-cli/actions/workflows/build.yml/badge.svg) [![codecov](https://codecov.io/github/victorbersy/docker-hub-cli/branch/main/graph/badge.svg?token=ODN9OUF0A4)](https://codecov.io/github/victorbersy/docker-hub-cli)|
| Triage        |[![Open Source Helpers](https://www.codetriage.com/victorbersy/docker-hub-cli/badges/users.svg)](https://www.codetriage.com/victorbersy/docker-hub-cli)|

CLI that helps users manage and browse Docker Hub resources (repositories, organizations, members, etc...)

![image](https://user-images.githubusercontent.com/2109178/180596993-6b6638d8-6dfb-4a84-9bc8-172f282e8af3.png)

## Launch

### Go
```console
go install github.com/victorbersy/docker-hub-cli@latest
docker-hub-cli
```
### Docker
```console
# Docker Hub
docker pull victorbersy/docker-hub-cli:latest
docker run --rm -it victorbersy/docker-hub-cli

# Github Container Registry
docker pull ghcr.io/victorbersy/docker-hub-cli:latest
docker run --rm -it ghcr.io/victorbersy/docker-hub-cli
```
## Screenshots

![image](https://user-images.githubusercontent.com/2109178/180597089-22be7878-8a27-4fe6-8401-be28cc26fa0b.png)
![image](https://user-images.githubusercontent.com/2109178/180597084-c1e26447-91ce-4b82-994f-481886776fad.png)
![image](https://user-images.githubusercontent.com/2109178/180597083-66beebdb-3b60-401f-ab78-343b866b3986.png)
![image](https://user-images.githubusercontent.com/2109178/180597082-180ff55a-18e1-4056-83a9-26694c1fcc23.png)

## Credits

- Heavily inspired from [gh-dash](https://github.com/dlvhdr/gh-dash/) created by [@dlvhdr](https://github.com/dlvhdr/)
- [bubbletea by Charm](https://github.com/charmbracelet/bubbletea)
- [glamour by Charm](https://github.com/charmbracelet/glamour)
- [lipgloss by Charm](https://github.com/charmbracelet/lipgloss)
- [Docker Hub](https://hub.docker.com/)
