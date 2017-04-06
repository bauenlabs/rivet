# Rivet ➕
Efficient string concatenation package. (Inspired directly by [this article](http://herman.asia/efficient-string-concatenation-in-go))

## Installation
As with any Go package, rivet can be installed via `go get`:

```bash
go get github.com/bauenlabs/rivet
```

And then import in your application:
```go
import "github.com/bauenlabs/rivet"
```

## Usage
This package exports a method, `Concat`, which takes an unlimited number of strings, concatenates them together, and returns the resulting string.

Example usage:
```go
s := rivet.Concat("a", "b", "c"); // s = abc.
```
