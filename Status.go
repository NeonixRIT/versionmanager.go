package versionmanagergo

// Status variables
type Status int

const (
	OUTDATED Status = -1
	CURRENT  Status = 0
	DEV      Status = 1
)

func (s Status) Base() int {
	return int(s)
}
