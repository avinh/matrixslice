## Graph Matrix Slice

### Installation
With Go installed, package installation is performed using go get.

```
go get -u github.com/avinh/matrixslice
```

### Example:

``` Go
func main() {
	g, err := matrixslice.New(5)

	if err != nil {
		errors.New("Error creating New()")
	}

	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(1, 2)
	g.AddEdge(2, 0)
	g.AddEdge(2, 3)
	g.AddEdge(3, 1)

	g.RemoveEdge(2, 3)

	g.PrintMatrix()

	row1, err := g.GetRow(3)

	if err != nil {
		panic(err)
	}

	err = g.SetRow(4, row1)

	if err != nil {
		panic(err)
	}

	g.PrintMatrix()

	//merge row to matrix
	g.MergeEdgesToRow(1,[3,4,5,7])

	edge1, err := g.GetEdges(2)
	if err != nil {
		panic(err)
	}
	fmt.Println(edge1)

	edgerow, err := g.GetEdgesWithRow(row1, 1)
	if err != nil {
		panic(err)
	}
	fmt.Println(edgerow)

	check, err := g.CheckEdge(0, 4)

	if err != nil {
		panic(err)
	}

	fmt.Println(check)

}
```
