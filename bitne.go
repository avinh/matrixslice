package matrixslice

func NewUint64(n int) []uint64 {
	return make([]uint64, (n+63)/64)
}

func GetBit(b []uint64, index int) bool {
	pos := index / 64
	j := uint(index % 64)
	return (b[pos] & (uint64(1) << j)) != 0
}

func SetBit(b []uint64, index int, value bool, scale int) []uint64 {
	if len(b) < 1 {
		b = NewUint64(scale)
	}
	pos := index / 64
	j := uint(index % 64)
	if value {
		b[pos] |= (uint64(1) << j)
	} else {
		b[pos] &= ^(uint64(1) << j)
	}

	return b
}

func Len(b []uint64) int {
	return 64 * len(b)
}
