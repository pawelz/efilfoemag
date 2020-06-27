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

package canvas

import (
	"testing"

	"github.com/pawelz/efilfoemag/src/bits"
)

func TestEqualsTo(t *testing.T) {
	for _, td := range []struct {
		name     string
		this     *Canvas
		that     *Canvas
		expected bool
	}{
		{
			name: "simply equal",
			this: &Canvas{
				width:  8,
				height: 8,
				b: []uint8{
					bits.Byte("00000000"),
					bits.Byte("00010000"),
					bits.Byte("00111000"),
					bits.Byte("01111100"),
					bits.Byte("00111110"),
					bits.Byte("00011100"),
					bits.Byte("00001000"),
					bits.Byte("00000000"),
				},
			},
			that: &Canvas{
				width:  8,
				height: 8,
				b: []uint8{
					bits.Byte("00000000"),
					bits.Byte("00010000"),
					bits.Byte("00111000"),
					bits.Byte("01111100"),
					bits.Byte("00111110"),
					bits.Byte("00011100"),
					bits.Byte("00001000"),
					bits.Byte("00000000"),
				},
			},
			expected: true,
		},
		{
			name: "bad width",
			this: &Canvas{
				width:  8,
				height: 8,
				b: []uint8{
					bits.Byte("00000000"),
					bits.Byte("00010000"),
					bits.Byte("00111000"),
					bits.Byte("01111100"),
					bits.Byte("00111110"),
					bits.Byte("00011100"),
					bits.Byte("00001000"),
					bits.Byte("00000000"),
				},
			},
			that: &Canvas{
				width:  16,
				height: 8,
				b: []uint8{
					bits.Byte("00000000"),
					bits.Byte("00010000"),
					bits.Byte("00111000"),
					bits.Byte("01111100"),
					bits.Byte("00111110"),
					bits.Byte("00011100"),
					bits.Byte("00001000"),
					bits.Byte("00000000"),
				},
			},
			expected: false,
		},
		{
			name: "bad height",
			this: &Canvas{
				width:  8,
				height: 8,
				b: []uint8{
					bits.Byte("00000000"),
					bits.Byte("00010000"),
					bits.Byte("00111000"),
					bits.Byte("01111100"),
					bits.Byte("00111110"),
					bits.Byte("00011100"),
					bits.Byte("00001000"),
					bits.Byte("00000000"),
				},
			},
			that: &Canvas{
				width:  8,
				height: 16,
				b: []uint8{
					bits.Byte("00000000"),
					bits.Byte("00010000"),
					bits.Byte("00111000"),
					bits.Byte("01111100"),
					bits.Byte("00111110"),
					bits.Byte("00011100"),
					bits.Byte("00001000"),
					bits.Byte("00000000"),
				},
			},
			expected: false,
		},
		{
			name: "bad pattern",
			this: &Canvas{
				width:  8,
				height: 8,
				b: []byte{
					bits.Byte("00000000"),
					bits.Byte("00010000"),
					bits.Byte("00111000"),
					bits.Byte("01111100"),
					bits.Byte("00111110"),
					bits.Byte("00011100"),
					bits.Byte("00001000"),
					bits.Byte("00000000"),
				},
			},
			that: &Canvas{
				width:  8,
				height: 8,
				b: []byte{
					bits.Byte("00000000"),
					bits.Byte("00010000"),
					bits.Byte("00101000"),
					bits.Byte("01000100"),
					bits.Byte("00100010"),
					bits.Byte("00010100"),
					bits.Byte("00001000"),
					bits.Byte("00000000"),
				},
			},
			expected: false,
		},
	} {
		t.Run(td.name, func(t *testing.T) {
			if actual := td.this.equalsTo(td.that); actual != td.expected {
				t.Errorf("want %v, got %v", td.expected, actual)
			}
		})
	}
}

func TestParse(t *testing.T) {
	for _, td := range []struct {
		name     string
		input    []byte
		expected *Canvas
		failure  bool
	}{
		{
			name: "simple 8x8",
			input: []byte(`8x8
++++++++
+++#++++
++###+++
+#####++
++#####+
+++###++
++++#+++
++++++++
`),
			expected: &Canvas{
				width:  8,
				height: 8,
				b: []byte{
					bits.Byte("00000000"),
					bits.Byte("00010000"),
					bits.Byte("00111000"),
					bits.Byte("01111100"),
					bits.Byte("00111110"),
					bits.Byte("00011100"),
					bits.Byte("00001000"),
					bits.Byte("00000000"),
				},
			},
		},
		{
			name: "badHeight",
			input: []byte(`8x7
++++++++
+++#++++
++###+++
+#####++
++#####+
+++###++
++++#+++
++++++++
`),
			failure: true,
		},
	} {
		t.Run(td.name, func(t *testing.T) {
			actual, err := Parse(td.input)

			if td.failure {
				if err == nil {
					t.Errorf("expected a failure, got %x", actual)
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if !actual.equalsTo(td.expected) {
				t.Errorf("expected %x, got %x", td.expected, actual)
			}
		})
	}
}
