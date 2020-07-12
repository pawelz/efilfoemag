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
	"bufio"
	"bytes"
	"fmt"
	"strconv"

	"github.com/pawelz/efilfoemag/src/bits"
	"github.com/pawelz/efilfoemag/src/state"
)

const (
	endl = byte('\n')
)

type Canvas struct {
	width  uint
	height uint
	b      []uint8
}

// create is a factory of blank Canvas objects.
func create(width, height int) (*Canvas, error) {
	if width <= 0 || height <= 0 {
		return nil, fmt.Errorf("width and height of a Canvas must be positive, got: width = %d, height = %d", width, height)
	}
	if w, h := width%8, height%8; w != 0 || h != 0 {
		return nil, fmt.Errorf("width and height of a Canvas must be divisible by 8, got width = %d (mod 8 = %d), height = %d (mod 8 = %d)", width, w, height, h)
	}
	if width < 8 || height < 8 {
		return nil, fmt.Errorf("width and height of a Canvas must be greater or equal 8, got width = %d, height = %d", width, height)
	}
	return &Canvas{
		width:  uint(width),
		height: uint(height),
		b:      make([]uint8, width*height/8),
	}, nil
}

func stripFinalChar(s string) string {
	if s == "" {
		return ""
	}
	return s[:len(s)-1]
}

// Parse parses the content of .efil file to produce a Canvas object.
func Parse(inputData []byte) (*Canvas, error) {
	r := bufio.NewReader(bytes.NewReader(inputData))
	widthString, err := r.ReadString(byte('x'))
	if err != nil {
		return nil, fmt.Errorf("error reading width: %v", err)
	}
	widthString = stripFinalChar(widthString)
	heightString, err := r.ReadString(endl)
	if err != nil {
		return nil, fmt.Errorf("error reading height: %v", err)
	}
	heightString = stripFinalChar(heightString)
	width, err := strconv.Atoi(widthString)
	if err != nil {
		return nil, fmt.Errorf("error parsing width %q: %v", widthString, err)
	}
	height, err := strconv.Atoi(heightString)
	if err != nil {
		return nil, fmt.Errorf("error parsing height %q: %v", heightString, err)
	}

	canvas, err := create(width, height)
	if err != nil {
		return nil, fmt.Errorf("error creating canvas: %v", err)
	}

	for rowNum := 0; rowNum < height; rowNum++ {
		rowData, err := r.ReadBytes(endl)
		if err != nil {
			return nil, fmt.Errorf("error creating canvas while reading row %d: %v", rowNum, err)
		}
		if l := len(rowData); l != width+1 {
			return nil, fmt.Errorf("error reading row %d, want %d characters (including \\n), got %d", rowNum, width+1, l)
		}
		for oNum := 0; oNum*8 < width; oNum++ {
			var octet uint8
			for bitNum := 0; bitNum < 8; bitNum++ {
				symbol := rowData[oNum*8+bitNum]
				switch symbol {
				case '#':
					octet |= (bits.Byte("10000000") >> bitNum)
				case '+':
					// pass
				default:
					return nil, fmt.Errorf("encountered invalid byte %c at (%d, %d)", symbol, rowNum, oNum*8+bitNum)
				}
			}
			canvas.b[rowNum*width/8+oNum] = octet
		}
	}

	return canvas, nil
}

func (c *Canvas) byteshift(x, y uint) uint {
	return y*c.width/8 + x/8
}

func bitmask(x uint) uint8 {
	return bits.Byte("10000000") >> (x % 8)
}

func (c *Canvas) validateAddress(x, y uint) error {
	if x >= c.width || y >= c.height {
		return fmt.Errorf("address (%d, %d) is out of band (0-%d, 0-%d)", x, y, c.width, c.height)
	}
	return nil
}

// Get returns the state of the cell at the given address.
func (c *Canvas) Get(x, y uint) (state.State, error) {
	if err := c.validateAddress(x, y); err != nil {
		return state.Dead, fmt.Errorf("cannot Get: %v", err)
	}
	octet := c.b[c.byteshift(x, y)]
	if octet&bitmask(x) != uint8(0) {
		return state.Alive, nil
	}
	return state.Dead, nil
}

// Set sets the state of the cell at the given address.
func (c *Canvas) Set(x, y uint, s state.State) error {
	if err := c.validateAddress(x, y); err != nil {
		return fmt.Errorf("cannot Get: %v", err)
	}
	c.b[c.byteshift(x, y)] |= bitmask(x)
	return nil
}

// equalsTo compares this canvas to anover one.
func (c *Canvas) equalsTo(other *Canvas) bool {
	if c.height != other.height {
		return false
	}
	if c.width != other.width {
		return false
	}
	// this hould not happen for valid Canas.
	if len(c.b) != len(other.b) {
		return false
	}
	for i := 0; i < len(c.b); i++ {
		if c.b[i] != other.b[i] {
			return false
		}
	}
	return true
}
