package organization_user

import (
	data_user "github.com/victorbersy/docker-hub-cli/internal/data/user"
	"github.com/victorbersy/docker-hub-cli/internal/ui/components/table"
)

type Organization struct {
	Data data_user.Organization
}

func (org Organization) ToTableRow() table.Row {
	return table.Row{
		org.renderName(),
	}
}

func (org Organization) renderName() string {
	return org.Data.Name
}
