# Auth0 Scopes
Consume Auth0 Scopes in Go
A [Go](http://golang.org) middleware for [gin](https://gin-gonic.com/).

[![License MIT][License-Image]][License-Url] 

[License-Url]: https://opensource.org/license/mit/
[License-Image]: https://img.shields.io/badge/License-mit-blue.svg

**Check out [Examples](https://github.com/colwill/go-auth0-scopes/tree/main/examples) - for an idea on how to use this with Gin**

## Installation

```bash
# Go package install
go get github.com/colwill/go-auth0-scopes/
```

## Basic Usage

```go
import (
    middleware "github.com/colwill/go-auth0-scopes"
)

var router = gin.Default()

func main() {

    router.Use(middleware.ScopeUnwrap())
    router.GET("/protected", middleware.Authorise("read:protected"), func(c *gin.Context) {
        c.IndentedJSON(http.StatusOK, gin.H{"message": "this is a protected route!"})
    })

    err = router.Run()
	if err != nil {
		return
	}
}
```
