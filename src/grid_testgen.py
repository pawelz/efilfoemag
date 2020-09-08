#!/bin/python3

# Copyright 2020 Google LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     https://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

"""Generates testcases for grid.torusGet"""

grid = [
    "++#+++++",
    "#++#++++",
    "#+###+++",
    "######++",
    "++#####+",
    "+++###++",
    "++++#+++",
    "+++++++#",
]

cases = []
for x in range(-1, 9):
  for y in range(-1, 9):
    cases.append(
"""		{{
			x: {},
			y: {},
			expected: state.{},
		}},""".format(x, y, "Alive" if grid[y % 8][x % 8] == "#" else "Dead"))

print("""
package grid

import (
	"fmt"
	"testing"

	"github.com/pawelz/efilfoemag/src/state"
)

func TestTorusGet(t *testing.T) {{
	testGrid, err := Parse([]byte(`8x8
{}
`))
	if err != nil {{
                t.Fatalf("Cannot Parse test data: %v", err)
	}}
	for _, td := range []struct {{
		x        int
		y        int
		expected state.State
	}}{{
{}
	}}{{
		t.Run(fmt.Sprintf("x=%d,y=%d", td.x, td.y), func(t *testing.T) {{
			actual, err := testGrid.torusGet(td.x, td.y)
			if err != nil {{
				t.Errorf("expected no error; got %v", err)
			}}
			if actual != td.expected {{
				t.Errorf("for (x=%d, y=%d) want %v; got %v", td.x, td.y, td.expected.ToStr(), actual.ToStr())
			}}
		}})
	}}
}}
""".format("\n".join(grid), "\n".join(cases)))
