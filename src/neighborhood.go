// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package neighborhood

import (
	"fmt"

	"github.com/pawelz/efilfoemag/src/bits"
	"github.com/pawelz/efilfoemag/src/state"
)

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

	b10000000 uint8 = 0x80
)

var (
	ancestorsOfAlive []Neighborhood
	ancestorsOfDead  []Neighborhood
	sides            = []Side{NW, N, NE, W, C, E, SW, S, SE}
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
		sumbit := bits.Sum(uint16(n))
		switch {
		case n.C().IsAlive() && sumbit == 4:
			isAlive()
		case n.C().IsAlive() && sumbit == 3:
			isAlive()
		case !n.C().IsAlive() && sumbit == 3:
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

func mask(s Side) uint16 {
	return 0x1ff & ^(1 << uint(s))
}

// NW returns the state of the NW cell of the neighborhood.
func (n Neighborhood) NW() state.State {
	return state.Of(n&(1<<uint(NW)) != 0)
}

// N returns the state of the N cell of the neighborhood.
func (n Neighborhood) N() state.State {
	return state.Of(n&(1<<uint(N)) != 0)
}

// NE returns the state of the NE cell of the neighborhood.
func (n Neighborhood) NE() state.State {
	return state.Of(n&(1<<uint(NE)) != 0)
}

// W returns the state of the W cell of the neighborhood.
func (n Neighborhood) W() state.State {
	return state.Of(n&(1<<uint(W)) != 0)
}

// C returns the state of the C cell of the neighborhood.
func (n Neighborhood) C() state.State {
	return state.Of(n&(1<<uint(C)) != 0)
}

// E returns the state of the E cell of the neighborhood.
func (n Neighborhood) E() state.State {
	return state.Of(n&(1<<uint(E)) != 0)
}

// SW returns the state of the SW cell of the neighborhood.
func (n Neighborhood) SW() state.State {
	return state.Of(n&(1<<uint(SW)) != 0)
}

// S returns the state of the S cell of the neighborhood.
func (n Neighborhood) S() state.State {
	return state.Of(n&(1<<uint(S)) != 0)
}

// SE returns the state of the SE cell of the neighborhood.
func (n Neighborhood) SE() state.State {
	return state.Of(n&(1<<uint(SE)) != 0)
}

// Set sets the given side of the neighborhood to a vaue.
func (n *Neighborhood) Set(side Side, s state.State) {
	var b uint16
	if s.IsAlive() {
		b = 1
	} else {
		b = 0
	}
	*n = Neighborhood((uint16(*n) & mask(side)) | b<<uint(side))
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

// ToStr renders a Neighborhood as a human readable string.
func (n Neighborhood) ToStr() string {
	var v string
	for i, s := range sides {
		v += string(state.Of(n&(1<<uint(s)) != 0).ToRune())
		if i == 2 || i == 5 {
			v += ","
		}
	}
	return v
}

// Set represents a set of Neighborhoods.
//
// There are 2^9 (= 8*64) possible Neighborhood. So we represent the set as an
// array of 64 bytes. The bit at index n is set if the Neighborhood n belongs to
// the set.
//
// TODO(pawelz@execve.pl): since we are overoptimizing anyway, would this be
//                         more efficient if it used uint32 or even uint64?
type Set [64]uint8

// Returns the address of the neighborhood in the Set.
func setAddress(n Neighborhood) (uint8, uint8, error) {
	if n > 0x1ff {
		return 0xff, 0xff, fmt.Errorf("Illegal Neighborhood. Want integer in 0 .. 0x1ff, got %x.", n)
	}
	theByte := uint8(n / 8)
	theBit := uint8(n % 8)
	return theByte, theBit, nil
}

// Add addes a Neighborhood to the Set.
func (s *Set) Add(n Neighborhood) error {
	theByte, theBit, err := setAddress(n)
	if err != nil {
		return err
	}
  (*s)[theByte] |= (b10000000 >> theBit)
	return nil
}

// Contains checks whether the Set contains a Neighborhood.
func (s *Set) Contains(n Neighborhood) (bool, error) {
	theByte, theBit, err := setAddress(n)
	if err != nil {
		return false, err
	}
	return (*s)[theByte] & (b10000000 >> theBit) > 0, nil
}

// Equals returns true if neighborhoods are equal.
func Equals(first *Set, second *Set) bool {
	for i := 0; i < 64; i++ {
		if (*first)[i] != (*second)[i] {
			return false
		}
	}
	return true
}

// Implements fmt.Stringer
func (s *Set) String() string {
	ns := []Neighborhood{}
	for n := 0; n < 0x200; n++ {
		c, err := s.Contains(Neighborhood(n))
		if err != nil {
			return fmt.Sprintf("[ERROR: %v]", err)
		}
		if c {
			ns = append(ns, Neighborhood(n))
		}
	}
	if len(ns) == 0 {
		return "[]"
	}
	rv := fmt.Sprintf("[%q", ns[0].ToStr())
	for _, n := range ns[1:] {
		rv = fmt.Sprintf("%s, %q", rv, n.ToStr())
	}
	rv = fmt.Sprintf("%s]", rv)
	return rv
}

// ShiftIntersect performs "shift & intersect" operation.
//
// Given two sets of neighborhoods (left and right), tries to find all elements
// that overlap with at least one element in the other set if shifted to the
// given side.
//
// Returns two sets containing "overlapping" elements each of the inupt sets.
func ShiftIntersect(left *Set, right *Set, s Side) (*Set, *Set, error) {
	resL := &Set{}
	resR := &Set{}
	var l, r Neighborhood
	for l = 0; l < 0x200; l++ {
		lc, err := left.Contains(l)
		if err != nil {
			return nil, nil, fmt.Errorf("ShiftIntersect left.Contains: %v", err)
		}
		if !lc {
			continue
		}
		// TODO(pawelz@execve.pl): Set needs some iterator util.
		for r = 0; r < 0x200; r++ {
			rc, err := right.Contains(r)
			if err != nil {
				return nil, nil, fmt.Errorf("ShiftIntersect right.Contains: %v", err)
			}
			if !rc {
				continue
			}
			match, err := l.Matches(r, 1, s)
			if err != nil {
				return nil, nil, fmt.Errorf("ShiftIntersect Matches: %v", err)
			}
			if match {
				if err := resL.Add(l); err != nil {
					return nil, nil, fmt.Errorf("ShiftIntersect resL.Add: %v", err)
				}
				if err := resR.Add(r); err != nil {
					return nil, nil, fmt.Errorf("ShiftIntersect resR.Add: %v", err)
				}
			}
		}
	}
	return resL, resR, nil
}
