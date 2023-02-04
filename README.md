# clearmap

The `clearmap` package clears maps.

## Install

Run `go get github.com/justindfuller/clearmap` to add it to your package.

## Usage

To clear a map of all Zero values, including:

* Zeroes
* Empty Strings 
* False Booleans 
* Empty Arrays
* Empty Maps
* Nil values

You can use `clearmap.Clear`.

```go
func main() {
  myMap := map[string]interface{}{
    "int": 0,
    "string": "",
    "bool": false,
  }

  clearmap.Clear(myMap)

  fmt.Println(myMap) // nil
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

