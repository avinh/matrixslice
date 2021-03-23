package matrixslice

type BitNe struct {
	Length int
}

func newUint64(n int) ([]uint64, BitNe) {
	bit := BitNe{Length: (n + 63) / 64}
	return make([]uint64, (n+63)/64), bit
}

func (bit *BitNe) getBit(b []uint64, index int) bool {
	pos := index / 64
	j := uint(index % 64)
	if pos < ((bit.Length + 63) / 64) {
		for i := 0; i < (bit.Length+63)/64; i++ {
			b = append(b, 0)
		}
	}
	return (b[pos] & (uint64(1) << j)) != 0
}

func (bit *BitNe) scanBit(b []uint64) []int {
	list := make([]int, 0)
	for i, v := range b {
		if v != 0 {
			for j := 0; j < 64; j++ {
				if v&(uint64(1)<<j) != 0 {
					res := j
					if i != 0 {
						res = 64*i + j
					}
					list = append(list, res)
				}
			}
		}
	}
	return list
}

func (bit *BitNe) setBit(b []uint64, index int, value bool, scale int) []uint64 {
	pos := index / 64
	j := uint(index % 64)

	if len(b) < 1 {
		b, _ = newUint64(scale)
	} else if pos < (bit.Length+63)/64 {
		for i := 0; i < (bit.Length+63)/64; i++ {
			b = append(b, 0)
		}
	}

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

func (bit *BitNe) len(b []uint64) int {
	return 64 * len(b)
}

func remove(s []uint64, i int) []uint64 {
	s[len(s)-1], s[i] = s[i], s[len(s)-1]
	return s[:len(s)-1]
}

func compress(list []uint64) (results []uint64) {
	for i, v := range list {
		if v != 0 {
			results = append(results, uint64(i))
			results = append(results, v)
		}
	}
	return results
}

func uncompress(list []uint64) (results []uint64) {
	var max uint64 = 0
	for i := 0; i < len(list); i++ {
		if i%2 == 0 && list[i] > max {
			max = list[i]
		}
	}
	resultne := make([]uint64, max+1)
	for i, v := range list {
		if i%2 == 0 {
			resultne[v] = list[i+1]
		}
	}
	return resultne
}
