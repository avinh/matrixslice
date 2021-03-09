package matrixslice

import (
	"errors"
	"fmt"
)

type Graph struct {
	AdjMatrix [][]int
}

//New create graph
func New(size int) (Graph, error) {

	if size < 0 {
		return Graph{}, errors.New("dimensions must be non-negative")
	}

	a := make([][]int, size)
	for i := range a {
		a[i] = make([]int, size)
	}

	return Graph{AdjMatrix: a}, nil
}

//AddEdge edges
func (g *Graph) AddEdge(v1, v2 uint32) error {

	if !g.inRange(v1, v2) {
		return errors.New("index out of range")
	}
	if len(g.AdjMatrix[v1]) < 1 || len(g.AdjMatrix[v2]) < 1 {
		return errors.New("Matrix is empty")
	}

	g.AdjMatrix[v1][v2] = 1
	g.AdjMatrix[v2][v1] = 1

	return nil
}

//RemoveEdge edges
func (g *Graph) RemoveEdge(v1, v2 uint32) error {
	if !g.inRange(v1, v2) {
		return errors.New("index out of range")
	}
	if g.AdjMatrix[v1][v2] == 0 {
		return nil
	}
	g.AdjMatrix[v1][v2] = 0
	g.AdjMatrix[v2][v1] = 0

	return nil
}

func (g *Graph) CheckEdge(index1, index2 int) (bool, error) {
	if !g.inRange(uint32(index1), uint32(index2)) {
		return false, errors.New("index out of range")
	}

	a := g.AdjMatrix

	if a[index1][index2] == 1 && a[index2][index1] == 1 {
		return true, nil
	}
	return false, nil
}

func (g *Graph) GetEdges(index int) ([]int, error) {
	a := g.AdjMatrix

	if !g.inRangeOne(uint32(index)) {
		return nil, errors.New("index out of range")
	}

	edge := make([]int, 0)

	for i, v := range a[index] {
		if v == 1 {
			edge = append(edge, int(i))
		}
	}
	return edge, nil
}

func (g *Graph) MergeEdgesToRow(index int, row []int) ([]int, error) {
	if len(row) < 1 {
		return nil, errors.New("Row is emmty")
	}
	s := make([]int, len(g.AdjMatrix))
	for _, v := range row {
		s[v] = 1
	}
	err := g.SetRow(index, s)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (g *Graph) GetEdgesWithRow(row []int, index int) ([]int, error) {
	a := g.AdjMatrix

	if len(row) < 1 {
		return nil, errors.New("Row is emmty")
	}

	if !g.inRangeOne(uint32(index)) {
		return nil, errors.New("index out of range")
	}

	if len(row) < len(a[index]) {
		for i := len(row); i <= len(a[index]); i++ {
			row = append(row, 0)
		}
	}

	edge := make([]int, 0)

	for i, v := range row {
		if v == 1 {
			edge = append(edge, int(i))
		}
	}

	return edge, nil
}

func (g *Graph) GetRow(index int) ([]int, error) {
	if !g.inRangeOne(uint32(index)) {
		return nil, errors.New("index out of range")
	}

	a := g.AdjMatrix
	s := make([]int, 0)
	for i, v := range a {
		if index == i {
			s = v
		}
	}

	return s, nil
}

func (g *Graph) SetRow(index int, row []int) error {
	if !g.inRangeOne(uint32(index)) {
		return errors.New("index out of range")
	}
	g.AdjMatrix[index] = row
	return nil
}

//PrintMatrix print matrix
func (g *Graph) PrintMatrix() {
	for i, v := range g.AdjMatrix {
		fmt.Println(i)
		fmt.Println(v)
	}
}

// inRange returns true if (r, c) is a valid index into v.
func (g *Graph) inRange(r, c uint32) bool {
	n := g.Dim()
	return (c < n) && (r < n)
}

func (g *Graph) inRangeOne(r uint32) bool {
	n := g.Dim()
	return (r < n)
}

// Dim returns the (single-axis) dimension of the Graph.
func (g *Graph) Dim() uint32 {
	return uint32(len(g.AdjMatrix))
}
