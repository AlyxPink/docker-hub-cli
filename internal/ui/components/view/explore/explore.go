package view_explore

import (
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/victorbersy/docker-hub-cli/internal/data"
	"github.com/victorbersy/docker-hub-cli/internal/ui/components/repository"
	"github.com/victorbersy/docker-hub-cli/internal/ui/components/table"
	"github.com/victorbersy/docker-hub-cli/internal/ui/components/view"
	"github.com/victorbersy/docker-hub-cli/internal/ui/constants"
	"github.com/victorbersy/docker-hub-cli/internal/ui/context"
	"github.com/victorbersy/docker-hub-cli/internal/utils"
)

const ViewType = "explore"

type Model struct {
	Repositories []data.RepositoryData
	view         view.Model
}

func NewModel(id int, ctx *context.ProgramContext) Model {
	m := Model{
		Repositories: []data.RepositoryData{},
		view: view.Model{
			Id:        id,
			Ctx:       ctx,
			Spinner:   spinner.Model{Spinner: spinner.Dot},
			IsLoading: true,
			Type:      ViewType,
		},
	}

	m.view.Table = table.NewModel(
		m.view.GetDimensions(),
		m.GetViewColumns(),
		m.BuildRows(),
		"Repositories",
		utils.StringPtr(emptyStateStyle.Render("No repositories were found")),
	)

	return m
}

func (m *Model) Id() int {
	return m.view.Id
}

func (m Model) Update(msg tea.Msg) (view.View, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case ViewRepositoriesFetchedMsg:
		m.Repositories = msg.Repositories
		m.view.IsLoading = false
		m.view.Table.SetRows(m.BuildRows())

	case view.ViewTickMsg:
		if !m.view.IsLoading {
			return &m, nil
		}

		var internalTickCmd tea.Cmd
		m.view.Spinner, internalTickCmd = m.view.Spinner.Update(msg.InternalTickMsg)
		cmd = m.view.CreateNextTickCmd(internalTickCmd)
	}

	return &m, cmd
}

func (m *Model) View() string {
	var spinnerText *string
	if m.view.IsLoading {
		spinnerText = utils.StringPtr(lipgloss.JoinHorizontal(lipgloss.Top,
			spinnerStyle.Copy().Render(m.view.Spinner.View()),
			"Fetching Repositories...",
		))
	}

	return containerStyle.Copy().Render(
		m.view.Table.View(spinnerText),
	)
}

func (m *Model) UpdateProgramContext(ctx *context.ProgramContext) {
	m.view.UpdateProgramContext(ctx)
}

func renderColumnTitleLabels() string {
	return lipgloss.JoinHorizontal(
		lipgloss.Top,
		view.LabelTitle.Copy().Foreground(constants.ColorLabelDockerOfficial).Render(constants.GlyphLabelDockerOfficial),
		view.LabelTitle.Copy().Foreground(constants.ColorLabelVerifiedPublisher).Render(constants.GlyphLabelVerifiedPublisher),
		view.LabelTitle.Copy().Foreground(constants.ColorLabelOpenSourceProgram).Render(constants.GlyphLabelOpenSourceProgram),
	)
}

func (m *Model) GetViewColumns() []table.Column {
	return []table.Column{
		{
			Title: view.ColumnTitle.Render("Name"),
			Width: &nameWidth,
		},
		{
			Title: renderColumnTitleLabels(),
			Width: &labelsWidth,
		},
		{
			Title: view.ColumnTitle.Render("Organization"),
			Width: &organizationsnameWidth,
		},
		{
			Title: columnTitleStatsDownloads.Render(constants.GlyphStatsDownloads),
			Width: &statsWidth,
		},
		{
			Title: columnTitleStatsStars.Render(constants.GlyphStatsStars),
			Width: &statsWidth,
		},
		{
			Title: view.ColumnTitle.Render("Updated At"),
			Width: &LastUpdateCellWidth,
		},
		{
			Title: view.ColumnTitle.Render("Description"),
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

type ViewRepositoriesFetchedMsg struct {
	ViewId       int
	Repositories []data.RepositoryData
}

func (msg ViewRepositoriesFetchedMsg) GetViewId() int {
	return msg.ViewId
}

func (msg ViewRepositoriesFetchedMsg) GetViewType() string {
	return ViewType
}

func (m *Model) GetCurrRow() data.RowData {
	if len(m.Repositories) == 0 {
		return nil
	}
	repo := m.Repositories[m.view.Table.GetCurrItem()]
	return &repo
}

func (m *Model) NextRow() int {
	return m.view.NextRow()
}

func (m *Model) PrevRow() int {
	return m.view.PrevRow()
}

func (m *Model) FirstItem() int {
	return m.view.FirstItem()
}

func (m *Model) LastItem() int {
	return m.view.LastItem()
}

func (m *Model) FetchViewRows() tea.Cmd {
	if m == nil {
		return nil
	}
	m.Repositories = nil
	m.view.Table.ResetCurrItem()
	m.view.Table.Rows = nil
	m.view.IsLoading = true
	var cmds []tea.Cmd
	cmds = append(cmds, m.view.CreateNextTickCmd(spinner.Tick))

	cmds = append(cmds, func() tea.Msg {
		fetchedRepos, err := data.FetchRepositories()
		if err != nil {
			return ViewRepositoriesFetchedMsg{
				ViewId:       m.view.Id,
				Repositories: []data.RepositoryData{},
			}
		}

		return ViewRepositoriesFetchedMsg{
			ViewId:       m.view.Id,
			Repositories: fetchedRepos,
		}
	})

	return tea.Batch(cmds...)
}

func (m *Model) GetIsLoading() bool {
	return m.view.IsLoading
}

func Fetch(ctx context.ProgramContext) (view view.View, fetchCmd tea.Cmd) {
	viewModel := NewModel(0, &ctx)
	fetchCmd = viewModel.FetchViewRows()
	return &viewModel, tea.Batch(fetchCmd)
}
