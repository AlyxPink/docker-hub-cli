package sidebar_repository

import (
	"fmt"
	"strings"

	"github.com/VictorBersy/docker-hub-cli/internal/data"
	"github.com/VictorBersy/docker-hub-cli/internal/ui/components/repository"
	"github.com/VictorBersy/docker-hub-cli/internal/ui/constants"
	"github.com/VictorBersy/docker-hub-cli/internal/ui/styles"
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
	s := strings.Builder{}
	s.WriteString(m.renderTitle())
	s.WriteString("\n")
	s.WriteString(m.renderArchs())
	s.WriteString("\n")
	s.WriteString("\n")
	s.WriteString(m.renderTags())
	s.WriteString("\n")
	s.WriteString("\n")
	s.WriteString(m.renderPullCmd())
	s.WriteString("\n")
	s.WriteString(m.renderReadme())
	s.WriteString("\n")

	return s.String()
}

func (m *Model) renderLabels() string {
	render := strings.Builder{}
	if m.repo.Data.Labels.DockerOfficial {
		render.WriteString(label.Copy().Foreground(labelDockerOfficial).Align(lipgloss.Left).Render(constants.LabelDockerOfficialGlyph))
	}
	if m.repo.Data.Labels.VerifiedPublisher {
		render.WriteString(label.Copy().Foreground(labelVerifiedPublisher).Align(lipgloss.Left).Render(constants.LabelVerifiedPublisherGlyph))
	}
	if m.repo.Data.Labels.OpenSourceProgram {
		render.WriteString(label.Copy().Foreground(labelOpenSourceProgram).Align(lipgloss.Left).Render(constants.LabelOpenSourceProgramGlyph))
	}
	return render.String()
}

func (m *Model) renderTitle() string {
	return styles.MainTextStyle.Copy().
		Padding(1).
		Bold(true).
		Render(m.renderName() + m.renderLabels())
}

func (m *Model) renderName() string {
	return styles.MainTextStyle.Copy().
		Bold(true).
		Render(m.repo.Data.Name)
}

func (m *Model) renderArchs() string {
	render := strings.Builder{}
	render.WriteString(archTitle.Render("Archs"))
	render.WriteString("\n")
	archs := m.repo.Data.Architectures
	for _, arch := range archs {
		render.WriteString(archTag.Render(arch.Label))
	}
	return render.String()
}

func (m *Model) renderTags() string {
	render := strings.Builder{}
	render.WriteString(dockerImageTitle.Render("Tags"))
	render.WriteString("\n")
	tags := []string{"3.16.0", "3.16", "3", "latest"}
	for _, tag := range tags {
		render.WriteString(dockerImageTag.Render(tag))
	}
	return render.String()
}

func (m *Model) renderPullCmd() string {
	render := strings.Builder{}
	render.WriteString(dockerPullCmdTitle.Render("Pull cmd"))
	render.WriteString("\n")
	cmd := fmt.Sprintf("$ docker pull %s", m.repo.Data.Slug)
	render.WriteString(dockerPullCmdBox.Copy().Render(cmd))

	return render.String()
}

func (m *Model) renderReadme() string {
	return styles.MainTextStyle.Copy().
		Render("# Todo: Render README")
}
