package my_repos_sidebar

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	data_user "github.com/victorbersy/docker-hub-cli/internal/data/user"
	repository_user "github.com/victorbersy/docker-hub-cli/internal/ui/components/repository/user"
	"github.com/victorbersy/docker-hub-cli/internal/ui/constants"
	"github.com/victorbersy/docker-hub-cli/internal/utils"
)

type Model struct {
	repo  *repository_user.Repository
	width int
}

func NewModel(data *data_user.Repository, width int) Model {
	var r *repository_user.Repository
	if data == nil {
		r = nil
	} else {
		r = &repository_user.Repository{Data: *data}
	}
	return Model{
		repo:  r,
		width: width,
	}
}

func (m Model) View() string {
	return lipgloss.JoinVertical(
		lipgloss.Left,
		m.renderTitle(),
		m.renderVisibility(),
		m.renderStats(),
		m.renderTimestamps(),
	)
}

func (m *Model) renderTitle() string {
	return lipgloss.JoinHorizontal(
		lipgloss.Center,
		m.renderName(),
	)
}

func (m *Model) renderName() string {
	return titleRepo.Render(m.repo.Data.Name)
}

func (m *Model) renderVisibility() string {
	var visibility string
	if m.repo.Data.IsPrivate {
		visibility = visibilityPrivate.Render(fmt.Sprintf("%s Private", constants.GlyphPrivate))
	} else {
		visibility = visibilityPublic.Render(fmt.Sprintf("%s Public", constants.GlyphPublic))
	}
	return lipgloss.JoinVertical(
		lipgloss.Top,
		visibilityTitle.Render("Visibility:"),
		visibility,
	)
}

func (m *Model) renderStats() string {
	return lipgloss.JoinVertical(
		lipgloss.Top,
		statsTitle.Render("Stats:"),
		lipgloss.JoinHorizontal(
			lipgloss.Top,
			statsStars.Render(constants.GlyphStatsStars),
			fmt.Sprint(m.repo.Data.StarCount),
		),
		lipgloss.JoinHorizontal(
			lipgloss.Top,
			statsDownloads.Render(constants.GlyphStatsDownloads),
			fmt.Sprint(m.repo.Data.PullCount),
		),
	)
}

func (m *Model) renderTimestamps() string {
	updated_at := timestampName.Render("Last update:")
	created_at := timestampName.Render("Created at:")
	return lipgloss.JoinVertical(
		lipgloss.Top,
		timestampTitle.Render("Timestamps:"),
		fmt.Sprintf("%s %s", updated_at, utils.TimeElapsed(m.repo.Data.UpdatedAt)),
		fmt.Sprintf("%s %s", created_at, utils.TimeElapsed(m.repo.Data.CreatedAt)),
	)
}
