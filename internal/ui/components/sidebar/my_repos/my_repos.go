package my_repos_sidebar

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	data_user "github.com/victorbersy/docker-hub-cli/internal/data/user"
	repository_user "github.com/victorbersy/docker-hub-cli/internal/ui/components/repository/user"
	"github.com/victorbersy/docker-hub-cli/internal/ui/components/sidebar"
	"github.com/victorbersy/docker-hub-cli/internal/ui/context"
	"github.com/victorbersy/docker-hub-cli/internal/ui/styles"
	"github.com/victorbersy/docker-hub-cli/internal/utils"
)

type Model struct {
	repo  *repository_user.Repository
	width int
	ctx   *context.ProgramContext
}

func NewModel(data *data_user.Repository, width int, ctx *context.ProgramContext) Model {
	var r *repository_user.Repository
	if data == nil {
		r = nil
	} else {
		r = &repository_user.Repository{Data: *data}
	}
	return Model{
		repo:  r,
		width: width,
		ctx:   ctx,
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
	return sidebar.Title.Render(m.repo.Data.Name)
}

func (m *Model) renderVisibility() string {
	visibility_txt := m.ctx.Localizer.L("my_repos_sidebar_visibility")
	private_txt := m.ctx.Localizer.L("my_repos_sidebar_visibility_private")
	public_txt := m.ctx.Localizer.L("my_repos_sidebar_visibility_public")
	var visibility string
	if m.repo.Data.IsPrivate {
		visibility = visibilityPrivate.Render(fmt.Sprintf("%s %s", styles.DefaultGlyphs.Private, private_txt))
	} else {
		visibility = visibilityPublic.Render(fmt.Sprintf("%s %s", styles.DefaultGlyphs.Public, public_txt))
	}
	return lipgloss.JoinVertical(
		lipgloss.Top,
		sidebar.Title.Render(fmt.Sprintf("%s:", visibility_txt)),
		visibility,
	)
}

func (m *Model) renderStats() string {
	return lipgloss.JoinVertical(
		lipgloss.Top,
		sidebar.Title.Render(fmt.Sprintf("%s:", m.ctx.Localizer.L("my_repos_sidebar_stats"))),
		lipgloss.JoinHorizontal(
			lipgloss.Top,
			statsStars.Render(styles.DefaultGlyphs.StatsStars),
			fmt.Sprint(m.repo.Data.StarCount),
		),
		lipgloss.JoinHorizontal(
			lipgloss.Top,
			statsDownloads.Render(styles.DefaultGlyphs.StatsDownloads),
			fmt.Sprint(m.repo.Data.PullCount),
		),
	)
}

func (m *Model) renderTimestamps() string {
	timestamps_txt := m.ctx.Localizer.L("my_repos_sidebar_timestamps")
	updated_at_txt := m.ctx.Localizer.L("my_repos_sidebar_updated_at")
	created_at_txt := m.ctx.Localizer.L("my_repos_sidebar_created_at")
	updated_at := sidebar.AttributeName.Render(fmt.Sprintf("%s:", updated_at_txt))
	created_at := sidebar.AttributeName.Render(fmt.Sprintf("%s:", created_at_txt))
	return lipgloss.JoinVertical(
		lipgloss.Top,
		sidebar.Title.Render(fmt.Sprintf("%s:", timestamps_txt)),
		fmt.Sprintf("%s %s", updated_at, utils.TimeElapsed(m.repo.Data.UpdatedAt)),
		fmt.Sprintf("%s %s", created_at, utils.TimeElapsed(m.repo.Data.CreatedAt)),
	)
}
