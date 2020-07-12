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

package bits

import (
	"log"
)

// Sumbit returns number of bits set to 1.
func Sum(c uint16) uint16 {
	var s uint16
	for i := 0; i < 9; i++ {
		s += c & 1
		c >>= 1
	}
	return s
}

// Parses string like "01101001" as uint8.
// TODO(pawelz@execve.pl) replace this with ob-literalz once we're on go v1.13
//                        https://tip.golang.org/ref/spec#Integer_literals
func Byte(bs string) uint8 {
	if l := len(bs); l != 8 {
		log.Fatalf("expected 8 bits, got %d: %q", l, bs) // srsly
	}
	var v uint8
	for n, b := range bs {
		switch b {
		case '1':
			v |= (uint8(1) << (7 - n))
		case '0':
			// do nothing
		default:
			log.Fatalf("unexpected bit value %c in %q", b, bs)
		}
	}
	return v
}
