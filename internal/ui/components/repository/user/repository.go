package repository_user

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	data_user "github.com/victorbersy/docker-hub-cli/internal/data/user"
	"github.com/victorbersy/docker-hub-cli/internal/ui/components/table"
	"github.com/victorbersy/docker-hub-cli/internal/ui/constants"
	"github.com/victorbersy/docker-hub-cli/internal/utils"
)

type Repository struct {
	Data data_user.Repository
}

func (repo Repository) ToTableRow() table.Row {
	return table.Row{
		repo.renderName(),
		repo.renderIsPrivate(),
		repo.renderstatsDownloads(),
		repo.renderstatsStars(),
		repo.renderLastUpdate(),
		repo.renderCreatedAt(),
	}
}

func (repo Repository) renderName() string {
	return repo.Data.Name
}

func (repo Repository) renderIsPrivate() string {
	if repo.Data.IsPrivate {
		return lipgloss.NewStyle().Foreground(lipgloss.Color("#EF476F")).Render(constants.GlyphPrivate)
	} else {
		return lipgloss.NewStyle().Foreground(lipgloss.Color("#06D6A0")).Render(constants.GlyphPublic)
	}
}

func (repo Repository) renderstatsDownloads() string {
	return lipgloss.NewStyle().
		Render(fmt.Sprint(repo.Data.PullCount))
}

func (repo Repository) renderstatsStars() string {
	return lipgloss.NewStyle().
		Render(fmt.Sprint(repo.Data.StarCount))
}

func (repo Repository) renderLastUpdate() string {
	return lipgloss.NewStyle().
		Render(utils.TimeElapsed(repo.Data.UpdatedAt))
}

func (repo Repository) renderCreatedAt() string {
	return lipgloss.NewStyle().
		Render(utils.TimeElapsed(repo.Data.CreatedAt))
}
