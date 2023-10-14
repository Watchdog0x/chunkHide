package PNGChunkModifie

import (
	"encoding/binary"
	"hash/crc32"
)

func uint32ToBytes(value uint32) []byte {
	bytes := make([]byte, 4)
	binary.BigEndian.PutUint32(bytes, value)
	return bytes
}

func calculateCRC(chunkType, chunkData []byte) []byte {
	crc := crc32.NewIEEE()

	crc.Write(chunkType)
	crc.Write(chunkData)

	calculatedCRC := crc.Sum32()

	CRC := uint32ToBytes(calculatedCRC)

	return CRC
}
