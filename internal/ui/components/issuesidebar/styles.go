package issuesidebar

import (
	"github.com/docker/hack-docker-access-management-cli/internal/ui/styles"
)

var (
	pillStyle = styles.MainTextStyle.Copy().
		Foreground(styles.DefaultTheme.SubleMainText).
		PaddingLeft(1).
		PaddingRight(1)
)
