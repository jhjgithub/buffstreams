package buffstreams

import (
	"encoding/binary"
	//"log"
	//"math"
)

func byteArrayToUInt32(bytes []byte) (result int64, bytesRead int) {
	return binary.Varint(bytes)
}

func intToByteArray(value int64, bufferSize int) []byte {
	toWriteLen := make([]byte, bufferSize)
	binary.PutVarint(toWriteLen, value)

	return toWriteLen
}

func byteArrayToInt64(bytes []byte) (result int64, bytesRead int) {
	// network order
	size := int64(binary.BigEndian.Uint64(bytes))

	return size, 8
}

func int64ToByteArray(value int64, bufferSize int) []byte {
	toWriteLen := make([]byte, bufferSize)
	// network order
	binary.BigEndian.PutUint64(toWriteLen, uint64(value))

	return toWriteLen
}

// Formula for taking size in bytes and calculating # of bits to express that size
// http://www.exploringbinary.com/number-of-bits-in-a-decimal-integer/
func messageSizeToBitLength(messageSize int) int {
	/*
		bytes := float64(messageSize)
		header := math.Ceil(math.Floor(math.Log2(bytes)+1)/8.0) + 1

		return int(header)
	*/

	// fixed int64
	return 8
}
