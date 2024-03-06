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

### Packages
* `Capitalized` names are exported
* `interal` package is only available to packages directly above it and it's subpackages

### API
* `alias` is different name for `type` (basically)

### Imports
