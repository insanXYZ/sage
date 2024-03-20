![image](./sage.png)
Sage is a Go library designed to simplify the task of image validation, either from `*multipart.FileHeader` or `*os.File` 

# Installation
```sh
  go get github.com/insanXYZ/sage
```

# Example

#### <b>Image in your pc
```go
import (
    "github.com/insanXYZ/sage"
)

//open file in your pc with os.open
//note : giphy.gif have size 263kb
open, err := os.Open("giphy.gif")
if err != nil {
    fmt.Println(err.Error())
}

//image validation only
err := sage.Validate(open) //not error
if err != nil {
    fmt.Println(err.Error())
}

//validate format 
err := sage.Validate(open, "gif") //not error
if err != nil {
    fmt.Println(err.Error())
}

err = sage.Validate(open, "png") //error
if err != nil {
    fmt.Println(err.Error())
}

//validate size
err := sage.Validate(open, "minsize=100") //not error
if err != nil {
    fmt.Println(err.Error())
}

err = sage.Validate(open, "maxsize=100") //error
if err != nil {
    fmt.Println(err.Error())
}

//validate format and size
err = sage.Validate(open, "gif","minsize=100") //not error
if err != nil {
    fmt.Println(err.Error())
}
```
#### <b>Web Framework
```go
import (
    "github.com/insanXYZ/sage"
)

//Echo (https://echo.labstack.com/)
e := echo.New()

e.POST("/image-upload", func(c echo.Context) error {
    header, _ := c.FormFile("image")
    err := sage.Validate(header)
    ...
})
```
```go
import (
    "github.com/insanXYZ/sage"
)

//Fiber (https://gofiber.io/)
f := fiber.New()

fiber.Post("/image-upload", func(ctx *fiber.Ctx) error {
    file, _ := ctx.FormFile("file")
    err := s.Validate(header)
    ...
})
```

# Tag

| Tag     | Description                             |
|:--------|:----------------------------------------|
| png     | image must be of type png               |
| jpg     | image must be of type jpg               |
| jpeg    | image must be of type jpeg              |
| gif     | image must be of type gif               |
| bmp     | image must be of type bmp               |
| minsize | images must have a certain minimum size |
| maxsize | images must have a certain maximum size |