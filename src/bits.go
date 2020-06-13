package efilfoemag

// Sumbit returns number of bits set to 1.
func Sumbit(c uint16) uint16 {
	var s uint16
	for i := 0; i < 9; i++ {
		s += c & 1
		c >>= 1
	}
	return s
}
