package my_orgs_sidebar

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	data_user "github.com/AlyxPink/docker-hub-cli/internal/data/user"
	organization_user "github.com/AlyxPink/docker-hub-cli/internal/ui/components/organization/user"
	"github.com/AlyxPink/docker-hub-cli/internal/ui/components/sidebar"
	"github.com/AlyxPink/docker-hub-cli/internal/ui/context"
	"github.com/AlyxPink/docker-hub-cli/internal/ui/styles"
	"github.com/AlyxPink/docker-hub-cli/internal/utils"
)

type Model struct {
	organization *organization_user.Organization
	width        int
	ctx          *context.ProgramContext
}

func NewModel(data *data_user.Organization, width int, ctx *context.ProgramContext) Model {
	var o *organization_user.Organization
	if data == nil {
		o = nil
	} else {
		o = &organization_user.Organization{Data: *data}
	}
	return Model{
		organization: o,
		width:        width,
		ctx:          ctx,
	}
}

func (m Model) View() string {
	return lipgloss.JoinVertical(
		lipgloss.Left,
		m.renderTitle(),
		m.renderBadge(),
		m.renderLocation(),
		m.renderTimestamps(),
	)
}

func (m *Model) renderTitle() string {
	return lipgloss.JoinHorizontal(
		lipgloss.Center,
		m.renderName(),
	)
}

func (m *Model) renderBadge() string {
	var badge string
	switch m.organization.Data.Badge {
	case "verified_publisher":
		glyph := styles.DefaultGlyphs.LabelVerifiedPublisher
		text := fmt.Sprint(glyph, m.ctx.Localizer.L("my_orgs_label_verified_publisher"))
		color := styles.DefaultTheme.LabelVerifiedPublisher
		badge = lipgloss.NewStyle().Foreground(color).Render(text)
	case "open_source":
		glyph := styles.DefaultGlyphs.LabelOpenSourceProgram
		text := fmt.Sprint(glyph, m.ctx.Localizer.L("my_orgs_label_open_source_program"))
		color := styles.DefaultTheme.LabelOpenSourceProgram
		badge = lipgloss.NewStyle().Foreground(color).Render(text)
	}
	return lipgloss.JoinVertical(
		lipgloss.Top,
		badge,
	)
}

func (m *Model) renderName() string {
	return sidebar.Title.PaddingRight(1).Render(m.organization.Data.Name)
}

func (m *Model) renderLocation() string {
	return sidebar.TextBox.Render(m.organization.Data.Location)
}

func (m *Model) renderTimestamps() string {
	timestamps_txt := m.ctx.Localizer.L("my_orgs_sidebar_timestamps")
	created_at_txt := m.ctx.Localizer.L("my_orgs_sidebar_created_at")
	created_at := sidebar.AttributeName.Render(fmt.Sprintf("%s:", created_at_txt))
	return lipgloss.JoinVertical(
		lipgloss.Top,
		sidebar.Title.Render(fmt.Sprintf("%s:", timestamps_txt)),
		fmt.Sprintf("%s %s", created_at, utils.TimeElapsed(m.organization.Data.DateJoined)),
	)
}
