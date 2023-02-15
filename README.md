# Golang gin extension

Depends on https://github.com/gin-gonic/gin

## Features

- middleware implementations
- ...

## Install

`go get -u github.com/patsnapops/ginx`

## Example

```go
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/patsnapops/ginx/middleware"
	"github.com/patsnapops/http-headers"
)

func main() {
	// new default gin engine
	ginEngine := gin.Default()

	// attach middlewares to gin engine, that's it
	middleware.AttachTo(ginEngine).
		WithCacheDisabled().
		WithCORS().
		WithRecover().
		WithRequestID(hh.XRequestID).
		WithSecurity()
}
```
