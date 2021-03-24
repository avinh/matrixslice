package matrixslice

import (
	"errors"
	"fmt"
)

//Graph struct
type Graph struct {
	BitMatrix [][]uint64
}

//New create graph
func New(size uint64) (Graph, error) {

	if size < 0 {
		return Graph{}, errors.New("dimensions must be non-negative")
	}

	a := make([][]uint64, size)

	return Graph{BitMatrix: a}, nil
}

func (g *Graph) Expansion(count int) {
	if len(g.BitMatrix) < count {
		for i := len(g.BitMatrix); i < count; i++ {
			a := make([]uint64, 0)
			g.BitMatrix = append(g.BitMatrix, a)
		}
	}
}

//AddEdge edges
func (g *Graph) AddEdge(v1, v2 uint64) error {

	if v1 == v2 {
		return errors.New("the same index")
	}

	if !g.inRange(v1, v2) {
		return errors.New("index out of range")
	}

	g.BitMatrix[v1] = setBit(g.BitMatrix[v1], v2, true)

	g.BitMatrix[v2] = setBit(g.BitMatrix[v2], v1, true)

	return nil
}

//RemoveEdge edges
func (g *Graph) RemoveEdge(v1, v2 uint64) error {
	if !g.inRange(v1, v2) {
		return errors.New("index out of range")
	}

	g.BitMatrix[v1] = setBit(g.BitMatrix[v1], v2, false)

	g.BitMatrix[v2] = setBit(g.BitMatrix[v2], v1, false)

	return nil
}

func (g *Graph) CheckEdge(index1, index2 uint64) (bool, error) {
	if !g.inRange(uint64(index1), uint64(index2)) {
		return false, errors.New("index out of range")
	}
	if len(g.BitMatrix[index1]) < 1 || len(g.BitMatrix[index2]) < 1 {
		return false, nil
	}

	if getBit(g.BitMatrix[index1], index2) && getBit(g.BitMatrix[index2], index1) {
		return true, nil
	}
	return false, nil
}

func (g *Graph) GetEdges(index uint64) ([]uint64, error) {
	if !g.inRangeOne(uint64(index)) {
		return nil, errors.New("index out of range")
	}
	row, err := g.GetRow(uint64(index))
	if err != nil {
		return nil, err
	}

	edge, err := g.GetEdgesFromRow(row)

	if err != nil {
		return nil, err
	}

	results := make([]uint64, 0)
	for _, v := range edge {
		check, err := g.CheckEdge(v, index)
		if err != nil {
			return nil, err
		}
		if check {
			results = append(results, v)
		}
	}
	return edge, nil
}

func (g *Graph) GetEdgesFromRow(row []uint64) ([]uint64, error) {
	return scanBit(row), nil
}

func (g *Graph) CountRow(row []uint64) int {
	return len(scanBit(row))
}

func (g *Graph) GetRow(index uint64) ([]uint64, error) {
	if !g.inRangeOne(uint64(index)) {
		return nil, errors.New("index out of range")
	}
	g.BitMatrix[index] = resize(g.BitMatrix[index])
	return g.BitMatrix[index], nil
}

func (g *Graph) SetRow(index uint64, row []uint64) error {
	if !g.inRangeOne(uint64(index)) {
		return errors.New("index out of range")
	}
	g.BitMatrix[index] = unresize(row)
	return nil
}

//PrintMatrix print matrix
func (g *Graph) PrintMatrix() {
	for i, v := range g.BitMatrix {
		fmt.Println(i, v)
	}
}

// inRange returns true if (r, c) is a valid index into v.
func (g *Graph) inRange(r, c uint64) bool {
	n := g.Dim()
	return (c < n) && (r < n)
}

func (g *Graph) inRangeOne(r uint64) bool {
	n := g.Dim()
	return (r < n)
}

// Dim returns the (single-axis) dimension of the Graph.
func (g *Graph) Dim() uint64 {
	return uint64(len(g.BitMatrix))
}
