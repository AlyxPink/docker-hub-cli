package sidebar_repository

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/docker/hack-docker-access-management-cli/internal/data"
	"github.com/docker/hack-docker-access-management-cli/internal/ui/components/repository"
	"github.com/docker/hack-docker-access-management-cli/internal/ui/styles"
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
		render.WriteString(label.Copy().Foreground(labelDockerOfficial).Align(lipgloss.Left).Render(""))
	}
	if m.repo.Data.Labels.OpenSourceProgram {
		render.WriteString(label.Copy().Foreground(labelOpenSourceProgram).Align(lipgloss.Left).Render("﫠"))
	}
	if m.repo.Data.Labels.VerifiedPublisher {
		render.WriteString(label.Copy().Foreground(labelVerifiedPublisher).Align(lipgloss.Left).Render(""))
	}
	return render.String()
}

func (m *Model) renderTitle() string {
	return styles.MainTextStyle.Copy().
		Width(m.getIndentedContentWidth()).
		Padding(1).
		Bold(true).
		Render(m.renderLabels() + m.repo.Data.Name)
}

func (m *Model) renderArchs() string {
	render := strings.Builder{}
	render.WriteString(archTitle.Render("Archs:"))
	archs := []string{"arm64v8", "arm32v6", "arm32v7", "ppc64le", "s390x", "i386", "amd64"}
	for _, arch := range archs {
		render.WriteString(archTag.Render(arch))
	}
	return render.String()
}

func (m *Model) renderTags() string {
	render := strings.Builder{}
	render.WriteString(dockerImageTitle.Render("Tags:"))
	tags := []string{"3.16.0", "3.16", "3", "latest"}
	for _, tag := range tags {
		render.WriteString(dockerImageTag.Render(tag))
	}
	return render.String()
}

func (m *Model) renderPullCmd() string {
	render := strings.Builder{}
	render.WriteString(dockerPullCmdTitle.Render("Pull cmd:"))
	render.WriteString("\n")
	cmd := fmt.Sprintf("$ docker pull %s", m.repo.Data.Name)
	render.WriteString(dockerPullCmdBox.Copy().Render(cmd))

	return render.String()
}

func (m *Model) renderReadme() string {
	return styles.MainTextStyle.Copy().Width(m.getIndentedContentWidth()).
		Render("# Todo: Render README")
}

func (m *Model) getIndentedContentWidth() int {
	return m.width - 6
}
