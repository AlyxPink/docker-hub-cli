package my_repos

import (
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/victorbersy/docker-hub-cli/internal/data"
	data_user "github.com/victorbersy/docker-hub-cli/internal/data/user"
	repository_user "github.com/victorbersy/docker-hub-cli/internal/ui/components/repository/user"
	"github.com/victorbersy/docker-hub-cli/internal/ui/components/table"
	"github.com/victorbersy/docker-hub-cli/internal/ui/components/view"
	"github.com/victorbersy/docker-hub-cli/internal/ui/context"
	"github.com/victorbersy/docker-hub-cli/internal/ui/styles"
	"github.com/victorbersy/docker-hub-cli/internal/utils"
)

const ViewType = "my_repos"

type Model struct {
	Repositories []data_user.Repository
	view         view.Model
}

func NewModel(id int, ctx *context.ProgramContext) Model {
	m := Model{
		Repositories: []data_user.Repository{},
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
		m.view.Ctx.Localizer.T("my_repositories_item_type_label"),
		utils.StringPtr(
			lipgloss.JoinVertical(
				lipgloss.Top,
				emptyStateStyle.Render(m.view.Ctx.Localizer.T("my_repositories_not_found")),
				emptyStateStyle.Render(m.view.Ctx.Localizer.T("my_repositories_not_found_tip")),
			),
		),
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
			m.view.Ctx.Localizer.T("my_repositories_fetching"),
		))
	}

	return containerStyle.Copy().Render(
		m.view.Table.View(spinnerText),
	)
}

func (m *Model) UpdateProgramContext(ctx *context.ProgramContext) {
	m.view.UpdateProgramContext(ctx)
}

func (m *Model) GetViewColumns() []table.Column {
	return []table.Column{
		{
			Title: view.ColumnTitle.Render(m.view.Ctx.Localizer.T("column_header_name")),
			Grow:  utils.BoolPtr(true),
		},
		{
			Title: columnTitleIsPrivate.Render(styles.DefaultGlyphs.IsPrivate),
			Width: &isPrivateWidth,
		},
		{
			Title: columnTitleStatsDownloads.Render(styles.DefaultGlyphs.StatsDownloads),
			Width: &statsWidth,
		},
		{
			Title: columnTitleStatsStars.Render(styles.DefaultGlyphs.StatsStars),
			Width: &statsWidth,
		},
		{
			Title: view.ColumnTitle.Render(m.view.Ctx.Localizer.T("column_header_updated_at")),
			Width: &updatedAtWidth,
		},
		{
			Title: view.ColumnTitle.Render(m.view.Ctx.Localizer.T("column_header_created_at")),
			Width: &createdAtWidth,
		},
	}
}

func (m *Model) BuildRows() []table.Row {
	var rows []table.Row
	for _, currRepo := range m.Repositories {
		repoModel := repository_user.Repository{Data: currRepo}
		rows = append(rows, repoModel.ToTableRow())
	}

	return rows
}

func (m *Model) NumRows() int {
	return len(m.Repositories)
}

type ViewRepositoriesFetchedMsg struct {
	ViewId       int
	Repositories []data_user.Repository
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
		fetchedRepos, err := data_user.FetchRepositories()
		if err != nil {
			return ViewRepositoriesFetchedMsg{
				ViewId:       m.view.Id,
				Repositories: []data_user.Repository{},
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
	viewModel := NewModel(1, &ctx)
	fetchCmd = viewModel.FetchViewRows()
	return &viewModel, tea.Batch(fetchCmd)
}
