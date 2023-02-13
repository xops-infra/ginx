# Golang gin extension

Depends on https://github.com/gin-gonic/gin

## Features

- middleware implementations
- ...

## Install

`go get -u github.com/sfpprxy/ginx`

## Example

```go
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sfpprxy/ginx/middleware"
	"github.com/sfpprxy/http-headers"
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
