![image](./sage.png)
Sage is a Go library designed to simplify the task of image validation, either from `*multipart.FileHeader` or `*os.File` 

# Installation
```sh
  go get github.com/insanXYZ/sage
```

# Example

#### <b>Image in your pc
```go
//Sage instance
s := sage.New()

//open file in your pc with os.open
//note : giphy.gif have size 263kb
open, _ := os.Open("giphy.gif")

//image validation only
err := s.Validate(open) //not error
if err != nil {
    fmt.Println(err.Error())
}

//validate format 
err := s.Validate(open, "gif") //not error
if err != nil {
    fmt.Println(err.Error())
}

err = s.Validate(open, "png") //error
if err != nil {
    fmt.Println(err.Error())
}

//validate size
err := s.Validate(open, "minsize=100") //not error
if err != nil {
    fmt.Println(err.Error())
}

err = s.Validate(open, "maxsize=100") //error
if err != nil {
    fmt.Println(err.Error())
}

//validate format and size
err = s.Validate(open, "gif","minsize=100") //not error
if err != nil {
    fmt.Println(err.Error())
}
```
#### <b>Web Framework
```go
//Sage instance
s := sage.New()

//Echo (https://echo.labstack.com/)
e := echo.New()

e.POST("/image-upload", func(c echo.Context) error {
    header, _ := c.FormFile("image")
    err := s.Validate(header)
    ...
})
```
```go
//Sage instance
s := sage.New()

//Fiber (https://gofiber.io/)
f := fiber.New()

fiber.Post("/image-upload", func(ctx *fiber.Ctx) error {
    file, _ := ctx.FormFile("file")
    err := s.Validate(header)
    ...
})
```

# Tag

| Tag       | Description                               |
|:----------|:------------------------------------------|
| png     | image must be of type png               |
| jpg     | image must be of type jpg               |
| jpeg    | image must be of type jpeg              |
| gif     | image must be of type gif               |
| minsize | images must have a certain minimum size |
| maxsize | images must have a certain maximum size |