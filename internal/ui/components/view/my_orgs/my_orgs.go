package my_orgs

import (
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/AlyxPink/docker-hub-cli/internal/data"
	data_user "github.com/AlyxPink/docker-hub-cli/internal/data/user"
	organization_user "github.com/AlyxPink/docker-hub-cli/internal/ui/components/organization/user"
	"github.com/AlyxPink/docker-hub-cli/internal/ui/components/table"
	"github.com/AlyxPink/docker-hub-cli/internal/ui/components/view"
	"github.com/AlyxPink/docker-hub-cli/internal/ui/context"
	"github.com/AlyxPink/docker-hub-cli/internal/utils"
)

const ViewType = "my_orgs"

type Model struct {
	Organizations []data_user.Organization
	view          view.Model
}

func NewModel(id int, ctx *context.ProgramContext) Model {
	m := Model{
		Organizations: []data_user.Organization{},
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
		m.view.Ctx.Localizer.L("my_orgs_item_type_label"),
		utils.StringPtr(
			lipgloss.JoinVertical(
				lipgloss.Top,
				emptyStateStyle.Render(m.view.Ctx.Localizer.L("my_orgs_not_found")),
				emptyStateStyle.Render(m.view.Ctx.Localizer.L("my_orgs_not_found_tip")),
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
	case ViewOrganizationsFetchedMsg:
		m.Organizations = msg.Organizations
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
			m.view.Ctx.Localizer.L("my_orgs_fetching"),
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
			Title: view.ColumnTitle.Render(m.view.Ctx.Localizer.L("column_header_name")),
			Grow:  utils.BoolPtr(true),
		},
		{
			Title: view.ColumnTitle.Render(m.view.Ctx.Localizer.L("column_header_badge")),
			Grow:  utils.BoolPtr(true),
		},
		{
			Title: view.ColumnTitle.Render(m.view.Ctx.Localizer.L("column_header_created_at")),
			Grow:  utils.BoolPtr(true),
		},
	}
}

func (m *Model) BuildRows() []table.Row {
	var rows []table.Row
	for _, currOrg := range m.Organizations {
		orgModel := organization_user.Organization{Data: currOrg}
		rows = append(rows, orgModel.ToTableRow())
	}

	return rows
}

func (m *Model) NumRows() int {
	return len(m.Organizations)
}

type ViewOrganizationsFetchedMsg struct {
	ViewId        int
	Organizations []data_user.Organization
}

func (msg ViewOrganizationsFetchedMsg) GetViewId() int {
	return msg.ViewId
}

func (msg ViewOrganizationsFetchedMsg) GetViewType() string {
	return ViewType
}

func (m *Model) GetCurrRow() data.RowData {
	if len(m.Organizations) == 0 {
		return nil
	}
	org := m.Organizations[m.view.Table.GetCurrItem()]
	return &org
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
	m.Organizations = nil
	m.view.Table.ResetCurrItem()
	m.view.Table.Rows = nil
	m.view.IsLoading = true
	var cmds []tea.Cmd
	cmds = append(cmds, m.view.CreateNextTickCmd(spinner.Tick))

	cmds = append(cmds, func() tea.Msg {
		fetchedOrgs, err := data_user.FetchOrganizations()
		if err != nil {
			return ViewOrganizationsFetchedMsg{
				ViewId:        m.view.Id,
				Organizations: []data_user.Organization{},
			}
		}

		return ViewOrganizationsFetchedMsg{
			ViewId:        m.view.Id,
			Organizations: fetchedOrgs,
		}
	})

	return tea.Batch(cmds...)
}

func (m *Model) GetIsLoading() bool {
	return m.view.IsLoading
}

func Fetch(id int, ctx context.ProgramContext) (view view.View, fetchCmd tea.Cmd) {
	viewModel := NewModel(id, &ctx)
	fetchCmd = viewModel.FetchViewRows()
	return &viewModel, tea.Batch(fetchCmd)
}
