package efilfoemag

// Sum returns number of bits set to 1.
func Sum(c uint8) uint8 {
	var s uint8
	for i := 0; i < 8; i++ {
		s += c & 1
		c >>= 1
	}
	return s
}
