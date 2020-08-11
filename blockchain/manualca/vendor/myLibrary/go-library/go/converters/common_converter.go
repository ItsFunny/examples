package converter

import (
	. "encoding/binary"
	"math"
)

func BigEndianUInt642Bytes(i uint64) []byte {
	var buf = make([]byte, 8)
	BigEndian.PutUint64(buf, i)
	return buf
}
func BigEndianInt642Bytes(i int64) []byte {
	var buf = make([]byte, 8)
	BigEndian.PutUint64(buf, uint64(i))
	return buf
}
func LittelEndianInt642Bytes(i int64) []byte {
	var buf = make([]byte, 8)
	LittleEndian.PutUint64(buf, uint64(i))
	return buf
}

func BigEndianBytes2Int64(bytes []byte) uint64 {
	// return uint64(BigEndian.Uint32(bytes))<<16 | uint64(LittleEndian.Uint16(bytes[4:]))
	return uint64(BigEndian.Uint64(bytes))
}
func LittleEndianBytes2Int64(bytes []byte) uint64 {
	// return uint64(BigEndian.Uint32(bytes))<<16 | uint64(LittleEndian.Uint16(bytes[4:]))
	return uint64(LittleEndian.Uint64(bytes))
}

func BigEndianFloat64ToByte(float float64) []byte {
	bits := math.Float64bits(float)
	bytes := make([]byte, 8)
	BigEndian.PutUint64(bytes, bits)

	return bytes
}

func BigEndianBytesToFloat64(bytes []byte) float64 {
	bits := BigEndian.Uint64(bytes)

	return math.Float64frombits(bits)
}
