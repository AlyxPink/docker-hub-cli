package view

import (
	"github.com/VictorBersy/docker-hub-cli/internal/data"
	"github.com/VictorBersy/docker-hub-cli/internal/ui/components/table"
	"github.com/VictorBersy/docker-hub-cli/internal/ui/constants"
	"github.com/VictorBersy/docker-hub-cli/internal/ui/context"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	Id        int
	Ctx       *context.ProgramContext
	Spinner   spinner.Model
	IsLoading bool
	Table     table.Model
	Type      string
}

type View interface {
	Id() int
	Update(msg tea.Msg) (View, tea.Cmd)
	View() string
	NumRows() int
	GetCurrRow() data.RowData
	NextRow() int
	PrevRow() int
	FirstItem() int
	LastItem() int
	FetchViewRows() tea.Cmd
	GetIsLoading() bool
	GetViewColumns() []table.Column
	BuildRows() []table.Row
	UpdateProgramContext(ctx *context.ProgramContext)
}

func (m *Model) CreateNextTickCmd(nextTickCmd tea.Cmd) tea.Cmd {
	if m == nil || nextTickCmd == nil {
		return nil
	}
	return func() tea.Msg {
		return ViewTickMsg{
			ViewId:          m.Id,
			InternalTickMsg: nextTickCmd(),
			Type:            m.Type,
		}
	}

}

func (m *Model) GetDimensions() constants.Dimensions {
	return constants.Dimensions{
		Width:  m.Ctx.MainContentWidth - containerStyle.GetHorizontalPadding(),
		Height: m.Ctx.MainContentHeight - 2,
	}
}

func (m *Model) UpdateProgramContext(ctx *context.ProgramContext) {
	oldDimensions := m.GetDimensions()
	m.Ctx = ctx
	newDimensions := m.GetDimensions()
	m.Table.SetDimensions(newDimensions)

	if oldDimensions.Height != newDimensions.Height || oldDimensions.Width != newDimensions.Width {
		m.Table.SyncViewPortContent()
	}
}

type ViewMsg interface {
	GetViewId() int
	GetViewType() string
}

type ViewRowsFetchedMsg struct {
	ViewId int
}

func (msg ViewRowsFetchedMsg) GetViewId() int {
	return msg.ViewId
}

type ViewTickMsg struct {
	ViewId          int
	InternalTickMsg tea.Msg
	Type            string
}

func (msg ViewTickMsg) GetViewId() int {
	return msg.ViewId
}

func (msg ViewTickMsg) GetViewType() string {
	return msg.Type
}

func (m *Model) NextRow() int {
	return m.Table.NextItem()
}

func (m *Model) PrevRow() int {
	return m.Table.PrevItem()
}

func (m *Model) FirstItem() int {
	return m.Table.FirstItem()
}

func (m *Model) LastItem() int {
	return m.Table.LastItem()
}

func (m *Model) GetIsLoading() bool {
	return m.IsLoading
}
