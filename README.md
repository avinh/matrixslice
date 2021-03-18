## Graph Matrix Slice

Sparse matrix representation of unweighted graphs in Go, and splitting to index, helping in limited memory storage for example in Dynamodb


### Installation
With Go installed, package installation is performed using go get.

```
go get -u github.com/avinh/matrixslice
```

### Example:

``` Go
func main() {
	g, err := matrixslice.New(10000)

	if err != nil {
		fmt.Println("Error")
	}

	g.AddEdge(1, 3)
	g.AddEdge(1, 5)
	g.AddEdge(1, 2)
	g.AddEdge(2, 5)
	g.AddEdge(1, 39)
	g.AddEdge(36, 39)
	g.AddEdge(37, 39)
	g.AddEdge(38, 39)
	g.AddEdge(1, 69)
	g.AddEdge(1, 999)
	g.AddEdge(1, 9999)

	//Remove Edge
	g.RemoveEdge(1, 3)
	// Check Edge of two vertices
	fmt.Println(g.CheckEdge(36, 39))
	//Get all Edge of a index
	fmt.Println(g.GetEdges(1))
	//Set a row of matrix
	g.SetRow(9999, []uint64{1})
	//scale matrix
	g.Expansion(100)
	//Print Matrix
	g.PrintMatrix()
	//Length of matrix
	fmt.Println(g.Dim())


	row, err := g.GetRow(1)
	if err != nil {
		fmt.Println(err)
	}
	
	//Get Count row
	fmt.Println(g.CountRow(row))
	fmt.Println(g.GetEdgesFromRow(row))
}
```
