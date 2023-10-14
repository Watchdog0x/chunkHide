package PNGChunkModifie

import (
	"encoding/binary"
	"fmt"
	"hash/crc32"
	"io"
	"os"
)



func Validate(file *os.File) {

	// Read and validate PNG header
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

		// Verify CRC
		crc := crc32.NewIEEE()
		crc.Write(chunk.Type[:])
		crc.Write(chunk.Data)
		calculatedCRC := crc.Sum32()

		if calculatedCRC != chunk.CRC {
			fmt.Println("CRC mismatch for chunk:", string(chunk.Type[:]))
			return
		}

		// Print or process the chunk data as needed
		fmt.Printf("Chunk Type: %s, Length: %d\n", string(chunk.Type[:]), chunk.Length)
	}
}
