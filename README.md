# Docker Hub CLI

![build](https://github.com/VictorBersy/docker-hub-cli/actions/workflows/build.yml/badge.svg)

CLI that helps users manage and browse Docker Hub resources (repositories, organizations, members, etc...)

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

![image](https://user-images.githubusercontent.com/2109178/177809263-02338c63-b354-4d4e-9db1-8eb122449346.png)

![image](https://user-images.githubusercontent.com/2109178/177809251-48e3fcf5-8825-4963-8d8a-7924c91b3eb5.png)

## Credits

- Heavily inspired by [gh-dash](https://github.com/dlvhdr/gh-dash/) from @dlvhdr
- [bubbletea by Charm](https://github.com/charmbracelet/bubbletea)
- [glamour by Charm](https://github.com/charmbracelet/glamour)
- [lipgloss by Charm](https://github.com/charmbracelet/lipgloss)
