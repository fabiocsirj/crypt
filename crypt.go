package crypt

func x(d, k1, k2 uint64) uint64 {
	if k1 == 0 && k2 == 0 {
		return 1
	}
	return k2 - d*k1
}
