package explore_sidebar

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	data_search "github.com/AlyxPink/docker-hub-cli/internal/data/search"
	repository_search "github.com/AlyxPink/docker-hub-cli/internal/ui/components/repository/search"
	"github.com/AlyxPink/docker-hub-cli/internal/ui/components/sidebar"
	"github.com/AlyxPink/docker-hub-cli/internal/ui/context"
)

type Model struct {
	repo  *repository_search.Repository
	width int
	ctx   *context.ProgramContext
}

func NewModel(data *data_search.Repository, width int, ctx *context.ProgramContext) Model {
	var r *repository_search.Repository
	if data == nil {
		r = nil
	} else {
		r = &repository_search.Repository{Data: *data}
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
		m.renderLabels(),
		m.renderDescription(),
		m.renderArchs(),
		m.renderPullCmd(),
	)
}

func (m *Model) renderLabels() string {
	labels := []string{}
	for _, label := range m.repo.Data.Labels {
		if label.Enabled {
			labels = append(labels, lipgloss.JoinHorizontal(
				lipgloss.Top,
				labelGlyph.Foreground(label.Color).Render(label.Glyph),
				labelGlyph.Foreground(label.Color).Render(label.Name),
			))
		}
	}
	return lipgloss.JoinVertical(
		lipgloss.Top,
		labels...,
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

func (m *Model) renderDescription() string {
	return sidebar.TextBox.Render(m.repo.Data.Description)
}

func (m *Model) renderArchs() string {
	architecture_title := m.ctx.Localizer.L("explore_sidebar_architectures")
	archs := []string{}
	for _, arch := range m.repo.Data.Architectures {
		archs = append(archs, archLabel.Render(arch.Name))
	}
	return lipgloss.JoinVertical(
		lipgloss.Top,
		archsTitle.Render(architecture_title),
		lipgloss.JoinHorizontal(
			lipgloss.Top,
			archs...,
		),
	)
}

func (m *Model) renderPullCmd() string {
	howToPullInstructions := &i18n.Message{ID: "explore_sidebar_how_to_pull_instructions"}
	howToPullInstructionsTemplate := map[string]interface{}{"Name": m.repo.Data.Slug, "Count": 1}
	how_to_txt := m.ctx.Localizer.Localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: howToPullInstructions,
		TemplateData:   howToPullInstructionsTemplate,
	})
	cmd := fmt.Sprintf("$ docker pull %s", m.repo.Data.Slug)
	return lipgloss.JoinVertical(
		lipgloss.Top,
		dockerPullCmdTitle.Render(how_to_txt),
		dockerPullCmdBox.Copy().Render(cmd),
	)
}
