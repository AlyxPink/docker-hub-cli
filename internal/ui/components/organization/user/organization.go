package organization_user

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	data_user "github.com/victorbersy/docker-hub-cli/internal/data/user"
	"github.com/victorbersy/docker-hub-cli/internal/ui/components/table"
	"github.com/victorbersy/docker-hub-cli/internal/utils"
)

type Organization struct {
	Data data_user.Organization
}

func (org Organization) ToTableRow() table.Row {
	return table.Row{
		org.renderName(),
		org.renderBadge(),
		org.renderCreatedAt(),
	}
}

func (org Organization) renderName() string {
	return lipgloss.NewStyle().
		Render(fmt.Sprint(org.Data.Name))
}

func (org Organization) renderBadge() string {
	return lipgloss.NewStyle().
		Render(fmt.Sprint(org.Data.Badge))
}

func (org Organization) renderCreatedAt() string {
	return lipgloss.NewStyle().
		Render(utils.TimeElapsed(org.Data.DateJoined))
}
