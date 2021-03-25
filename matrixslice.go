package matrixslice

import (
	"errors"
	"fmt"
)

const rangene uint64 = 64000000

type Graph struct {
	BitMatrix [][][]uint64
}

func New() Graph {
	a := make([][][]uint64, rangene)
	return Graph{BitMatrix: a}
}

func (g *Graph) AddEdge(v1, v2 uint64) error {
	if v1 == v2 {
		return errors.New("the same index")
	}

	pos1 := v1 / rangene
	pos2 := v2 / rangene

	v2s := v2
	if pos2 > 0 {
		v2s = uint64(v2 % rangene)
	}

	if len(g.BitMatrix[pos2]) < 1 {
		g.BitMatrix[pos2] = make([][]uint64, rangene)
	}

	v1s := v1
	if pos1 > 0 {
		v1s = uint64(v1 % rangene)
	}

	if len(g.BitMatrix[pos1]) < 1 {
		g.BitMatrix[pos1] = make([][]uint64, rangene)
	}

	bit1 := setBit(g.BitMatrix[pos1][v1s], v2, true)
	bit2 := setBit(g.BitMatrix[pos2][v2s], v1, true)

	g.BitMatrix[pos1][v1s] = bit1
	g.BitMatrix[pos2][v2s] = bit2
	return nil
}

func (g *Graph) RemoveEdge(v1, v2 uint64) error {

	pos1 := v1 / rangene
	pos2 := v2 / rangene

	v1s := v1
	if pos1 > 0 {
		v1s = uint64(v1 % rangene)
	}

	v2s := v2
	if pos2 > 0 {
		v2s = uint64(v2 % rangene)
	}

	if uint64(len(g.BitMatrix[pos2])) < v2s+1 || uint64(len(g.BitMatrix[pos1])) < v1s+1 {
		return errors.New("index out of range")
	}

	if len(g.BitMatrix[pos1][v1s]) < 1 || len(g.BitMatrix[pos2][v2s]) < 1 {
		return errors.New("index out of range")
	}

	bit1 := setBit(g.BitMatrix[pos1][v1s], v2, false)
	bit2 := setBit(g.BitMatrix[pos2][v2s], v1, false)

	g.BitMatrix[pos1][v1s] = bit1
	g.BitMatrix[pos2][v2s] = bit2
	return nil
}

func (g *Graph) CheckEdge(v1, v2 uint64) (bool, error) {
	pos1 := v1 / rangene
	pos2 := v2 / rangene

	v2s := v2
	if pos2 > 0 {
		v2s = uint64(v2 % rangene)
	}

	v1s := v1
	if pos1 > 0 {
		v1s = uint64(v1 % rangene)
	}

	if uint64(len(g.BitMatrix[pos2])) < v2s+1 || uint64(len(g.BitMatrix[pos1])) < v1s+1 {
		return false, nil
	}

	if len(g.BitMatrix[pos1][v1s]) < 1 || len(g.BitMatrix[pos2][v2s]) < 1 {
		return false, nil
	}

	if getBit(g.BitMatrix[pos1][v1s], v2) && getBit(g.BitMatrix[pos2][v2s], v1) {
		return true, nil
	}
	return false, nil
}

func (g *Graph) GetEdges(index uint64) ([]uint64, error) {
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
	return nil, nil
}

func (g *Graph) GetEdgesFromRow(row []uint64) ([]uint64, error) {
	row = unresize(row)
	return scanBit(row), nil
}

func (g *Graph) CountRow(row []uint64) int {
	row = unresize(row)
	return len(scanBit(row))
}

func (g *Graph) GetRow(index uint64) ([]uint64, error) {
	pos := index / rangene

	if len(g.BitMatrix[pos]) < 1 {
		return []uint64{}, nil
	}

	if pos > 0 {
		index = uint64(index % rangene)
	}

	g.BitMatrix[pos][index] = resize(g.BitMatrix[pos][index])
	return g.BitMatrix[pos][index], nil
}

func (g *Graph) SetRow(v uint64, row []uint64) error {
	pos := v / rangene

	if len(g.BitMatrix[pos]) < 1 {
		g.BitMatrix[pos] = make([][]uint64, rangene)
	}

	vs := v
	if pos > 0 {
		vs = uint64(v % rangene)
	}

	g.BitMatrix[pos][vs] = unresize(row)
	return nil
}

func (g *Graph) PrintMatrix() {
	for i, v := range g.BitMatrix {
		fmt.Println(i, v)
	}
}
