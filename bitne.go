package matrixslice

type BitNe struct {
	Set []uint64
}

func NewUint64(n int) BitNe {
	items := make([]uint64, (n+63)/64)
	return BitNe{Set: items}
}

func GetBit(b []uint64, index int) bool {
	pos := index / 64
	j := uint(index % 64)
	return (b[pos] & (uint64(1) << j)) != 0
}

func SetBit(b BitNe, index int, value bool, scale int) BitNe {
	if len(b.Set) < 1 {
		b = NewUint64(scale)
	}
	pos := index / 64
	j := uint(index % 64)
	if value {
		b.Set[pos] |= (uint64(1) << j)
	} else {
		b.Set[pos] &= ^(uint64(1) << j)
	}

	return b
}

func Len(b []uint64) int {
	return 64 * len(b)
}
