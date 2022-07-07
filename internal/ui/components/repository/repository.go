package repository

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/docker/hack-docker-access-management-cli/internal/data"
	"github.com/docker/hack-docker-access-management-cli/internal/ui/components/table"
	"github.com/docker/hack-docker-access-management-cli/internal/utils"
)

type Repository struct {
	Data data.RepositoryData
}

type sectionRepositoriesFetchedMsg struct {
	SectionId    int
	Repositories []Repository
}

func (repo Repository) ToTableRow() table.Row {
	return table.Row{
		repo.renderName(),
		repo.renderlabelDockerOfficial(),
		repo.renderlabelVerifiedPublisher(),
		repo.renderlabelOpenSourceProgram(),
		repo.renderOrganization(),
		repo.renderstatsDownloads(),
		repo.renderstatsStars(),
		repo.renderLastUpdate(),
		repo.renderDescription(),
	}
}

func (repo Repository) renderName() string {
	return repo.Data.Name
}
func (repo Repository) renderlabelDockerOfficial() string {
	return renderLabel(
		repo.Data.Labels.DockerOfficial,
		labelDockerOfficial,
		"")
}
func (repo Repository) renderlabelVerifiedPublisher() string {
	return renderLabel(
		repo.Data.Labels.VerifiedPublisher,
		labelVerifiedPublisher,
		"﫠")
}
func (repo Repository) renderlabelOpenSourceProgram() string {
	return renderLabel(
		repo.Data.Labels.OpenSourceProgram,
		labelOpenSourceProgram,
		"")
}
func (repo Repository) renderOrganization() string {
	return lipgloss.NewStyle().
		Render(repo.Data.Organization.Name)
}
func (repo Repository) renderstatsDownloads() string {
	return lipgloss.NewStyle().
		Render(repo.Data.Stats.Downloads)
}
func (repo Repository) renderstatsStars() string {
	return lipgloss.NewStyle().
		Render(repo.Data.Stats.Stars)
}
func (repo Repository) renderDescription() string {
	return lipgloss.NewStyle().
		Render(repo.Data.Description)
}
func (repo Repository) renderLastUpdate() string {
	return lipgloss.NewStyle().
		Render(utils.TimeElapsed(repo.Data.LastUpdate))
}

func renderLabel(enable bool, color lipgloss.AdaptiveColor, label string) string {
	if enable {
		return lipgloss.NewStyle().Foreground(color).Render(label)
	}
	return ""
}