# 09_modules_packages_imports

### Intro
* repositories
* modules
* packages
* 1 or more packages per module
* 1 module per 1 repository

### Modules
* `go.mod` file is necessary (and needs to be correct of course)
* `go.mod` file contains:
    - module name
    - go version
    - dependencies
    - version of dependencies
    - (optional) replace 
    - (optional) exclude
* `pkg` (for actual code) and `cmd` (for executables)
* [modules repository](https://pkg.go.dev/), kinda like `pypi`, but not really - this is just for reference, while pypi is for distribution
* Semantic Versioning (SemVer) - `v1.2.3` (major.minor.patch)
* sometimes `go.sum` file contains `// indirect` comment, which means there is an older module, without `go.mod` file, `go.mod` file is corrupeted, or similar reason
* non-backward-compatible changes means, practically, new module, thus `github.com/.../v2` (or similar) is used is `go.mod` file
* `vendoring` means copying all project dependencies to the project directory, so nothing is downloaded while compiling, thus the project is self-contained and guaranteed to work
* `vendor` top-level directory is used for vendoring
* [golang modules proxy server](https://proxy.golang.org/) is used for caching and security reasons
* [Athens project](https://docs.gomods.io/) is used for caching and security reasons

### Packages
* `Capitalized` names are exported
* `interal` package is only available to packages directly above it and it's subpackages

### API
* `alias` is different name for `type` (basically)

### Imports
