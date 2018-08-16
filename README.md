<p align="center"><a href="#readme"><img src="https://gh.kaos.st/go-crowd.svg"/></a></p>

<p align="center"><a href="#installation">Installation</a> • <a href="#usage-example">Usage example</a> • <a href="#build-status">Build Status</a> • <a href="#license">License</a></p>

<p align="center">
  <a href="https://godoc.org/pkg.re/essentialkaos/go-crowd.v2"><img src="https://godoc.org/pkg.re/essentialkaos/go-crowd.v2?status.svg"></a>
  <a href="https://goreportcard.com/report/github.com/essentialkaos/go-crowd"><img src="https://goreportcard.com/badge/github.com/essentialkaos/go-crowd"></a>
  <a href="https://travis-ci.org/essentialkaos/go-crowd"><img src="https://travis-ci.org/essentialkaos/go-crowd.svg"></a>
  <a href="https://codebeat.co/projects/github-com-essentialkaos-go-crowd-master"><img alt="codebeat badge" src="https://codebeat.co/badges/9aaa0412-47a5-4555-924e-9c9e1d61a3e4" /></a>
  <a href="https://essentialkaos.com/ekol"><img src="https://gh.kaos.st/ekol.svg"></a>
</p>

`go-crowd` is a Go package for working with [Crowd REST API](https://developer.atlassian.com/server/crowd/crowd-rest-resources/).

Currently, this package support only getting data from API (_i.e., you cannot create or modify data using this package_).

_**Note, that this is beta software, so it's entirely possible that there will be some significant bugs. Please report bugs so that we are aware of the issues.**_

### Installation

Before the initial install allows git to use redirects for [pkg.re](https://github.com/essentialkaos/pkgre) service (_reason why you should do this described [here](https://github.com/essentialkaos/pkgre#git-support)_):

```
git config --global http.https://pkg.re.followRedirects true
```

Make sure you have a working Go 1.9+ workspace (_[instructions](https://golang.org/doc/install)_), then:

````
go get pkg.re/essentialkaos/go-crowd.v2
````

For update to latest stable release, do:

```
go get -u pkg.re/essentialkaos/go-crowd.v2
```

### Usage example

```go
package main

import (
  "fmt"
  "pkg.re/essentialkaos/go-crowd.v2"
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
| `master` (_Stable_) | [![Build Status](https://travis-ci.org/essentialkaos/go-crowd.svg?branch=master)](https://travis-ci.org/essentialkaos/go-crowd) |
| `develop` (_Unstable_) | [![Build Status](https://travis-ci.org/essentialkaos/go-crowd.svg?branch=develop)](https://travis-ci.org/essentialkaos/go-crowd) |

### License

[EKOL](https://essentialkaos.com/ekol)

<p align="center"><a href="https://essentialkaos.com"><img src="https://gh.kaos.st/ekgh.svg"/></a></p>
