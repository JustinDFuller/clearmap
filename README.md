# clearmap

[![Go Reference](https://pkg.go.dev/badge/github.com/justindfuller/clearmap.svg)](https://pkg.go.dev/github.com/justindfuller/clearmap)
[![Go Report Card](https://goreportcard.com/badge/github.com/justindfuller/clearmap)](https://goreportcard.com/report/github.com/justindfuller/clearmap)

The `clearmap` package clears maps.

## Install

Run `go get github.com/justindfuller/clearmap` to add it to your package.

## Usage

Use `clearmap.Clear` to clear a map of all [zero](https://go.dev/tour/basics/12), [empty](https://go.dev/tour/moretypes/11), and [nil](https://go.dev/tour/moretypes/12) values, including:

* Zero Numbers (float, int, complex)
* Empty Strings 
* Zero Runes
* Zero Bytes
* False Booleans 
* Empty Arrays
* Empty Maps
* Nil values

```go
func main() {
  myMap := map[string]interface{}{
    "int": 0,
    "string": "",
    "bool": false,
  }

  fmt.Println(clearmap.Clear(myMap)) // nil
  fmt.Println(myMap) // map[string]interface{}{}
}
``` 

Here's how it works:

* If an int, int32, or int64 is `0`, the key is removed.
* If a string is empty, the key is removed.
* If a boolean is false, the key is removed.
* If an array is empty or nil, the key is removed.
* If a map is empty or nil, the key is removed.
* If any other value is nil, the key is removed.
* If the entire map is empty, nil is returned.

## Tests

You can run the tests with `make test` or `make test-watch`.

The `test-watch` command requires [reflex](https://github.com/cespare/reflex).
