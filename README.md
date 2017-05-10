# IoC Container for Go

[![Build Status](https://travis-ci.org/bt/btioc.svg)](https://travis-ci.org/bt/btioc)

Go doesn't currently have a well documented or easy to use IoC container. As I needed one for the project I'm working on, I'm trying to make one that closely resembles [Laravel's IoC Container](https://laravel.com/docs/5.3/container), due to my strong PHP background.

## Getting Started

```go
// Instantiate an object.
obj := "foobar"

// Register into the container.
btioc.Register("foo_object", obj)

// Retrieve the object.
res, _ := btioc.Retrieve("foo_object")

// Prints "what the foobar?!".
fmt.Printf("What the %s?!", res)
```

For a more detailed example (or how you can use structs with the container), check out the [example file](https://github.com/bt/btioc/blob/master/example/main.go).