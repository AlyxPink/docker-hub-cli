package sidebar_explore

import (
	"fmt"
	"strings"

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
	s.WriteString(m.renderName())
	s.WriteString("\n")
	s.WriteString(m.renderLabels())
	s.WriteString("\n")
	s.WriteString(m.renderOrganization())
	s.WriteString("\n")
	s.WriteString(m.renderStats())
	s.WriteString("\n")
	s.WriteString(m.renderDescription())
	s.WriteString("\n")
	s.WriteString(m.renderLastUpdate())
	s.WriteString("\n")

	return s.String()
}

func (m *Model) renderName() string {
	return styles.MainTextStyle.Copy().Width(m.getIndentedContentWidth()).
		Render(m.repo.Data.Name)
}
func (m *Model) renderLabels() string {
	return styles.MainTextStyle.Copy().Width(m.getIndentedContentWidth()).
		Render(fmt.Sprint(m.repo.Data.Labels))
}
func (m *Model) renderOrganization() string {
	return styles.MainTextStyle.Copy().Width(m.getIndentedContentWidth()).
		Render(fmt.Sprint(m.repo.Data.Organization))
}
func (m *Model) renderStats() string {
	return styles.MainTextStyle.Copy().Width(m.getIndentedContentWidth()).
		Render(fmt.Sprint(m.repo.Data.Stats))
}
func (m *Model) renderDescription() string {
	return styles.MainTextStyle.Copy().Width(m.getIndentedContentWidth()).
		Render(m.repo.Data.Description)
}
func (m *Model) renderLastUpdate() string {
	return styles.MainTextStyle.Copy().Width(m.getIndentedContentWidth()).
		Render(fmt.Sprint(m.repo.Data.LastUpdate))
}

func (m *Model) getIndentedContentWidth() int {
	return m.width - 6
}
