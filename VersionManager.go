package versionmanager

// VersionManager is The main struct to compare a local project to the latest GitHub release and perform actions based on the result
type VersionManager struct {
	author    string
	name      string
	separator string
	local     local
	remote    remote
	status    int
	observers []func(status int, data Release)
}

// notifyObservers invokes all observers when check status is called. status should be used to tell the observer if it should do anything or not
func (vm *VersionManager) notifyObservers() {
	for _, o := range vm.observers {
		o(vm.status, vm.remote.Data)
	}
}

// RegisterObserver registers a function to be triggered whenever CheckStatus is called
func (vm *VersionManager) RegisterObserver(observer func(status int, data Release)) {
	vm.observers = append(vm.observers, observer)
}

// RegisterObservers registers multiple functions to be triggered whenever CheckStatus is called
func (vm *VersionManager) RegisterObservers(observers []func(status int, data Release)) {
	vm.observers = append(vm.observers, observers...)
}

// CheckStatus compares local project version to the latest GitHub release version. Trigger events based on comparison result
func (vm *VersionManager) CheckStatus() int {
	vm.remote.refresh()

	if vm.local.Version.compareTo(vm.remote.Version) == 0 {
		vm.status = CURRENT
	} else if vm.local.Version.compareTo(vm.remote.Version) < 0 {
		vm.status = OUTDATED
	} else {
		vm.status = DEV
	}
	vm.notifyObservers()
	return vm.status
}

// MakeVersionManager Instantiates a VersionManager for a project with a default separator of a period
func MakeVersionManager(author string, projectName string, localVersion string) VersionManager {
	sep := "."
	vm := VersionManager{
		author:    author,
		name:      projectName,
		separator: sep,
		local:     local{Author: author, Name: projectName, Version: makeVersion(localVersion, sep)},
		remote:    makeRemote(author, projectName, sep),
		status:    0,
		observers: []func(status int, data Release){},
	}
	return vm
}

// MakeVersionManagerSep Instantiates a VersionManager for a project with a custom version separator
func MakeVersionManagerSep(author string, projectName string, localVersion string, separator string) VersionManager {
	vm := VersionManager{
		author:    author,
		name:      projectName,
		separator: separator,
		local:     local{Author: author, Name: projectName, Version: makeVersion(localVersion, separator)},
		remote:    makeRemote(author, projectName, separator),
		status:    0,
		observers: []func(status int, data Release){},
	}
	return vm
}
