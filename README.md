# Package matrixslice

The `matrixslice` package provides functionalities to manage a graph represented using a bit matrix structure, optimized for sparse graphs. It includes methods for adding and removing edges, checking edge existence, retrieving edges from a vertex, and manipulating rows of the graph matrix.

## Usage

### Graph Structure

The `Graph` struct represents the graph using a 3-dimensional slice of `uint64` arrays (`[][][]uint64`). This structure is designed to efficiently store and manipulate adjacency information.

### Constants

- `rangene`: A constant representing the range used for partitioning the graph matrix.

### Functions and Methods

#### Graph Struct

- **New() Graph**: Initializes a new instance of `Graph`.

- **AddEdge(v1, v2 uint64) error**: Adds an edge between vertices `v1` and `v2`.

- **RemoveEdge(v1, v2 uint64) error**: Removes an edge between vertices `v1` and `v2`.

- **CheckEdge(v1, v2 uint64) (bool, error)**: Checks if an edge exists between vertices `v1` and `v2`.

- **GetEdges(index uint64) ([]uint64, error)**: Retrieves edges connected to a vertex specified by `index`.

- **GetEdgesFromRow(row []uint64) ([]uint64, error)**: Retrieves edges from a given row of the graph matrix.

- **CountRow(row []uint64) int**: Counts the number of edges in a given row of the graph matrix.

- **GetRow(index uint64) ([]uint64, error)**: Retrieves a specific row from the graph matrix.

- **SetRow(v uint64, row []uint64) error**: Sets a row in the graph matrix.

- **PrintMatrix()**: Prints the current state of the graph matrix.

#### Helper Functions

- **newUint64(n uint64) []uint64**: Creates a new `uint64` slice of specified length.

- **getBit(b []uint64, index uint64) bool**: Retrieves the bit value at a specific index in the `uint64` slice.

- **scanBit(b []uint64) []uint64**: Scans a `uint64` slice for set bits and returns their indices.

- **setBit(b []uint64, index uint64, value bool) []uint64**: Sets or clears a bit at a specific index in the `uint64` slice.

- **bitLen(b []uint64) int**: Calculates the total number of bits in the `uint64` slice.

- **lenUint64(b []uint64) uint64**: Returns the length of a `uint64` slice.

- **remove(s []uint64, i int) []uint64**: Removes an element at index `i` from the `uint64` slice `s`.

- **resize(list []uint64) []uint64**: Converts a list of indices and values into a sparse representation.

- **unresize(list []uint64) []uint64**: Converts a sparse representation back into indices and values.

## Installation

To use this package, ensure you have Go installed and set up properly. You can then install the package using:

```sh
go get github.com/avinh/matrixslice
```

## Example

```go
package main

import (
	"fmt"
	"github.com/[your_username]/matrixslice"
)

func main() {
	graph := matrixslice.New()

	err := graph.AddEdge(1, 2)
	if err != nil {
		fmt.Println("Error adding edge:", err)
	}

	edges, err := graph.GetEdges(1)
	if err != nil {
		fmt.Println("Error retrieving edges:", err)
	} else {
		fmt.Println("Edges from vertex 1:", edges)
	}
}
```

This example initializes a graph, adds an edge between vertices 1 and 2, and retrieves edges from vertex 1.

## License
This package is licensed under the MIT License. See the LICENSE file for more information.


