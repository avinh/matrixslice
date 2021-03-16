package matrixslice

type BitNe struct {
	Length int
}

func newUint64(n int) ([]uint64, BitNe) {
	bit := BitNe{Length: n}
	return make([]uint64, (n+63)/64), bit
}

func (bit *BitNe) GetBit(b []uint64, index int) bool {
	if index < bit.Length {
		for i := 0; i < bit.Length; i++ {
			b = append(b, 0)
		}
	}
	pos := index / 64
	j := uint(index % 64)
	return (b[pos] & (uint64(1) << j)) != 0
}

func (bit *BitNe) SetBit(b []uint64, index int, value bool, scale int) []uint64 {
	if len(b) < 1 {
		b, _ = newUint64(scale)
	} else if index < bit.Length {
		for i := 0; i < bit.Length; i++ {
			b = append(b, 0)
		}
	}

	pos := index / 64
	j := uint(index % 64)
	if value {
		b[pos] |= (uint64(1) << j)
	} else {
		b[pos] &= ^(uint64(1) << j)
	}

	for i := (len(b) - 1); i > 0; i-- {
		if b[i] == 0 {
			b = remove(b, i)
		} else {
			break
		}
	}
	return b
}

func (bit *BitNe) Len(b []uint64) int {
	return 64 * len(b)
}

func remove(s []uint64, i int) []uint64 {
	s[len(s)-1], s[i] = s[i], s[len(s)-1]
	return s[:len(s)-1]
}
