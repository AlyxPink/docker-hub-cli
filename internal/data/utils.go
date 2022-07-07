package data

import "time"

type RowData interface {
	GetRepoNameWithOwner() string
	GetUrl() string
	GetLastUpdate() time.Time
}
