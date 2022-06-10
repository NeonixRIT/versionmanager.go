package versionmanager

import (
	"strconv"
	"strings"
)

// version represents a version number to easily compare project versions
type version struct {
	String string
	Slice  []string
}

// compareTo compares each version category to tell if local version is less than, equal to, or greater than the latest GitHub release
func (v version) compareTo(o version) int {
	i := 0
	for i < len(v.Slice) && i < len(o.Slice) {
		thisVer, errT := strconv.Atoi(v.Slice[i])
		otherVer, errO := strconv.Atoi(o.Slice[i])
		if errT != nil || errO != nil {
			return 0
		}

		if thisVer < otherVer {
			return OUTDATED
		} else if thisVer > otherVer {
			return DEV
		}
		i++
	}
	return CURRENT
}

// makeVersion creates comparable instance of a project's version number
func makeVersion(versionStr string, separator string) version {
	return version{
		String: versionStr,
		Slice:  strings.Split(versionStr, separator)}
}
