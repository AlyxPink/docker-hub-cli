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
		visibility = lipgloss.NewStyle().Foreground(lipgloss.Color("#EF476F")).Padding(1, 0, 2, 0).Render(fmt.Sprintf("%s Private", constants.GlyphPrivate))
	} else {
		visibility = lipgloss.NewStyle().Foreground(lipgloss.Color("#06D6A0")).Padding(1, 0, 2, 0).Render(fmt.Sprintf("%s Public", constants.GlyphPublic))
	}
	return lipgloss.JoinVertical(
		lipgloss.Top,
		lipgloss.NewStyle().Bold(true).Underline(true).Render("Visibility:"),
		visibility,
	)
}

func (m *Model) renderStats() string {
	return lipgloss.JoinVertical(
		lipgloss.Top,
		lipgloss.NewStyle().Bold(true).Underline(true).PaddingBottom(1).Render("Stats:"),
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
	updated_at := lipgloss.NewStyle().Bold(true).Render("Last update:")
	created_at := lipgloss.NewStyle().Bold(true).Render("Created at:")
	return lipgloss.JoinVertical(
		lipgloss.Top,
		lipgloss.NewStyle().Bold(true).Underline(true).Padding(1, 0).Render("Timestamps:"),
		lipgloss.NewStyle().Render(fmt.Sprintf("%s %s", updated_at, utils.TimeElapsed(m.repo.Data.UpdatedAt))),
		lipgloss.NewStyle().Render(fmt.Sprintf("%s %s", created_at, utils.TimeElapsed(m.repo.Data.CreatedAt))),
	)
}
