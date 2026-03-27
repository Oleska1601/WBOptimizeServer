package models

import "time"

type ItemV1 struct {
	// 136 байт
	Active    bool      // 1 + 7 padding
	SKU       string    // 16
	Price     float64   // 8
	Qty       int32     // 4 + 4 padding
	Name      string    // 16
	Weight    float32   // 4 + 4 padding
	Width     int32     // 4 + 4 padding
	Height    int32     // 4 + 4 padding
	Depth     int32     // 4 + 4 padding
	CreatedAt time.Time // 24
	UpdatedAt time.Time // 24
}

type ItemV2 struct {
	CreatedAt time.Time // 24
	UpdatedAt time.Time // 24

	SKU  string // 16
	Name string // 16

	Price float64 // 8

	Qty    int32 // 4
	Width  int32 // 4
	Height int32 // 4
	Depth  int32 // 4

	Weight float32 // 4

	Active bool // 1
	// 112 байт
}
