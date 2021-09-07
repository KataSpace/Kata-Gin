# Kata-Gin

Package for Gin. Automatic generation of route maps

## How to use this package?

+ Declare a struct and implement some gin functions like this:

```go
type Driver struct {
}

func (d *Driver) GetAllName(c *gin.Context) {

c.JSON(http.StatusOK, "GetAllName")
}
```

+ Then register *Driver with `KataGin`:

```go
r = kg.RegisterRouter(r, nil, nil, &Driver{})
r.Run()
```

## Details about `KataGin`
