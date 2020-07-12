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

package state

type State bool

const (
	Alive State = true
	Dead State = false
)

// Of converts a boolean to the corresponding State.
func Of(isAlive bool) State {
	if isAlive {
		return Alive
	}
	return Dead
}

// IsAlive converts a State to the corresponding boolean.
func (s State) IsAlive() bool {
	return s == Alive
}

// ToRune converts a State to its corresponding rune. Alive -> '#'; Dead -> '+'.
func (s State) ToRune() rune {
	if s.IsAlive() {
		return '#'
	}
	return '+'
}

// ToStr converts a State to a human readable string.
func (s State) ToStr() string {
	if s.IsAlive() {
		return "Alive"
	}
	return "Dead"
}
