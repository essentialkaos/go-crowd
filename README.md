<p align="center"><a href="#readme"><img src="https://gh.kaos.st/go-crowd.svg"/></a></p>

<p align="center">
  <a href="https://pkg.re/essentialkaos/go-crowd.v3?docs"><img src="https://gh.kaos.st/godoc.svg" alt="PkgGoDev"></a>
  <a href="https://goreportcard.com/report/github.com/essentialkaos/go-crowd"><img src="https://goreportcard.com/badge/github.com/essentialkaos/go-crowd"></a>
  <a href="https://github.com/essentialkaos/go-crowd/actions"><img src="https://github.com/essentialkaos/go-crowd/workflows/CI/badge.svg" alt="GitHub Actions Status" /></a>
  <a href="https://github.com/essentialkaos/go-crowd/actions?query=workflow%3ACodeQL"><img src="https://github.com/essentialkaos/go-crowd/workflows/CodeQL/badge.svg" /></a>
  <a href="https://codebeat.co/projects/github-com-essentialkaos-go-crowd-master"><img alt="codebeat badge" src="https://codebeat.co/badges/9aaa0412-47a5-4555-924e-9c9e1d61a3e4" /></a>
  <a href="#license"><img src="https://gh.kaos.st/apache2.svg"></a>
</p>

<p align="center"><a href="#installation">Installation</a> • <a href="#usage-example">Usage example</a> • <a href="#build-status">Build Status</a> • <a href="#license">License</a></p>

<br/>

`go-crowd` is a Go package for working with [Crowd REST API](https://developer.atlassian.com/server/crowd/crowd-rest-resources/).

_**Note, that this is beta software, so it's entirely possible that there will be some significant bugs. Please report bugs so that we are aware of the issues.**_

### Installation

Make sure you have a working Go 1.14+ workspace (_[instructions](https://golang.org/doc/install)_), then:

````
go get pkg.re/essentialkaos/go-crowd.v3
````

For update to latest stable release, do:

```
go get -u pkg.re/essentialkaos/go-crowd.v3
```

### Usage example

```go
package main

import (
  "fmt"
  "pkg.re/essentialkaos/go-crowd.v3"
)

func main() {
  api, err := crowd.NewAPI("https://crowd.domain.com/crowd/", "myapp", "MySuppaPAssWOrd")
  api.SetUserAgent("MyApp", "1.2.3")

  if err != nil {
    fmt.Printf("Error: %v\n", err)
    return
  }

  user, err := api.GetUser("john", true)

  if err != nil {
    fmt.Printf("Error: %v\n", err)
    return
  }

  fmt.Println("%-v\n", user)
}
```

### Build Status

| Branch     | Status |
|------------|--------|
| `master` (_Stable_) | [![CI](https://github.com/essentialkaos/go-crowd/workflows/CI/badge.svg?branch=master)](https://github.com/essentialkaos/go-crowd/actions) |
| `develop` (_Unstable_) | [![CI](https://github.com/essentialkaos/go-crowd/workflows/CI/badge.svg?branch=develop)](https://github.com/essentialkaos/go-crowd/actions) |

### License

[Apache License, Version 2.0](http://www.apache.org/licenses/LICENSE-2.0)

<p align="center"><a href="https://essentialkaos.com"><img src="https://gh.kaos.st/ekgh.svg"/></a></p>
