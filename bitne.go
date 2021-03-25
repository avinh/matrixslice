package matrixslice

func newUint64(n uint64) []uint64 {
	return make([]uint64, (n+63)/64)
}

func getBit(b []uint64, index uint64) bool {
	pos := index / 64
	j := uint64(index % 64)
	return (b[pos] & (uint64(1) << j)) != 0
}

func scanBit(b []uint64) []uint64 {
	list := make([]uint64, 0)
	for i, v := range b {
		if v != 0 {
			for j := 0; j < 64; j++ {
				if v&(uint64(1)<<j) != 0 {
					res := j
					if i != 0 {
						res = 64*i + j
					}
					list = append(list, uint64(res))
				}
			}
		}
	}
	return list
}

func setBit(b []uint64, index uint64, value bool) []uint64 {
	pos := index / 64
	j := uint(index % 64)

	temp := make([]uint64, (index+64)/64)

	if (len(b)) > len(temp) {
		temp = b
	}

	for i, v := range b {
		if v != 0 && uint64(len(b)) < (index+64)/64 {
			temp[i] = v
		}
	}

	if value {
		temp[pos] |= (uint64(1) << j)
	} else {
		temp[pos] &= ^(uint64(1) << j)
	}

	return temp
}

func bitLen(b []uint64) int {
	return int(64 * len(b))
}

func lenUint64(b []uint64) uint64 {
	return uint64(len(b))
}

func remove(s []uint64, i int) []uint64 {
	s[len(s)-1], s[i] = s[i], s[len(s)-1]
	return s[:len(s)-1]
}

func resize(list []uint64) (results []uint64) {
	for i, v := range list {
		if v != 0 {
			results = append(results, uint64(i))
			results = append(results, v)
		}
	}
	return results
}

func unresize(list []uint64) (results []uint64) {
	var max uint64 = 0
	for i := 0; uint64(i) < lenUint64(list); i++ {
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
