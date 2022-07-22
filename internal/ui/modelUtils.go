package ui

import (
	"github.com/VictorBersy/docker-hub-cli/internal/data"
	"github.com/VictorBersy/docker-hub-cli/internal/ui/components/view"
)

func (m *Model) getCurrView() view.View {
	views := m.getCurrentViews()
	if len(views) == 0 {
		return nil
	}
	return views[m.currViewId]
}

func (m *Model) getCurrRowData() data.RowData {
	view := m.getCurrView()
	if view == nil {
		return nil
	}
	return view.GetCurrRow()
}

func (m *Model) getViewAt(id int) view.View {
	views := m.getCurrentViews()
	if len(views) <= id {
		return nil
	}
	return views[id]
}

func (m *Model) getPrevViewId() int {
	viewsConfigs := m.ctx.GetViewsConfig()
	m.currViewId = (m.currViewId - 1) % len(viewsConfigs)
	if m.currViewId < 0 {
		m.currViewId += len(viewsConfigs)
	}

	return m.currViewId
}

func (m *Model) getNextViewId() int {
	return (m.currViewId + 1) % len(m.ctx.GetViewsConfig())
}
