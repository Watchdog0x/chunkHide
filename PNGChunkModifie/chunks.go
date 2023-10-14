package PNGChunkModifie

import (
	"bytes"
	"compress/zlib"
	"encoding/binary"
	"fmt"
	"io"
	"os"
)

type Chunk struct {
	Length uint32
	Type   [4]byte
	Data   []byte
	CRC    uint32
}

func NewtEXt(keyword, data string) []byte {

	const tEXtChunkType = "tEXt"

	// Construct the tEXt chunk
	chunkData := append([]byte(keyword), append([]byte{0}, []byte(data)...)...)

	lengthOfData := uint32ToBytes(uint32(len(chunkData)))

	return append(append(append(lengthOfData, []byte(tEXtChunkType)...), chunkData...), calculateCRC([]byte(tEXtChunkType), chunkData)...)

}

func NewZTXt(keyword, data string) []byte {

	const zTXtChunkType = "zTXt"

	// Compress the text data using zlib
	var compressedData bytes.Buffer
	writer := zlib.NewWriter(&compressedData)
	_, err := writer.Write([]byte(data))
	if err != nil {
		fmt.Println("Error compressing the data:", err)
		os.Exit(1)
	}
	writer.Close()

	// Construct the zTXt chunk
	chunkData := append([]byte(keyword), 0)
	chunkData = append(chunkData, 0)
	chunkData = append(chunkData, compressedData.Bytes()...)

	lengthOfData := uint32ToBytes(uint32(len(chunkData)))

	return append(append(append(lengthOfData, []byte(zTXtChunkType)...), chunkData...), calculateCRC([]byte(zTXtChunkType), chunkData)...)
}

func NewPLTE(data string) []byte {

	const PLTEChunkType = "PLTE"

	// Construct the PLTE chunk
	for len(data)%3 != 0 {
		data += " "
	}
	chunkData := []byte(data)

	lengthOfData := uint32ToBytes(uint32(len(chunkData)))

	return append(append(append(lengthOfData, []byte(PLTEChunkType)...), chunkData...), calculateCRC([]byte(PLTEChunkType), chunkData)...)
}

func Constructor(file *os.File, newChunk, addBeforeChunk []byte) []byte {

	fileData, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading PNG", err)
	}

	position := bytes.Index(fileData, addBeforeChunk)
	if position == -1 {
		fmt.Println(addBeforeChunk, "not found in the png.")
		os.Exit(1)
	}

	return append(fileData[:position-4], append(newChunk, fileData[position-4:]...)...)
}

func ReadChunkData(reader io.Reader) ([]byte, error) {
	// Read the length field (4 bytes)
	lengthBytes := make([]byte, 4)
	_, err := reader.Read(lengthBytes)
	if err != nil {
		return nil, err
	}

	// Convert lengthBytes to an integer
	length := binary.BigEndian.Uint32(lengthBytes)

	// Read the chunk type (4 bytes)
	chunkTypeBytes := make([]byte, 4)
	_, err = reader.Read(chunkTypeBytes)
	if err != nil {
		return nil, err
	}

	// Read the data field based on the length
	data := make([]byte, length)
	_, err = reader.Read(data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
