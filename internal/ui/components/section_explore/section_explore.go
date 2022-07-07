package section_explore

import (
	"fmt"
	"sort"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/docker/hack-docker-access-management-cli/internal/config"
	"github.com/docker/hack-docker-access-management-cli/internal/data"
	"github.com/docker/hack-docker-access-management-cli/internal/ui/components/repository"
	"github.com/docker/hack-docker-access-management-cli/internal/ui/components/section"
	"github.com/docker/hack-docker-access-management-cli/internal/ui/components/table"
	"github.com/docker/hack-docker-access-management-cli/internal/ui/context"
	"github.com/docker/hack-docker-access-management-cli/internal/utils"
)

const SectionType = "explore"

type Model struct {
	Repositories []data.RepositoryData
	section      section.Model
}

func NewModel(id int, ctx *context.ProgramContext, config config.SectionConfig) Model {
	m := Model{
		Repositories: []data.RepositoryData{},
		section: section.Model{
			Id:        id,
			Config:    config,
			Ctx:       ctx,
			Spinner:   spinner.Model{Spinner: spinner.Dot},
			IsLoading: true,
			Type:      SectionType,
		},
	}

	m.section.Table = table.NewModel(
		m.section.GetDimensions(),
		m.GetSectionColumns(),
		m.BuildRows(),
		"PR",
		utils.StringPtr(emptyStateStyle.Render(fmt.Sprintf(
			"No PRs were found that match the given filters: %s",
			lipgloss.NewStyle().Italic(true).Render(m.section.Config.Title),
		))),
	)

	return m
}

func (m *Model) Id() int {
	return m.section.Id
}

func (m Model) Update(msg tea.Msg) (section.Section, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case SectionRepositoriesFetchedMsg:
		m.Repositories = msg.Repositories
		m.section.IsLoading = false
		m.section.Table.SetRows(m.BuildRows())

	case section.SectionTickMsg:
		if !m.section.IsLoading {
			return &m, nil
		}

		var internalTickCmd tea.Cmd
		m.section.Spinner, internalTickCmd = m.section.Spinner.Update(msg.InternalTickMsg)
		cmd = m.section.CreateNextTickCmd(internalTickCmd)
	}

	return &m, cmd
}

func (m *Model) View() string {
	var spinnerText *string
	if m.section.IsLoading {
		spinnerText = utils.StringPtr(lipgloss.JoinHorizontal(lipgloss.Top,
			spinnerStyle.Copy().Render(m.section.Spinner.View()),
			"Fetching Pull Requests...",
		))
	}

	return containerStyle.Copy().Render(
		m.section.Table.View(spinnerText),
	)
}

func (m *Model) UpdateProgramContext(ctx *context.ProgramContext) {
	m.section.UpdateProgramContext(ctx)
}

func (m *Model) GetSectionColumns() []table.Column {
	columnTitle := lipgloss.NewStyle().Bold(true)
	return []table.Column{
		{
			Title: columnTitle.Copy().Render("Name"),
			Width: &nameWidth,
		},
		{
			Title: columnTitle.Copy().Foreground(labelDockerOfficial).Render(""),
			Width: &labelDockerOfficialWidth,
		},
		{
			Title: columnTitle.Copy().Foreground(labelVerifiedPublisher).Render("﫠"),
			Width: &labelVerifiedPublisherWidth,
		},
		{
			Title: columnTitle.Copy().Foreground(labelOpenSourceProgram).Render(""),
			Width: &labelOpenSourceProgramWidth,
		},
		{
			Title: columnTitle.Copy().Render("Organization"),
			Width: &organizationsnameWidth,
		},
		{
			Title: columnTitle.Copy().Foreground(statsDownloads).Render(""),
			Width: &statsWidth,
		},
		{
			Title: columnTitle.Copy().Foreground(statsStars).Render(""),
			Width: &statsWidth,
		},
		{
			Title: columnTitle.Copy().Render("Updated At"),
			Width: &LastUpdateCellWidth,
		},
		{
			Title: columnTitle.Copy().Render("Description"),
			Grow:  utils.BoolPtr(true),
		},
	}
}

func (m *Model) BuildRows() []table.Row {
	var rows []table.Row
	for _, currRepo := range m.Repositories {
		repoModel := repository.Repository{Data: currRepo}
		rows = append(rows, repoModel.ToTableRow())
	}

	return rows
}

func (m *Model) NumRows() int {
	return len(m.Repositories)
}

type SectionRepositoriesFetchedMsg struct {
	SectionId    int
	Repositories []data.RepositoryData
}

func (msg SectionRepositoriesFetchedMsg) GetSectionId() int {
	return msg.SectionId
}

func (msg SectionRepositoriesFetchedMsg) GetSectionType() string {
	return SectionType
}

func (m *Model) GetCurrRow() data.RowData {
	if len(m.Repositories) == 0 {
		return nil
	}
	repo := m.Repositories[m.section.Table.GetCurrItem()]
	return &repo
}

func (m *Model) NextRow() int {
	return m.section.NextRow()
}

func (m *Model) PrevRow() int {
	return m.section.PrevRow()
}

func (m *Model) FirstItem() int {
	return m.section.FirstItem()
}

func (m *Model) LastItem() int {
	return m.section.LastItem()
}

func (m *Model) FetchSectionRows() tea.Cmd {
	if m == nil {
		return nil
	}
	m.Repositories = nil
	m.section.Table.ResetCurrItem()
	m.section.Table.Rows = nil
	m.section.IsLoading = true
	var cmds []tea.Cmd
	cmds = append(cmds, m.section.CreateNextTickCmd(spinner.Tick))

	cmds = append(cmds, func() tea.Msg {
		limit := m.section.Config.Limit
		if limit == nil {
			limit = &m.section.Ctx.Config.Defaults.PrsLimit
		}
		fetchedPrs, err := data.FetchRepositories(m.section.Config.Title, *limit)
		if err != nil {
			return SectionRepositoriesFetchedMsg{
				SectionId:    m.section.Id,
				Repositories: []data.RepositoryData{},
			}
		}

		sort.Slice(fetchedPrs, func(i, j int) bool {
			return fetchedPrs[i].LastUpdate.After(fetchedPrs[j].LastUpdate)
		})
		return SectionRepositoriesFetchedMsg{
			SectionId:    m.section.Id,
			Repositories: fetchedPrs,
		}
	})

	return tea.Batch(cmds...)
}

func (m *Model) GetIsLoading() bool {
	return m.section.IsLoading
}

func FetchAllSections(ctx context.ProgramContext) (sections []section.Section, fetchAllCmd tea.Cmd) {
	fetchPRsCmds := make([]tea.Cmd, 0, len(ctx.Config.PRSections))
	sections = make([]section.Section, 0, len(ctx.Config.PRSections))
	for i, sectionConfig := range ctx.Config.PRSections {
		sectionModel := NewModel(i, &ctx, sectionConfig)
		sections = append(sections, &sectionModel)
		fetchPRsCmds = append(fetchPRsCmds, sectionModel.FetchSectionRows())
	}
	return sections, tea.Batch(fetchPRsCmds...)
}
