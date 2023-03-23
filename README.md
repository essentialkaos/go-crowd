<p align="center"><a href="#readme"><img src="https://gh.kaos.st/go-crowd.svg"/></a></p>

<p align="center">
  <a href="https://kaos.sh/g/go-crowd.v3"><img src="https://gh.kaos.st/godoc.svg" alt="PkgGoDev" /></a>
  <a href="https://kaos.sh/r/go-crowd"><img src="https://kaos.sh/r/go-crowd.svg" alt="GoReportCard" /></a>
  <a href="https://kaos.sh/w/go-crowd/ci"><img src="https://kaos.sh/w/go-crowd/ci.svg" alt="GitHub Actions CI Status" /></a>
  <a href="https://kaos.sh/w/go-crowd/codeql"><img src="https://kaos.sh/w/go-crowd/codeql.svg" alt="GitHub Actions CodeQL Status" /></a>
  <a href="https://kaos.sh/b/go-crowd"><img src="https://kaos.sh/b/9aaa0412-47a5-4555-924e-9c9e1d61a3e4.svg" alt="Codebeat badge" /></a>
  <a href="#license"><img src="https://gh.kaos.st/apache2.svg"></a>
</p>

<p align="center"><a href="#installation">Installation</a> • <a href="#usage-example">Usage example</a> • <a href="#build-status">Build Status</a> • <a href="#license">License</a></p>

<br/>

`go-crowd` is a Go package for working with [Crowd REST API](https://developer.atlassian.com/server/crowd/crowd-rest-resources/).

_**Note, that this is beta software, so it's entirely possible that there will be some significant bugs. Please report bugs so that we are aware of the issues.**_

### Installation

Make sure you have a working Go 1.18+ workspace (_[instructions](https://golang.org/doc/install)_), then:

```bash
go get -u github.com/essentialkaos/go-crowd/v3
```

### Usage example

```go
package main

import (
  "fmt"
  "github.com/essentialkaos/go-crowd/v3"
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
| `master` (_Stable_) | [![CI](https://kaos.sh/w/go-crowd/ci.svg?branch=master)](https://kaos.sh/w/go-crowd/ci?query=branch:master) |
| `develop` (_Unstable_) | [![CI](https://kaos.sh/w/go-crowd/ci.svg?branch=develop)](https://kaos.sh/w/go-crowd/ci?query=branch:develop) |

### License

[Apache License, Version 2.0](http://www.apache.org/licenses/LICENSE-2.0)

<p align="center"><a href="https://essentialkaos.com"><img src="https://gh.kaos.st/ekgh.svg"/></a></p>
