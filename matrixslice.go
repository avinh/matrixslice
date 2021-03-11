package matrixslice

import (
	"errors"
	"fmt"
)

//Graph struct
type Graph struct {
	BitMatrix []int
}

//New create graph
func New(size int) (Graph, error) {

	if size < 0 {
		return Graph{}, errors.New("dimensions must be non-negative")
	}

	a := make([]int, size)

	return Graph{BitMatrix: a}, nil
}

func (g *Graph) Expansion(count int) {
	if len(g.BitMatrix) < count {
		for i := len(g.BitMatrix); i < count; i++ {
			g.BitMatrix = append(g.BitMatrix, 0)
		}
	}
}

//AddEdge edges
func (g *Graph) AddEdge(v1, v2 uint32) error {
	if v1 == v2 {
		return errors.New("the same index")
	}

	if !g.inRange(v1, v2) {
		return errors.New("index out of range")
	}

	g.BitMatrix[v1] = setBit(g.BitMatrix[v1], uint(v2))
	g.BitMatrix[v2] = setBit(g.BitMatrix[v2], uint(v1))

	return nil
}

//RemoveEdge edges
func (g *Graph) RemoveEdge(v1, v2 uint32) error {
	if !g.inRange(v1, v2) {
		return errors.New("index out of range")
	}

	g.BitMatrix[v1] = clearBit(g.BitMatrix[v1], uint(v2))
	g.BitMatrix[v2] = clearBit(g.BitMatrix[v2], uint(v1))

	return nil
}

func (g *Graph) CheckEdge(index1, index2 int) (bool, error) {
	if !g.inRange(uint32(index1), uint32(index2)) {
		return false, errors.New("index out of range")
	}

	if hasBit(g.BitMatrix[index1], uint(index2)) && hasBit(g.BitMatrix[index2], uint(index1)) {
		return true, nil
	}
	return false, nil
}

func (g *Graph) GetEdges(index int) ([]int, error) {

	if !g.inRangeOne(uint32(index)) {
		return nil, errors.New("index out of range")
	}

	edge := make([]int, 0)

	for i, v := range g.BitMatrix {
		if i == index && v != 0 {
			for j := 0; j < len(g.BitMatrix); j++ {
				check, err := g.CheckEdge(index, j)

				if err != nil {
					return nil, err
				}

				if hasBit(v, uint(j)) && check {
					edge = append(edge, j)
				}
			}
		}
	}
	return edge, nil
}

func (g *Graph) GetRow(index int) (int, error) {
	if !g.inRangeOne(uint32(index)) {
		return 0, errors.New("index out of range")
	}
	return g.BitMatrix[index], nil
}

func (g *Graph) SetRow(index int, row int) error {
	if !g.inRangeOne(uint32(index)) {
		return errors.New("index out of range")
	}
	g.BitMatrix[index] = row
	return nil
}

//PrintMatrix print matrix
func (g *Graph) PrintMatrix() {
	for i, v := range g.BitMatrix {
		fmt.Println(i, v)
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
	return uint32(len(g.BitMatrix))
}
