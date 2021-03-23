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
func (g *Graph) AddEdge(v1, v2 uint32) error {

	if v1 == v2 {
		return errors.New("the same index")
	}

	if !g.inRange(v1, v2) {
		return errors.New("index out of range")
	}

	bit := BitNe{Length: len(g.BitMatrix)}

	g.BitMatrix[v1] = bit.setBit(g.BitMatrix[v1], int(v2), true, len(g.BitMatrix))

	g.BitMatrix[v2] = bit.setBit(g.BitMatrix[v2], int(v1), true, len(g.BitMatrix))

	return nil
}

//RemoveEdge edges
func (g *Graph) RemoveEdge(v1, v2 uint32) error {
	if !g.inRange(v1, v2) {
		return errors.New("index out of range")
	}
	bit := BitNe{Length: len(g.BitMatrix)}
	g.BitMatrix[v1] = bit.setBit(g.BitMatrix[v1], int(v2), false, len(g.BitMatrix))
	g.BitMatrix[v2] = bit.setBit(g.BitMatrix[v2], int(v1), false, len(g.BitMatrix))
	return nil
}

func (g *Graph) CheckEdge(index1, index2 int) (bool, error) {
	if !g.inRange(uint32(index1), uint32(index2)) {
		return false, errors.New("index out of range")
	}
	if len(g.BitMatrix[index1]) < 1 || len(g.BitMatrix[index2]) < 1 {
		return false, nil
	}
	bit := BitNe{Length: len(g.BitMatrix)}
	if bit.getBit(g.BitMatrix[index1], int(index2)) && bit.getBit(g.BitMatrix[index2], int(index1)) {
		return true, nil
	}
	return false, nil
}

func (g *Graph) GetEdges(index int) ([]int, error) {
	if !g.inRangeOne(uint32(index)) {
		return nil, errors.New("index out of range")
	}
	row, err := g.GetRow(uint32(index))
	if err != nil {
		return nil, err
	}

	edge, err := g.GetEdgesFromRow(row)

	if err != nil {
		return nil, err
	}

	results := make([]int, 0)
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

func (g *Graph) GetEdgesFromRow(row []uint64) ([]int, error) {
	bit := BitNe{Length: len(g.BitMatrix)}
	row = uncompress(row)
	return bit.scanBit(row), nil
}

func (g *Graph) CountRow(row []uint64) int {
	bit := BitNe{Length: len(g.BitMatrix)}
	row = uncompress(row)
	return len(bit.scanBit(row))
}

func (g *Graph) GetRow(index uint32) ([]uint64, error) {
	if !g.inRangeOne(uint32(index)) {
		return nil, errors.New("index out of range")
	}
	return compress(g.BitMatrix[index]), nil
}

func (g *Graph) SetRow(index uint32, row []uint64) error {
	if !g.inRangeOne(uint32(index)) {
		return errors.New("index out of range")
	}
	row = compress(row)
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
