package versionmanager

// local representation of project version a user is currently using.
type local struct {
	Author  string
	Name    string
	Version version
}
