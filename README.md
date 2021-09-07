# Kata-Gin

Package for Gin. Automatic generation of route maps

## How to use this package?

+ Declare a struct and implement some gin functions like this:

```go
package main

import (
	"net/http"

	kg "github.com/KataSpace/Kata-Gin"
	"github.com/gin-gonic/gin"
)

type example struct{}

func (e *example) GetAllName(c *gin.Context) {
	c.JSON(http.StatusOK, "GetAllName")
}

func main() {
	r := gin.Default()
	r = kg.RegisterRouter(r, nil, nil, new(example))

	r.Run()
}

```

## Details about `KataGin`
