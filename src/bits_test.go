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
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	for _, td := range []struct {
		number   uint16
		expected uint16
	}{
		{number: 0x000, expected: 0},
		{number: 0x001, expected: 1},
		{number: 0x002, expected: 1},
		{number: 0x003, expected: 2},
		{number: 0x004, expected: 1},
		{number: 0x005, expected: 2},
		{number: 0x006, expected: 2},
		{number: 0x007, expected: 3},
		{number: 0x008, expected: 1},
		{number: 0x009, expected: 2},
		{number: 0x00a, expected: 2},
		{number: 0x00b, expected: 3},
		{number: 0x00c, expected: 2},
		{number: 0x00d, expected: 3},
		{number: 0x00e, expected: 3},
		{number: 0x00f, expected: 4},
		{number: 0x010, expected: 1},
		{number: 0x011, expected: 2},
		{number: 0x012, expected: 2},
		{number: 0x013, expected: 3},
		{number: 0x014, expected: 2},
		{number: 0x015, expected: 3},
		{number: 0x016, expected: 3},
		{number: 0x017, expected: 4},
		{number: 0x018, expected: 2},
		{number: 0x019, expected: 3},
		{number: 0x01a, expected: 3},
		{number: 0x01b, expected: 4},
		{number: 0x01c, expected: 3},
		{number: 0x01d, expected: 4},
		{number: 0x01e, expected: 4},
		{number: 0x01f, expected: 5},
		{number: 0x020, expected: 1},
		{number: 0x021, expected: 2},
		{number: 0x022, expected: 2},
		{number: 0x023, expected: 3},
		{number: 0x024, expected: 2},
		{number: 0x025, expected: 3},
		{number: 0x026, expected: 3},
		{number: 0x027, expected: 4},
		{number: 0x028, expected: 2},
		{number: 0x029, expected: 3},
		{number: 0x02a, expected: 3},
		{number: 0x02b, expected: 4},
		{number: 0x02c, expected: 3},
		{number: 0x02d, expected: 4},
		{number: 0x02e, expected: 4},
		{number: 0x02f, expected: 5},
		{number: 0x030, expected: 2},
		{number: 0x031, expected: 3},
		{number: 0x032, expected: 3},
		{number: 0x033, expected: 4},
		{number: 0x034, expected: 3},
		{number: 0x035, expected: 4},
		{number: 0x036, expected: 4},
		{number: 0x037, expected: 5},
		{number: 0x038, expected: 3},
		{number: 0x039, expected: 4},
		{number: 0x03a, expected: 4},
		{number: 0x03b, expected: 5},
		{number: 0x03c, expected: 4},
		{number: 0x03d, expected: 5},
		{number: 0x03e, expected: 5},
		{number: 0x03f, expected: 6},
		{number: 0x040, expected: 1},
		{number: 0x041, expected: 2},
		{number: 0x042, expected: 2},
		{number: 0x043, expected: 3},
		{number: 0x044, expected: 2},
		{number: 0x045, expected: 3},
		{number: 0x046, expected: 3},
		{number: 0x047, expected: 4},
		{number: 0x048, expected: 2},
		{number: 0x049, expected: 3},
		{number: 0x04a, expected: 3},
		{number: 0x04b, expected: 4},
		{number: 0x04c, expected: 3},
		{number: 0x04d, expected: 4},
		{number: 0x04e, expected: 4},
		{number: 0x04f, expected: 5},
		{number: 0x050, expected: 2},
		{number: 0x051, expected: 3},
		{number: 0x052, expected: 3},
		{number: 0x053, expected: 4},
		{number: 0x054, expected: 3},
		{number: 0x055, expected: 4},
		{number: 0x056, expected: 4},
		{number: 0x057, expected: 5},
		{number: 0x058, expected: 3},
		{number: 0x059, expected: 4},
		{number: 0x05a, expected: 4},
		{number: 0x05b, expected: 5},
		{number: 0x05c, expected: 4},
		{number: 0x05d, expected: 5},
		{number: 0x05e, expected: 5},
		{number: 0x05f, expected: 6},
		{number: 0x060, expected: 2},
		{number: 0x061, expected: 3},
		{number: 0x062, expected: 3},
		{number: 0x063, expected: 4},
		{number: 0x064, expected: 3},
		{number: 0x065, expected: 4},
		{number: 0x066, expected: 4},
		{number: 0x067, expected: 5},
		{number: 0x068, expected: 3},
		{number: 0x069, expected: 4},
		{number: 0x06a, expected: 4},
		{number: 0x06b, expected: 5},
		{number: 0x06c, expected: 4},
		{number: 0x06d, expected: 5},
		{number: 0x06e, expected: 5},
		{number: 0x06f, expected: 6},
		{number: 0x070, expected: 3},
		{number: 0x071, expected: 4},
		{number: 0x072, expected: 4},
		{number: 0x073, expected: 5},
		{number: 0x074, expected: 4},
		{number: 0x075, expected: 5},
		{number: 0x076, expected: 5},
		{number: 0x077, expected: 6},
		{number: 0x078, expected: 4},
		{number: 0x079, expected: 5},
		{number: 0x07a, expected: 5},
		{number: 0x07b, expected: 6},
		{number: 0x07c, expected: 5},
		{number: 0x07d, expected: 6},
		{number: 0x07e, expected: 6},
		{number: 0x07f, expected: 7},
		{number: 0x080, expected: 1},
		{number: 0x081, expected: 2},
		{number: 0x082, expected: 2},
		{number: 0x083, expected: 3},
		{number: 0x084, expected: 2},
		{number: 0x085, expected: 3},
		{number: 0x086, expected: 3},
		{number: 0x087, expected: 4},
		{number: 0x088, expected: 2},
		{number: 0x089, expected: 3},
		{number: 0x08a, expected: 3},
		{number: 0x08b, expected: 4},
		{number: 0x08c, expected: 3},
		{number: 0x08d, expected: 4},
		{number: 0x08e, expected: 4},
		{number: 0x08f, expected: 5},
		{number: 0x090, expected: 2},
		{number: 0x091, expected: 3},
		{number: 0x092, expected: 3},
		{number: 0x093, expected: 4},
		{number: 0x094, expected: 3},
		{number: 0x095, expected: 4},
		{number: 0x096, expected: 4},
		{number: 0x097, expected: 5},
		{number: 0x098, expected: 3},
		{number: 0x099, expected: 4},
		{number: 0x09a, expected: 4},
		{number: 0x09b, expected: 5},
		{number: 0x09c, expected: 4},
		{number: 0x09d, expected: 5},
		{number: 0x09e, expected: 5},
		{number: 0x09f, expected: 6},
		{number: 0x0a0, expected: 2},
		{number: 0x0a1, expected: 3},
		{number: 0x0a2, expected: 3},
		{number: 0x0a3, expected: 4},
		{number: 0x0a4, expected: 3},
		{number: 0x0a5, expected: 4},
		{number: 0x0a6, expected: 4},
		{number: 0x0a7, expected: 5},
		{number: 0x0a8, expected: 3},
		{number: 0x0a9, expected: 4},
		{number: 0x0aa, expected: 4},
		{number: 0x0ab, expected: 5},
		{number: 0x0ac, expected: 4},
		{number: 0x0ad, expected: 5},
		{number: 0x0ae, expected: 5},
		{number: 0x0af, expected: 6},
		{number: 0x0b0, expected: 3},
		{number: 0x0b1, expected: 4},
		{number: 0x0b2, expected: 4},
		{number: 0x0b3, expected: 5},
		{number: 0x0b4, expected: 4},
		{number: 0x0b5, expected: 5},
		{number: 0x0b6, expected: 5},
		{number: 0x0b7, expected: 6},
		{number: 0x0b8, expected: 4},
		{number: 0x0b9, expected: 5},
		{number: 0x0ba, expected: 5},
		{number: 0x0bb, expected: 6},
		{number: 0x0bc, expected: 5},
		{number: 0x0bd, expected: 6},
		{number: 0x0be, expected: 6},
		{number: 0x0bf, expected: 7},
		{number: 0x0c0, expected: 2},
		{number: 0x0c1, expected: 3},
		{number: 0x0c2, expected: 3},
		{number: 0x0c3, expected: 4},
		{number: 0x0c4, expected: 3},
		{number: 0x0c5, expected: 4},
		{number: 0x0c6, expected: 4},
		{number: 0x0c7, expected: 5},
		{number: 0x0c8, expected: 3},
		{number: 0x0c9, expected: 4},
		{number: 0x0ca, expected: 4},
		{number: 0x0cb, expected: 5},
		{number: 0x0cc, expected: 4},
		{number: 0x0cd, expected: 5},
		{number: 0x0ce, expected: 5},
		{number: 0x0cf, expected: 6},
		{number: 0x0d0, expected: 3},
		{number: 0x0d1, expected: 4},
		{number: 0x0d2, expected: 4},
		{number: 0x0d3, expected: 5},
		{number: 0x0d4, expected: 4},
		{number: 0x0d5, expected: 5},
		{number: 0x0d6, expected: 5},
		{number: 0x0d7, expected: 6},
		{number: 0x0d8, expected: 4},
		{number: 0x0d9, expected: 5},
		{number: 0x0da, expected: 5},
		{number: 0x0db, expected: 6},
		{number: 0x0dc, expected: 5},
		{number: 0x0dd, expected: 6},
		{number: 0x0de, expected: 6},
		{number: 0x0df, expected: 7},
		{number: 0x0e0, expected: 3},
		{number: 0x0e1, expected: 4},
		{number: 0x0e2, expected: 4},
		{number: 0x0e3, expected: 5},
		{number: 0x0e4, expected: 4},
		{number: 0x0e5, expected: 5},
		{number: 0x0e6, expected: 5},
		{number: 0x0e7, expected: 6},
		{number: 0x0e8, expected: 4},
		{number: 0x0e9, expected: 5},
		{number: 0x0ea, expected: 5},
		{number: 0x0eb, expected: 6},
		{number: 0x0ec, expected: 5},
		{number: 0x0ed, expected: 6},
		{number: 0x0ee, expected: 6},
		{number: 0x0ef, expected: 7},
		{number: 0x0f0, expected: 4},
		{number: 0x0f1, expected: 5},
		{number: 0x0f2, expected: 5},
		{number: 0x0f3, expected: 6},
		{number: 0x0f4, expected: 5},
		{number: 0x0f5, expected: 6},
		{number: 0x0f6, expected: 6},
		{number: 0x0f7, expected: 7},
		{number: 0x0f8, expected: 5},
		{number: 0x0f9, expected: 6},
		{number: 0x0fa, expected: 6},
		{number: 0x0fb, expected: 7},
		{number: 0x0fc, expected: 6},
		{number: 0x0fd, expected: 7},
		{number: 0x0fe, expected: 7},
		{number: 0x0ff, expected: 8},
		{number: 0x100, expected: 1},
		{number: 0x101, expected: 2},
		{number: 0x102, expected: 2},
		{number: 0x103, expected: 3},
		{number: 0x104, expected: 2},
		{number: 0x105, expected: 3},
		{number: 0x106, expected: 3},
		{number: 0x107, expected: 4},
		{number: 0x108, expected: 2},
		{number: 0x109, expected: 3},
		{number: 0x10a, expected: 3},
		{number: 0x10b, expected: 4},
		{number: 0x10c, expected: 3},
		{number: 0x10d, expected: 4},
		{number: 0x10e, expected: 4},
		{number: 0x10f, expected: 5},
		{number: 0x110, expected: 2},
		{number: 0x111, expected: 3},
		{number: 0x112, expected: 3},
		{number: 0x113, expected: 4},
		{number: 0x114, expected: 3},
		{number: 0x115, expected: 4},
		{number: 0x116, expected: 4},
		{number: 0x117, expected: 5},
		{number: 0x118, expected: 3},
		{number: 0x119, expected: 4},
		{number: 0x11a, expected: 4},
		{number: 0x11b, expected: 5},
		{number: 0x11c, expected: 4},
		{number: 0x11d, expected: 5},
		{number: 0x11e, expected: 5},
		{number: 0x11f, expected: 6},
		{number: 0x120, expected: 2},
		{number: 0x121, expected: 3},
		{number: 0x122, expected: 3},
		{number: 0x123, expected: 4},
		{number: 0x124, expected: 3},
		{number: 0x125, expected: 4},
		{number: 0x126, expected: 4},
		{number: 0x127, expected: 5},
		{number: 0x128, expected: 3},
		{number: 0x129, expected: 4},
		{number: 0x12a, expected: 4},
		{number: 0x12b, expected: 5},
		{number: 0x12c, expected: 4},
		{number: 0x12d, expected: 5},
		{number: 0x12e, expected: 5},
		{number: 0x12f, expected: 6},
		{number: 0x130, expected: 3},
		{number: 0x131, expected: 4},
		{number: 0x132, expected: 4},
		{number: 0x133, expected: 5},
		{number: 0x134, expected: 4},
		{number: 0x135, expected: 5},
		{number: 0x136, expected: 5},
		{number: 0x137, expected: 6},
		{number: 0x138, expected: 4},
		{number: 0x139, expected: 5},
		{number: 0x13a, expected: 5},
		{number: 0x13b, expected: 6},
		{number: 0x13c, expected: 5},
		{number: 0x13d, expected: 6},
		{number: 0x13e, expected: 6},
		{number: 0x13f, expected: 7},
		{number: 0x140, expected: 2},
		{number: 0x141, expected: 3},
		{number: 0x142, expected: 3},
		{number: 0x143, expected: 4},
		{number: 0x144, expected: 3},
		{number: 0x145, expected: 4},
		{number: 0x146, expected: 4},
		{number: 0x147, expected: 5},
		{number: 0x148, expected: 3},
		{number: 0x149, expected: 4},
		{number: 0x14a, expected: 4},
		{number: 0x14b, expected: 5},
		{number: 0x14c, expected: 4},
		{number: 0x14d, expected: 5},
		{number: 0x14e, expected: 5},
		{number: 0x14f, expected: 6},
		{number: 0x150, expected: 3},
		{number: 0x151, expected: 4},
		{number: 0x152, expected: 4},
		{number: 0x153, expected: 5},
		{number: 0x154, expected: 4},
		{number: 0x155, expected: 5},
		{number: 0x156, expected: 5},
		{number: 0x157, expected: 6},
		{number: 0x158, expected: 4},
		{number: 0x159, expected: 5},
		{number: 0x15a, expected: 5},
		{number: 0x15b, expected: 6},
		{number: 0x15c, expected: 5},
		{number: 0x15d, expected: 6},
		{number: 0x15e, expected: 6},
		{number: 0x15f, expected: 7},
		{number: 0x160, expected: 3},
		{number: 0x161, expected: 4},
		{number: 0x162, expected: 4},
		{number: 0x163, expected: 5},
		{number: 0x164, expected: 4},
		{number: 0x165, expected: 5},
		{number: 0x166, expected: 5},
		{number: 0x167, expected: 6},
		{number: 0x168, expected: 4},
		{number: 0x169, expected: 5},
		{number: 0x16a, expected: 5},
		{number: 0x16b, expected: 6},
		{number: 0x16c, expected: 5},
		{number: 0x16d, expected: 6},
		{number: 0x16e, expected: 6},
		{number: 0x16f, expected: 7},
		{number: 0x170, expected: 4},
		{number: 0x171, expected: 5},
		{number: 0x172, expected: 5},
		{number: 0x173, expected: 6},
		{number: 0x174, expected: 5},
		{number: 0x175, expected: 6},
		{number: 0x176, expected: 6},
		{number: 0x177, expected: 7},
		{number: 0x178, expected: 5},
		{number: 0x179, expected: 6},
		{number: 0x17a, expected: 6},
		{number: 0x17b, expected: 7},
		{number: 0x17c, expected: 6},
		{number: 0x17d, expected: 7},
		{number: 0x17e, expected: 7},
		{number: 0x17f, expected: 8},
		{number: 0x180, expected: 2},
		{number: 0x181, expected: 3},
		{number: 0x182, expected: 3},
		{number: 0x183, expected: 4},
		{number: 0x184, expected: 3},
		{number: 0x185, expected: 4},
		{number: 0x186, expected: 4},
		{number: 0x187, expected: 5},
		{number: 0x188, expected: 3},
		{number: 0x189, expected: 4},
		{number: 0x18a, expected: 4},
		{number: 0x18b, expected: 5},
		{number: 0x18c, expected: 4},
		{number: 0x18d, expected: 5},
		{number: 0x18e, expected: 5},
		{number: 0x18f, expected: 6},
		{number: 0x190, expected: 3},
		{number: 0x191, expected: 4},
		{number: 0x192, expected: 4},
		{number: 0x193, expected: 5},
		{number: 0x194, expected: 4},
		{number: 0x195, expected: 5},
		{number: 0x196, expected: 5},
		{number: 0x197, expected: 6},
		{number: 0x198, expected: 4},
		{number: 0x199, expected: 5},
		{number: 0x19a, expected: 5},
		{number: 0x19b, expected: 6},
		{number: 0x19c, expected: 5},
		{number: 0x19d, expected: 6},
		{number: 0x19e, expected: 6},
		{number: 0x19f, expected: 7},
		{number: 0x1a0, expected: 3},
		{number: 0x1a1, expected: 4},
		{number: 0x1a2, expected: 4},
		{number: 0x1a3, expected: 5},
		{number: 0x1a4, expected: 4},
		{number: 0x1a5, expected: 5},
		{number: 0x1a6, expected: 5},
		{number: 0x1a7, expected: 6},
		{number: 0x1a8, expected: 4},
		{number: 0x1a9, expected: 5},
		{number: 0x1aa, expected: 5},
		{number: 0x1ab, expected: 6},
		{number: 0x1ac, expected: 5},
		{number: 0x1ad, expected: 6},
		{number: 0x1ae, expected: 6},
		{number: 0x1af, expected: 7},
		{number: 0x1b0, expected: 4},
		{number: 0x1b1, expected: 5},
		{number: 0x1b2, expected: 5},
		{number: 0x1b3, expected: 6},
		{number: 0x1b4, expected: 5},
		{number: 0x1b5, expected: 6},
		{number: 0x1b6, expected: 6},
		{number: 0x1b7, expected: 7},
		{number: 0x1b8, expected: 5},
		{number: 0x1b9, expected: 6},
		{number: 0x1ba, expected: 6},
		{number: 0x1bb, expected: 7},
		{number: 0x1bc, expected: 6},
		{number: 0x1bd, expected: 7},
		{number: 0x1be, expected: 7},
		{number: 0x1bf, expected: 8},
		{number: 0x1c0, expected: 3},
		{number: 0x1c1, expected: 4},
		{number: 0x1c2, expected: 4},
		{number: 0x1c3, expected: 5},
		{number: 0x1c4, expected: 4},
		{number: 0x1c5, expected: 5},
		{number: 0x1c6, expected: 5},
		{number: 0x1c7, expected: 6},
		{number: 0x1c8, expected: 4},
		{number: 0x1c9, expected: 5},
		{number: 0x1ca, expected: 5},
		{number: 0x1cb, expected: 6},
		{number: 0x1cc, expected: 5},
		{number: 0x1cd, expected: 6},
		{number: 0x1ce, expected: 6},
		{number: 0x1cf, expected: 7},
		{number: 0x1d0, expected: 4},
		{number: 0x1d1, expected: 5},
		{number: 0x1d2, expected: 5},
		{number: 0x1d3, expected: 6},
		{number: 0x1d4, expected: 5},
		{number: 0x1d5, expected: 6},
		{number: 0x1d6, expected: 6},
		{number: 0x1d7, expected: 7},
		{number: 0x1d8, expected: 5},
		{number: 0x1d9, expected: 6},
		{number: 0x1da, expected: 6},
		{number: 0x1db, expected: 7},
		{number: 0x1dc, expected: 6},
		{number: 0x1dd, expected: 7},
		{number: 0x1de, expected: 7},
		{number: 0x1df, expected: 8},
		{number: 0x1e0, expected: 4},
		{number: 0x1e1, expected: 5},
		{number: 0x1e2, expected: 5},
		{number: 0x1e3, expected: 6},
		{number: 0x1e4, expected: 5},
		{number: 0x1e5, expected: 6},
		{number: 0x1e6, expected: 6},
		{number: 0x1e7, expected: 7},
		{number: 0x1e8, expected: 5},
		{number: 0x1e9, expected: 6},
		{number: 0x1ea, expected: 6},
		{number: 0x1eb, expected: 7},
		{number: 0x1ec, expected: 6},
		{number: 0x1ed, expected: 7},
		{number: 0x1ee, expected: 7},
		{number: 0x1ef, expected: 8},
		{number: 0x1f0, expected: 5},
		{number: 0x1f1, expected: 6},
		{number: 0x1f2, expected: 6},
		{number: 0x1f3, expected: 7},
		{number: 0x1f4, expected: 6},
		{number: 0x1f5, expected: 7},
		{number: 0x1f6, expected: 7},
		{number: 0x1f7, expected: 8},
		{number: 0x1f8, expected: 6},
		{number: 0x1f9, expected: 7},
		{number: 0x1fa, expected: 7},
		{number: 0x1fb, expected: 8},
		{number: 0x1fc, expected: 7},
		{number: 0x1fd, expected: 8},
		{number: 0x1fe, expected: 8},
		{number: 0x1ff, expected: 9},
	} {
		t.Run(fmt.Sprintf("Test for %x", td.number), func(t *testing.T) {
			actual := Sum(td.number)
			if actual != td.expected {
				t.Errorf("excpected %d, got %d", td.expected, actual)
			}
		})
	}
}

func TestByte(t *testing.T) {
	for _, td := range []struct {
		input    string
		expected uint8
	}{
		{input: "00000000", expected: 0},
		{input: "00000001", expected: 1},
		{input: "00000010", expected: 2},
		{input: "00000011", expected: 3},
		{input: "00000100", expected: 4},
		{input: "00000101", expected: 5},
		{input: "00000110", expected: 6},
		{input: "00000111", expected: 7},
		{input: "00001000", expected: 8},
		{input: "00001001", expected: 9},
		{input: "00001010", expected: 10},
		{input: "00001011", expected: 11},
		{input: "00001100", expected: 12},
		{input: "00001101", expected: 13},
		{input: "00001110", expected: 14},
		{input: "00001111", expected: 15},
		{input: "00010000", expected: 16},
		{input: "00010001", expected: 17},
		{input: "00010010", expected: 18},
		{input: "00010011", expected: 19},
		{input: "00010100", expected: 20},
		{input: "00010101", expected: 21},
		{input: "00010110", expected: 22},
		{input: "00010111", expected: 23},
		{input: "00011000", expected: 24},
		{input: "00011001", expected: 25},
		{input: "00011010", expected: 26},
		{input: "00011011", expected: 27},
		{input: "00011100", expected: 28},
		{input: "00011101", expected: 29},
		{input: "00011110", expected: 30},
		{input: "00011111", expected: 31},
		{input: "00100000", expected: 32},
		{input: "00100001", expected: 33},
		{input: "00100010", expected: 34},
		{input: "00100011", expected: 35},
		{input: "00100100", expected: 36},
		{input: "00100101", expected: 37},
		{input: "00100110", expected: 38},
		{input: "00100111", expected: 39},
		{input: "00101000", expected: 40},
		{input: "00101001", expected: 41},
		{input: "00101010", expected: 42},
		{input: "00101011", expected: 43},
		{input: "00101100", expected: 44},
		{input: "00101101", expected: 45},
		{input: "00101110", expected: 46},
		{input: "00101111", expected: 47},
		{input: "00110000", expected: 48},
		{input: "00110001", expected: 49},
		{input: "00110010", expected: 50},
		{input: "00110011", expected: 51},
		{input: "00110100", expected: 52},
		{input: "00110101", expected: 53},
		{input: "00110110", expected: 54},
		{input: "00110111", expected: 55},
		{input: "00111000", expected: 56},
		{input: "00111001", expected: 57},
		{input: "00111010", expected: 58},
		{input: "00111011", expected: 59},
		{input: "00111100", expected: 60},
		{input: "00111101", expected: 61},
		{input: "00111110", expected: 62},
		{input: "00111111", expected: 63},
		{input: "01000000", expected: 64},
		{input: "01000001", expected: 65},
		{input: "01000010", expected: 66},
		{input: "01000011", expected: 67},
		{input: "01000100", expected: 68},
		{input: "01000101", expected: 69},
		{input: "01000110", expected: 70},
		{input: "01000111", expected: 71},
		{input: "01001000", expected: 72},
		{input: "01001001", expected: 73},
		{input: "01001010", expected: 74},
		{input: "01001011", expected: 75},
		{input: "01001100", expected: 76},
		{input: "01001101", expected: 77},
		{input: "01001110", expected: 78},
		{input: "01001111", expected: 79},
		{input: "01010000", expected: 80},
		{input: "01010001", expected: 81},
		{input: "01010010", expected: 82},
		{input: "01010011", expected: 83},
		{input: "01010100", expected: 84},
		{input: "01010101", expected: 85},
		{input: "01010110", expected: 86},
		{input: "01010111", expected: 87},
		{input: "01011000", expected: 88},
		{input: "01011001", expected: 89},
		{input: "01011010", expected: 90},
		{input: "01011011", expected: 91},
		{input: "01011100", expected: 92},
		{input: "01011101", expected: 93},
		{input: "01011110", expected: 94},
		{input: "01011111", expected: 95},
		{input: "01100000", expected: 96},
		{input: "01100001", expected: 97},
		{input: "01100010", expected: 98},
		{input: "01100011", expected: 99},
		{input: "01100100", expected: 100},
		{input: "01100101", expected: 101},
		{input: "01100110", expected: 102},
		{input: "01100111", expected: 103},
		{input: "01101000", expected: 104},
		{input: "01101001", expected: 105},
		{input: "01101010", expected: 106},
		{input: "01101011", expected: 107},
		{input: "01101100", expected: 108},
		{input: "01101101", expected: 109},
		{input: "01101110", expected: 110},
		{input: "01101111", expected: 111},
		{input: "01110000", expected: 112},
		{input: "01110001", expected: 113},
		{input: "01110010", expected: 114},
		{input: "01110011", expected: 115},
		{input: "01110100", expected: 116},
		{input: "01110101", expected: 117},
		{input: "01110110", expected: 118},
		{input: "01110111", expected: 119},
		{input: "01111000", expected: 120},
		{input: "01111001", expected: 121},
		{input: "01111010", expected: 122},
		{input: "01111011", expected: 123},
		{input: "01111100", expected: 124},
		{input: "01111101", expected: 125},
		{input: "01111110", expected: 126},
		{input: "01111111", expected: 127},
		{input: "10000000", expected: 128},
		{input: "10000001", expected: 129},
		{input: "10000010", expected: 130},
		{input: "10000011", expected: 131},
		{input: "10000100", expected: 132},
		{input: "10000101", expected: 133},
		{input: "10000110", expected: 134},
		{input: "10000111", expected: 135},
		{input: "10001000", expected: 136},
		{input: "10001001", expected: 137},
		{input: "10001010", expected: 138},
		{input: "10001011", expected: 139},
		{input: "10001100", expected: 140},
		{input: "10001101", expected: 141},
		{input: "10001110", expected: 142},
		{input: "10001111", expected: 143},
		{input: "10010000", expected: 144},
		{input: "10010001", expected: 145},
		{input: "10010010", expected: 146},
		{input: "10010011", expected: 147},
		{input: "10010100", expected: 148},
		{input: "10010101", expected: 149},
		{input: "10010110", expected: 150},
		{input: "10010111", expected: 151},
		{input: "10011000", expected: 152},
		{input: "10011001", expected: 153},
		{input: "10011010", expected: 154},
		{input: "10011011", expected: 155},
		{input: "10011100", expected: 156},
		{input: "10011101", expected: 157},
		{input: "10011110", expected: 158},
		{input: "10011111", expected: 159},
		{input: "10100000", expected: 160},
		{input: "10100001", expected: 161},
		{input: "10100010", expected: 162},
		{input: "10100011", expected: 163},
		{input: "10100100", expected: 164},
		{input: "10100101", expected: 165},
		{input: "10100110", expected: 166},
		{input: "10100111", expected: 167},
		{input: "10101000", expected: 168},
		{input: "10101001", expected: 169},
		{input: "10101010", expected: 170},
		{input: "10101011", expected: 171},
		{input: "10101100", expected: 172},
		{input: "10101101", expected: 173},
		{input: "10101110", expected: 174},
		{input: "10101111", expected: 175},
		{input: "10110000", expected: 176},
		{input: "10110001", expected: 177},
		{input: "10110010", expected: 178},
		{input: "10110011", expected: 179},
		{input: "10110100", expected: 180},
		{input: "10110101", expected: 181},
		{input: "10110110", expected: 182},
		{input: "10110111", expected: 183},
		{input: "10111000", expected: 184},
		{input: "10111001", expected: 185},
		{input: "10111010", expected: 186},
		{input: "10111011", expected: 187},
		{input: "10111100", expected: 188},
		{input: "10111101", expected: 189},
		{input: "10111110", expected: 190},
		{input: "10111111", expected: 191},
		{input: "11000000", expected: 192},
		{input: "11000001", expected: 193},
		{input: "11000010", expected: 194},
		{input: "11000011", expected: 195},
		{input: "11000100", expected: 196},
		{input: "11000101", expected: 197},
		{input: "11000110", expected: 198},
		{input: "11000111", expected: 199},
		{input: "11001000", expected: 200},
		{input: "11001001", expected: 201},
		{input: "11001010", expected: 202},
		{input: "11001011", expected: 203},
		{input: "11001100", expected: 204},
		{input: "11001101", expected: 205},
		{input: "11001110", expected: 206},
		{input: "11001111", expected: 207},
		{input: "11010000", expected: 208},
		{input: "11010001", expected: 209},
		{input: "11010010", expected: 210},
		{input: "11010011", expected: 211},
		{input: "11010100", expected: 212},
		{input: "11010101", expected: 213},
		{input: "11010110", expected: 214},
		{input: "11010111", expected: 215},
		{input: "11011000", expected: 216},
		{input: "11011001", expected: 217},
		{input: "11011010", expected: 218},
		{input: "11011011", expected: 219},
		{input: "11011100", expected: 220},
		{input: "11011101", expected: 221},
		{input: "11011110", expected: 222},
		{input: "11011111", expected: 223},
		{input: "11100000", expected: 224},
		{input: "11100001", expected: 225},
		{input: "11100010", expected: 226},
		{input: "11100011", expected: 227},
		{input: "11100100", expected: 228},
		{input: "11100101", expected: 229},
		{input: "11100110", expected: 230},
		{input: "11100111", expected: 231},
		{input: "11101000", expected: 232},
		{input: "11101001", expected: 233},
		{input: "11101010", expected: 234},
		{input: "11101011", expected: 235},
		{input: "11101100", expected: 236},
		{input: "11101101", expected: 237},
		{input: "11101110", expected: 238},
		{input: "11101111", expected: 239},
		{input: "11110000", expected: 240},
		{input: "11110001", expected: 241},
		{input: "11110010", expected: 242},
		{input: "11110011", expected: 243},
		{input: "11110100", expected: 244},
		{input: "11110101", expected: 245},
		{input: "11110110", expected: 246},
		{input: "11110111", expected: 247},
		{input: "11111000", expected: 248},
		{input: "11111001", expected: 249},
		{input: "11111010", expected: 250},
		{input: "11111011", expected: 251},
		{input: "11111100", expected: 252},
		{input: "11111101", expected: 253},
		{input: "11111110", expected: 254},
		{input: "11111111", expected: 255},
	} {
		t.Run(td.input, func(t *testing.T) {
			actual := Byte(td.input)
			if actual != td.expected {
				t.Errorf("want %d got %d", td.expected, actual)
			}
		})
	}
}
