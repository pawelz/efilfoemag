package neighborhood

import "fmt"

import "github.com/pawelz/efilfoemag/src/bits"

type Neighborhood uint16

// Side is an address of a cell in the neighborhood.
//
// It goes like that:
//
//    NW  N  NE
//
//    W   C   E
//
//    SW  S  SE
type Side int

const (
	NW Side = 8
	N  Side = 7
	NE Side = 6
	W  Side = 5
	C  Side = 4
	E  Side = 3
	SW Side = 2
	S  Side = 1
	SE Side = 0
)

var (
	ancestorsOfAlive []Neighborhood
	ancestorsOfDead []Neighborhood
)

func init() {
	var n Neighborhood
	for n = 0; n < 0x200; n++ {
		isAlive := func() {
			ancestorsOfAlive = append(ancestorsOfAlive, n)
		}
		isDead := func() {
			ancestorsOfDead = append(ancestorsOfDead, n)
		}
		sumbit := bits.Sumbit(uint16(n))
		switch {
		case n.C() && sumbit == 4:
			isAlive()
		case n.C() && sumbit == 3:
			isAlive()
		case !n.C() && sumbit == 3:
			isAlive()
		default:
			isDead()
		}
	}
}

func (s Side) ToStr() string {
	switch s {
	case NW:
		return "NW"
	case N:
		return "N"
	case NE:
		return "NE"
	case W:
		return "W"
	case C:
		return "C"
	case E:
		return "E"
	case SW:
		return "SW"
	case S:
		return "S"
	case SE:
		return "SE"
	}
	return fmt.Sprintf("[invalid side %d]", s)
}

// NW returns true if the NW cell of the neighborhood is alive.
func (n Neighborhood) NW() bool {
	return n&(1<<uint(NW)) != 0
}

// N returns true if the N cell of the neighborhood is alive.
func (n Neighborhood) N() bool {
	return n&(1<<uint(N)) != 0
}

// NE returns true if the NE cell of the neighborhood is alive.
func (n Neighborhood) NE() bool {
	return n&(1<<uint(NE)) != 0
}

// W returns true if the W cell of the neighborhood is alive.
func (n Neighborhood) W() bool {
	return n&(1<<uint(W)) != 0
}

// C returns true if the C cell of the neighborhood is alive.
func (n Neighborhood) C() bool {
	return n&(1<<uint(C)) != 0
}

// E returns true if the E cell of the neighborhood is alive.
func (n Neighborhood) E() bool {
	return n&(1<<uint(E)) != 0
}

// SW returns true if the SW cell of the neighborhood is alive.
func (n Neighborhood) SW() bool {
	return n&(1<<uint(SW)) != 0
}

// S returns true if the S cell of the neighborhood is alive.
func (n Neighborhood) S() bool {
	return n&(1<<uint(S)) != 0
}

// SE returns true if the SE cell of the neighborhood is alive.
func (n Neighborhood) SE() bool {
	return n&(1<<uint(SE)) != 0
}

// Matches checks whether the other neighborhood matches the current at given distance and side.
//
// For example, let:
//
//       OOO        ###
//   n = O##    k = #O#
//       O#O        ###
//
// then it is true that n.Matches(k, 1, SW), and false for any other side and distance.
func (n Neighborhood) Matches(k Neighborhood, dist int, s Side) (bool, error) {
	if s == C {
		return false, fmt.Errorf("C is not a valid side for matching neighborhoods")
	}
	var v bool
	switch dist {
	case 1:
		switch s {
		case NW:
			v = n.NW() == k.C() && n.N() == k.E() && n.W() == k.S() && n.C() == k.SE()
		case N:
			v = n.NW() == k.W() && n.N() == k.C() && n.NE() == k.E() && n.W() == k.SW() && n.C() == k.S() && n.E() == k.SE()
		case NE:
			v = n.N() == k.W() && n.NE() == k.C() && n.C() == k.SW() && n.E() == k.S()
		case W:
			v = n.NW() == k.N() && n.N() == k.NE() && n.W() == k.C() && n.C() == k.E() && n.SW() == k.S() && n.S() == k.SE()
		case E:
			v = n.N() == k.NW() && n.NE() == k.N() && n.C() == k.W() && n.E() == k.C() && n.S() == k.SW() && n.SE() == k.S()
		case SW:
			v = n.W() == k.N() && n.C() == k.NE() && n.SW() == k.C() && n.S() == k.E()
		case S:
			v = n.W() == k.NW() && n.C() == k.N() && n.E() == k.NE() && n.SW() == k.W() && n.S() == k.C() && n.SE() == k.E()
		case SE:
			v = n.C() == k.NW() && n.E() == k.N() && n.S() == k.W() && n.SW() == k.C()
		}
	case 2:
		switch s {
		case NW:
			v = n.NW() == k.SE()
		case N:
			v = n.NW() == k.SW() && n.N() == k.S() && n.NE() == k.SE()
		case NE:
			v = n.NE() == k.SW()
		case W:
			v = n.NW() == k.NE() && n.W() == k.E() && n.SW() == k.SE()
		case E:
			v = n.NE() == k.NW() && n.E() == k.W() && n.SE() == k.SW()
		case SW:
			v = n.SW() == k.NE()
		case S:
			v = n.SW() == k.NW() && n.S() == k.N() && n.SE() == k.NE()
		case SE:
			v = n.SE() == k.NW()
		}
	default:
		return false, fmt.Errorf("want dist equal 1 or 2, got %d", dist)
	}
	return v, nil
}

// Parse is a factory that produces a Neighborhood from a string.
//
// The parser interprets '#' characters as alive cells and '+' as dead cells. It ignores all other
// characters. The input string must contain exactly 9 meaningful characters.
func Parse(n string) (Neighborhood, error) {
	var v Neighborhood
	var sideIterator int
	nextSide := func() (uint, error) {
		if sideIterator >= 9 {
			return 0, fmt.Errorf("")
		}
		sides := []Side{NW, N, NE, W, C, E, SW, S, SE}
		side := sides[sideIterator]
		sideIterator++
		return uint(side), nil
	}
	for _, char := range n {
		switch char {
		case '#':
			shift, err := nextSide()
			if err != nil {
				return 0, fmt.Errorf("invalid Neighborhood, want exactly 9 meaningful characters ('#' or '+'), got: %q", n)
			}
			v |= 1 << shift
		case '+':
			if _, err := nextSide(); err != nil {
				return 0, fmt.Errorf("invalid Neighborhood, want exactly 9 meaningful characters ('#' or '+'), got: %q", n)
			}
		}
	}
	if sideIterator != 9 {
		return 0, fmt.Errorf("invalid Neighborhood, want exactly 9 meaningful characters ('#' or '+'), got: %q", n)
	}
	return v, nil
}
