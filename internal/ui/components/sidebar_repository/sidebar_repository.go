package sidebar_repository

import (
	"fmt"

	"github.com/VictorBersy/docker-hub-cli/internal/data"
	"github.com/VictorBersy/docker-hub-cli/internal/ui/components/repository"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	repo  *repository.Repository
	width int
}

func NewModel(data *data.RepositoryData, width int) Model {
	var r *repository.Repository
	if data == nil {
		r = nil
	} else {
		r = &repository.Repository{Data: *data}
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
		m.renderDescription(),
		m.renderArchs(),
		m.renderPullCmd(),
	)
}

func (m *Model) renderLabels() string {
	labels := []string{}
	for _, label := range m.repo.Data.Labels {
		if label.Enabled {
			labels = append(labels, labelGlyph.Foreground(label.Color).Render(label.Glyph))
		}
	}
	return lipgloss.JoinHorizontal(
		lipgloss.Top,
		labels...,
	)
}

func (m *Model) renderTitle() string {
	return lipgloss.JoinHorizontal(
		lipgloss.Top,
		m.renderName(),
		m.renderLabels(),
	)
}

func (m *Model) renderName() string {
	return titleRepo.Render(m.repo.Data.Name)
}

func (m *Model) renderDescription() string {
	return description.Render(m.repo.Data.Description)
}

func (m *Model) renderArchs() string {
	archs := []string{}
	for _, arch := range m.repo.Data.Architectures {
		archs = append(archs, archLabel.Render(arch.Name))
	}
	return lipgloss.JoinVertical(
		lipgloss.Top,
		archsTitle.Render("Archs"),
		lipgloss.JoinHorizontal(
			lipgloss.Top,
			archs...,
		),
	)
}

func (m *Model) renderPullCmd() string {
	cmd := fmt.Sprintf("$ docker pull %s", m.repo.Data.Slug)
	return lipgloss.JoinVertical(
		lipgloss.Top,
		dockerPullCmdTitle.Render(fmt.Sprintf("How to pull %s?", m.repo.Data.Name)),
		dockerPullCmdBox.Copy().Render(cmd),
	)
}
