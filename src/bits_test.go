package efilfoemag

import (
	"fmt"
	"testing"
)

func TestSumbit(t *testing.T) {
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
			actual := Sumbit(td.number)
			if actual != td.expected {
				t.Errorf("excpected %d, got %d", td.expected, actual)
			}
		})
	}
}
