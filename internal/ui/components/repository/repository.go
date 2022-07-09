package repository

import (
	"fmt"

	"github.com/VictorBersy/docker-hub-cli/internal/data"
	"github.com/VictorBersy/docker-hub-cli/internal/ui/components/table"
	"github.com/VictorBersy/docker-hub-cli/internal/ui/constants"
	"github.com/VictorBersy/docker-hub-cli/internal/utils"
	"github.com/charmbracelet/lipgloss"
)

type Repository struct {
	Data data.RepositoryData
}

func (repo Repository) ToTableRow() table.Row {
	return table.Row{
		repo.renderName(),
		repo.renderLabelDockerOfficial(),
		repo.renderLabelVerifiedPublisher(),
		repo.renderLabelOpenSourceProgram(),
		repo.renderPublisher(),
		repo.renderstatsDownloads(),
		repo.renderstatsStars(),
		repo.renderLastUpdate(),
		repo.renderDescription(),
	}
}

func (repo Repository) renderName() string {
	return repo.Data.Name
}

func (repo Repository) renderLabelDockerOfficial() string {
	return renderLabel(repo.Data.Labels.DockerOfficial, labelDockerOfficial, constants.LabelDockerOfficialGlyph)
}

func (repo Repository) renderLabelVerifiedPublisher() string {
	return renderLabel(repo.Data.Labels.VerifiedPublisher, labelVerifiedPublisher, constants.LabelVerifiedPublisherGlyph)
}

func (repo Repository) renderLabelOpenSourceProgram() string {
	return renderLabel(repo.Data.Labels.OpenSourceProgram, labelOpenSourceProgram, constants.LabelOpenSourceProgramGlyph)
}

func (repo Repository) renderPublisher() string {
	return lipgloss.NewStyle().
		Render(repo.Data.Publisher.Name)
}

func (repo Repository) renderstatsDownloads() string {
	return lipgloss.NewStyle().
		Render(repo.Data.PullCount)
}

func (repo Repository) renderstatsStars() string {
	return lipgloss.NewStyle().
		Render(fmt.Sprint(repo.Data.StarCount))
}

func (repo Repository) renderDescription() string {
	return lipgloss.NewStyle().
		Render(repo.Data.Description)
}

func (repo Repository) renderLastUpdate() string {
	return lipgloss.NewStyle().
		Render(utils.TimeElapsed(repo.Data.Updated_at))
}

func renderLabel(enable bool, color lipgloss.AdaptiveColor, label string) string {
	if enable {
		return lipgloss.NewStyle().Foreground(color).Width(3).Render(label)
	}
	return lipgloss.NewStyle().Width(1).Render("")
}
