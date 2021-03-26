## Graph Matrix Slice

Sparse matrix representation of unweighted graphs in Go, and splitting to index, helping in limited memory storage for example in Dynamodb

Support relationship up to 10 billion or more

### Installation
With Go installed, package installation is performed using go get.

```
go get -u github.com/avinh/matrixslice
```

### Example:

``` Go
func main() {
	g := matrixslice.New()

	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(0, 88)
	g.AddEdge(0, 99)
	g.AddEdge(0, 10000000000)
	g.AddEdge(2999999999, 34)

	g.AddEdge(0, 3434)
	g.AddEdge(0, 5)
	g.AddEdge(1, 5)
	g.AddEdge(1, 2)
	g.AddEdge(1, 39)
	g.AddEdge(1, 69)
	g.AddEdge(1, 999)
	g.AddEdge(1, 9999)

	g.AddEdge(36, 39)
	g.AddEdge(37, 39)
	g.AddEdge(38, 39)
	g.AddEdge(34, 9999)
	g.AddEdge(45, 1)
	g.AddEdge(768, 9999)
	g.AddEdge(0, 445353534)
	g.AddEdge(45, 123123)
	g.AddEdge(45, 62000000)
	g.AddEdge(45, 2341232)

	g.AddEdge(999, 2341232)
	g.AddEdge(999, 234)
	g.AddEdge(999, 23423)

	// Remove Edge
	err := g.RemoveEdge(0, 10000000000)
	if err != nil {
		fmt.Println(err)
	}

	// Check Edge of two vertices
	fmt.Println(g.CheckEdge(2999999999, 34))

	// Set a row of matrix
	g.SetRow(0, []uint64{0, 38, 1, 34376515584, 53, 4398046511104, 6958648, 4611686018427387904, 156250000, 1})

	row, err := g.GetRow(0)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(row)

	//Get edge from row
	res, err := g.GetEdgesFromRow(row)

	if err != nil {
		panic(err)
	}

	fmt.Println(res)

	// Print Matrix
	// g.PrintMatrix()
}
```
