package my_repos

import (
	"github.com/VictorBersy/docker-hub-cli/internal/data"
	"github.com/VictorBersy/docker-hub-cli/internal/ui/components/table"
	"github.com/VictorBersy/docker-hub-cli/internal/ui/components/view"
	"github.com/VictorBersy/docker-hub-cli/internal/ui/context"
	"github.com/VictorBersy/docker-hub-cli/internal/utils"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const ViewType = "my_repos"

type Model struct {
	view view.Model
}

func NewModel(id int, ctx *context.ProgramContext) Model {
	m := Model{
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
		utils.StringPtr(emptyStateStyle.Render("Nothing found")),
	)

	return m
}

func (m *Model) Id() int {
	return m.view.Id
}

func (m Model) Update(msg tea.Msg) (view.View, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case ViewFetchedMsg:
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

func (m *Model) GetViewColumns() []table.Column {
	return []table.Column{
		{
			Title: view.ColumnTitle.Render("TEST TODO"),
			Width: &testWidth,
		},
	}
}

func (m *Model) BuildRows() []table.Row {
	var rows []table.Row
	return rows
}

func (m *Model) NumRows() int {
	return 123
}

type ViewFetchedMsg struct {
	ViewId int
}

func (msg ViewFetchedMsg) GetViewId() int {
	return msg.ViewId
}

func (msg ViewFetchedMsg) GetViewType() string {
	return ViewType
}

func (m *Model) GetCurrRow() data.RowData {
	return nil
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
	return nil
}

func (m *Model) GetIsLoading() bool {
	return m.view.IsLoading
}

func Fetch(ctx context.ProgramContext) (view view.View, fetchCmd tea.Cmd) {
	viewModel := NewModel(1, &ctx)
	fetchCmd = viewModel.FetchViewRows()
	return &viewModel, tea.Batch(fetchCmd)
}
