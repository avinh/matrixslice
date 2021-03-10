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
	g, err := matrixslice.New(70)

	if err != nil {
		fmt.Println("Error")
	}

	g.AddEdge(1, 3)
	g.AddEdge(1, 5)
	g.AddEdge(1, 2)
	g.AddEdge(2, 5)
	g.AddEdge(5, 39)
	g.AddEdge(36, 39)
	g.AddEdge(37, 39)
	g.AddEdge(38, 39)
	g.AddEdge(39, 38)
	g.AddEdge(69, 55)

	//Remove Edge
	g.RemoveEdge(1, 3)
	//Check Edge of two vertices
	g.CheckEdge(39, 39)
	//Get all Edge of a index
	g.GetEdges(1)
	//Get a row of matrix
	g.GetRow(39)
	//Set a row of matrix
	g.SetRow(68, 39)
	//Print Matrix
	g.PrintMatrix()

	//Length of matrix
	fmt.Println(g.Dim())
}
```
