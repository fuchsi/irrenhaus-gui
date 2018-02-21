package main

import "github.com/mattn/go-gtk/gdkpixbuf"

var (
	logoPng = gdkpixbuf.PixbufData{Data: []uint8{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xe8, 0xc9, 0xb6, 0xff, 0xfd, 0xeb, 0xce, 0xff, 0xfc, 0xe8, 0xc9, 0xff, 0xfe, 0xfe, 0xfe, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xb1, 0x88, 0x7a, 0xff, 0xfd, 0xea, 0xce, 0xff, 0xfe, 0xee, 0xd8, 0xff, 0xfe, 0xe9, 0xd0, 0xff, 0xfe, 0xfe, 0xfe, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xfe, 0xfe, 0xff, 0xff, 0x75, 0x6a, 0x77, 0xff, 0xe0, 0xcb, 0xbf, 0xff, 0x3d, 0x3f, 0x60, 0xff, 0xfd, 0xfe, 0xfe, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xfe, 0xfe, 0xfe, 0xff, 0x3c, 0x2c, 0x41, 0xff, 0xfd, 0xea, 0xd2, 0xff, 0x54, 0x55, 0x75, 0xff, 0xa2, 0xa3, 0xaa, 0xff, 0xb1, 0x98, 0x74, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xff, 0xff, 0xff, 0xff, 0xf8, 0xf4, 0xef, 0xff, 0x70, 0x2c, 0x29, 0xff, 0x2c, 0x3, 0x8, 0xff, 0xf, 0x4, 0x3, 0xff, 0x9c, 0x5e, 0x45, 0xff, 0x24, 0x8, 0x8, 0xff, 0x33, 0x1c, 0x19, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xa7, 0x79, 0x63, 0xff, 0x7c, 0x1f, 0x1d, 0xff, 0xdb, 0x92, 0x60, 0xff, 0xd2, 0x84, 0x5e, 0xff, 0xe1, 0x85, 0x54, 0xff, 0xc6, 0xab, 0x85, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x76, 0x5b, 0x5d, 0xff, 0x43, 0xe, 0x18, 0xff, 0xfc, 0xd3, 0xa0, 0xff, 0xd6, 0x9e, 0x6c, 0xff, 0xe2, 0xba, 0x78, 0xff, 0x0, 0x0, 0x0, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xff, 0xff, 0xff, 0xff, 0xe, 0x1, 0xd, 0xff, 0xd, 0x1, 0x4, 0xff, 0x51, 0x34, 0x38, 0xff, 0xe8, 0xc7, 0x9d, 0xff, 0xdf, 0x93, 0x61, 0xff, 0x6a, 0x53, 0x42, 0xff, 0x47, 0xa, 0x9, 0xff, 0x4c, 0x4c, 0x70, 0xff, 0x7, 0x8, 0x2c, 0xff, 0xb, 0xd, 0x29, 0xff, 0xff, 0xff, 0xff, 0xff, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xfa, 0xfb, 0xfa, 0xff, 0x5, 0x1, 0x13, 0xff, 0x9, 0x0, 0x2, 0xff, 0x3f, 0x22, 0x24, 0xff, 0xc6, 0x76, 0x4d, 0xff, 0xf2, 0xa2, 0x6a, 0xff, 0x94, 0x17, 0x15, 0xff, 0x16, 0xc, 0x28, 0xff, 0x2, 0x2, 0x21, 0xff, 0x46, 0x48, 0x60, 0xff, 0x2, 0x4, 0x22, 0xff, 0xff, 0xff, 0xff, 0xff, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xfc, 0xfc, 0xfc, 0xff, 0x11, 0x0, 0x10, 0xff, 0xb, 0x1, 0x5, 0xff, 0xb5, 0x72, 0x69, 0xff, 0xaf, 0x51, 0x43, 0xff, 0xab, 0x62, 0x52, 0xff, 0xfd, 0xef, 0xd7, 0xff, 0x34, 0x27, 0x22, 0xff, 0x72, 0x73, 0x8d, 0xff, 0x12, 0x18, 0x33, 0xff, 0x1, 0x0, 0x6, 0xff, 0x2c, 0x23, 0x1d, 0xff, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x84, 0x74, 0x7e, 0xff, 0x15, 0xa, 0x25, 0xff, 0xd7, 0xbf, 0xb0, 0xff, 0x87, 0x3b, 0x3b, 0xff, 0xe, 0xa, 0xf, 0xff, 0xf8, 0xdf, 0xc3, 0xff, 0xfa, 0xe7, 0xc2, 0xff, 0x13, 0xd, 0x10, 0xff, 0x5, 0xb, 0x31, 0xff, 0x1, 0x9, 0x20, 0xff, 0x1, 0x0, 0x0, 0xff, 0xdf, 0xdd, 0xda, 0xff, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x3b, 0x24, 0x3a, 0xff, 0x38, 0x2e, 0x4e, 0xff, 0xfd, 0xec, 0xd8, 0xff, 0xce, 0x90, 0x78, 0xff, 0x99, 0x85, 0x6e, 0xff, 0xed, 0xd8, 0xbc, 0xff, 0xf7, 0xe2, 0xb2, 0xff, 0x3, 0x3, 0xb, 0xff, 0x44, 0x4f, 0x6b, 0xff, 0x1e, 0x2f, 0x59, 0xff, 0xd0, 0xd3, 0xdc, 0xff, 0x64, 0x6f, 0x86, 0xff, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x24, 0x15, 0x28, 0xff, 0x11, 0x4, 0x16, 0xff, 0x3d, 0x27, 0x2a, 0xff, 0x87, 0x54, 0x4f, 0xff, 0x4, 0x1, 0x12, 0xff, 0x7d, 0x7f, 0x83, 0xff, 0xfe, 0xec, 0xd8, 0xff, 0x7, 0x5, 0x10, 0xff, 0x0, 0x0, 0x15, 0xff, 0xa7, 0xa9, 0xb5, 0xff, 0xfd, 0xfd, 0xfe, 0xff, 0xff, 0xff, 0xff, 0xff, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x13, 0xd, 0x13, 0xff, 0x0, 0x0, 0x0, 0xff, 0x0, 0x0, 0x0, 0xff, 0x0, 0x0, 0x1, 0xff, 0x0, 0x0, 0x0, 0xff, 0x0, 0x0, 0x0, 0xff, 0xb4, 0xa4, 0x86, 0xff, 0xba, 0xa9, 0x98, 0xff, 0x2, 0x4, 0xf, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xff, 0xff, 0xff, 0xff, 0x1e, 0x19, 0x35, 0xff, 0xc, 0x2, 0x13, 0xff, 0x1, 0x1, 0x2, 0xff, 0xa, 0x14, 0x3a, 0xff, 0x0, 0x0, 0x0, 0xff, 0xf8, 0xe3, 0xcf, 0xff, 0xfc, 0xf7, 0xef, 0xff, 0x56, 0x54, 0x6d, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x82, 0x82, 0x82, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, Colorspace: 0, HasAlpha: true, BitsPerSample: 8, Width: 16, Height: 16, RowStride: 64}
)
