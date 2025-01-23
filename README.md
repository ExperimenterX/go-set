# go-set

[![Go Reference](https://pkg.go.dev/badge/github.com/ExperimenterX/go-set.svg)](https://pkg.go.dev/github.com/ExperimenterX/go-set)

A simple and efficient **Set** implementation for Go, designed for fast lookups, insertions, and deletions. Ideal for handling unique elements with ease.

## Installation
```sh
go get github.com/ExperimenterX/go-set
```

## Usage
```go
package main

import (
	"fmt"
	"github.com/ExperimenterX/go-set"
)

func main() {
	// Create a new set
	s := set.New[int]()

	// Add elements
	s.Add(1, 2, 3)

	// Check existence
	fmt.Println(s.Contains(2)) // true
}
```

## Reference
- **Download**: `go get github.com/ExperimenterX/go-set`
- **Usage**: See [pkg.go.dev/github.com/ExperimenterX/go-set](https://pkg.go.dev/github.com/ExperimenterX/go-set)

## License
This project is licensed under the **MIT License**.

---

Start using `go-set` today and make working with unique collections in Go easier! 🚀
