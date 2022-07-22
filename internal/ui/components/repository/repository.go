package repository

import (
	"fmt"

	"github.com/VictorBersy/docker-hub-cli/internal/data"
	"github.com/VictorBersy/docker-hub-cli/internal/ui/components/table"
	"github.com/VictorBersy/docker-hub-cli/internal/utils"
	"github.com/charmbracelet/lipgloss"
)

type Repository struct {
	Data data.RepositoryData
}

func (repo Repository) ToTableRow() table.Row {
	return table.Row{
		repo.renderName(),
		repo.renderLabels(),
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

func (repo Repository) renderLabels() string {
	labels := []string{}
	for _, label := range repo.Data.Labels {
		if label.Enabled {
			labels = append(labels, lipgloss.NewStyle().Foreground(label.Color).Width(3).Render(label.Glyph))
		}
	}
	return lipgloss.JoinHorizontal(
		lipgloss.Top,
		labels...,
	)
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
