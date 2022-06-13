# versionmanager.go
A GitHub Project Version Manager that polls latest version data from GitHub repo release tag.

# Installation
### Go get
`> go get -u github.com/NeonixRIT/versionmanager.go`

# Usage
versionmanager checks the version passed to its constructor against the tag attached to the latest release on a GitHub repo. Comparing the given version string against the `tag_name` value at [https://api.github.com/repos/{author}/{projectName}/releases/latest](). This means release tag names need to be formatted specifically for this. Versionmanager doesnt support letters in version categories* and assumes a separator of a period unless told otherwise.

*A version category is a set of numbers separated by a uniform character (e.g. 2.0.3 has categories 2, 0, and 3). Using Semantic Versioning there are usually 3 version categories (major, minor, and patch) but versionmanager supports more categories as well. 
```
package main

import (
	"fmt"

	vmg "github.com/NeonixRIT/versionmanager.go"
)

func main() {
	vm := vmg.MakeVersionManager("Aquatic-Labs", "Umbra-Mod-Menu", "2.0.4")

	vm.RegisterObserver(func(status vmg.Status, data vmg.Release) {
		if status == vmg.OUTDATED {
			fmt.Println("Outdated.")
		}
	})

	vm.RegisterObserver(func(status vmg.Status, data vmg.Release) {
		if status == vmg.CURRENT {
			fmt.Println("Current.")
		}
	})

	vm.RegisterObserver(func(status vmg.Status, data vmg.Release) {
		if status == vmg.DEV {
			fmt.Println("Dev.")
		}
	})

	vm.CheckStatus()
}

``` 