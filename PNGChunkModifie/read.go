package PNGChunkModifie

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"os"
)

func ReadData(file *os.File, matchChunk []byte) {

	header := make([]byte, 8)
	_, err := io.ReadFull(file, header)
	if err != nil {
		fmt.Println("Error reading PNG header:", err)
		return
	}

	if string(header) != "\x89PNG\r\n\x1a\n" {
		fmt.Println("Not a valid PNG file.")
		return
	}

	// Read chunks
	for {
		var chunk Chunk

		// Read chunk length
		err := binary.Read(file, binary.BigEndian, &chunk.Length)
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error reading chunk length:", err)
			return
		}

		// Read chunk type
		_, err = io.ReadFull(file, chunk.Type[:])
		if err != nil {
			fmt.Println("Error reading chunk type:", err)
			return
		}

		// Read chunk data
		chunk.Data = make([]byte, chunk.Length)
		_, err = io.ReadFull(file, chunk.Data)
		if err != nil {
			fmt.Println("Error reading chunk data:", err)
			return
		}
		// Read CRC
		err = binary.Read(file, binary.BigEndian, &chunk.CRC)
		if err != nil {
			fmt.Println("Error reading CRC:", err)
			return
		}
		if bytes.Equal(chunk.Type[:], matchChunk) {
			chunk.Data = bytes.Replace(chunk.Data, []byte("\x00"), []byte(":"), 1)
			fmt.Printf("Chunk Type: %s, Data: %s\n", string(chunk.Type[:]), string(chunk.Data[:]))
		}
	}
}
