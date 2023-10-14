package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/Watchdog0x/chunkHide/PNGChunkModifie"
)

func main() {
	const version = "1.0.0"

	var (
		imagePath string
		chunktype string
		keyword   string
		text      string
		output    string
		validate  bool
		read      bool
	)

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "By Watchdog0x Version %s \n\n", version)
		fmt.Fprintf(os.Stderr, "Usage: %s [options]\n", os.Args[0])
		fmt.Fprintln(os.Stderr, "Options:")
		fmt.Fprintln(os.Stderr, "  -i [image.png]\t Path to the PNG image file.")
		fmt.Fprintln(os.Stderr, "  -t [tEXt]\t\t Specify chunk types:")
		fmt.Fprintln(os.Stderr, "\t\t\t   tEXt (text chunk)")
		fmt.Fprintln(os.Stderr, "\t\t\t   zTXt (compressed text chunk)")
		fmt.Fprintln(os.Stderr, "\t\t\t   PLTE (palette chunk)")
		fmt.Fprintln(os.Stderr, "  -o [output.png]\t Where to output the modified image. Default: output.png")
		fmt.Fprintln(os.Stderr, "  -v [image.png]\t Validate the image chunks.")
		fmt.Fprintln(os.Stderr, "  -r\t\t\t Read the chunk data. Requires -t option.")
		fmt.Fprintln(os.Stderr, "  -keyword\t\t Keyword for the chunk. Required with -t tEXt or -t zTXt.")
		fmt.Fprintln(os.Stderr, "  -text\t\t\t Text for the chunk. Required with -t tEXt, -t zTXt or -t PLTE.")
		fmt.Println("\nExamples:")
		fmt.Println("  -t tEXt -keyword mykey -text Hello -i input.png -o output.png")
		fmt.Println("  -t zTXt -keyword mykey -text Hello -i input.png -o output.png")
		fmt.Println("  -r -t tEXt -i input.png")
		fmt.Println("  -v input.png")
	}

	flag.StringVar(&imagePath, "i", "", "Path to the PNG image file")
	flag.StringVar(&output, "o", "output.png", "Output")
	flag.StringVar(&chunktype, "t", "", "Chunk type")
	flag.BoolVar(&read, "r", false, "Read the chunk data. Requires -t option.")
	flag.BoolVar(&validate, "v", false, "Validate PNG. Example: -v image.png")
	flag.StringVar(&keyword, "keyword", "", "Keyword for the chunk. Required with -t tEXt or -t zTXt.")
	flag.StringVar(&text, "text", "", "Text for the chunk. Required with -t tEXt, -t zTXt or -t PLTE.")
	flag.Parse()

	if imagePath != "" {
		chunks := strings.ToLower(chunktype)

		file, err := os.Open(imagePath)
		if err != nil {
			fmt.Println("Error: Unable to open the file:", err)
			return
		}
		defer file.Close()

		if validate {
			PNGChunkModifie.Validate(file)
		} else if read {
			if chunks != "" {
				switch chunks {
				case "text":
					PNGChunkModifie.ReadData(file, []byte("tEXt"))
				case "ztxt":
					PNGChunkModifie.ReadData(file, []byte("zTXt"))
				case "plte":
					PNGChunkModifie.ReadData(file, []byte("PLTE"))
				default:
					fmt.Println("Error: Unsupported chunk type.")
					os.Exit(2)
				}
			} else {
				fmt.Println("Error: You need to specify a chunk type with -t.")
				os.Exit(2)
			}
		} else if chunktype != "" {
			switch chunks {
			case "text":
				if keyword != "" && text != "" {
					imageData := PNGChunkModifie.Constructor(file, PNGChunkModifie.NewtEXt(keyword, text), []byte("IDAT"))
					os.WriteFile(output, imageData, 0644)
				} else {
					fmt.Println("Error: You need to specify both -keyword and -text.")
					os.Exit(2)
				}
			case "ztxt":
				if keyword != "" && text != "" {
					imageData := PNGChunkModifie.Constructor(file, PNGChunkModifie.NewZTXt(keyword, text), []byte("IDAT"))
					os.WriteFile(output, imageData, 0644)
				} else {
					fmt.Println("Error: You need to specify both -keyword and -text.")
					os.Exit(2)
				}
			case "plte":
				if text != "" {
					imageData := PNGChunkModifie.Constructor(file, PNGChunkModifie.NewPLTE(text), []byte("IDAT"))
					os.WriteFile(output, imageData, 0644)
				} else {
					fmt.Println("Error: You need to specify -text.")
					os.Exit(2)
				}
			default:
				fmt.Println("Error: Unsupported chunk type.")
				os.Exit(2)
			}
		} else {
			fmt.Println("Error: You need to use -v, -r, or -t with the -i flag.")
			os.Exit(2)
		}

	} else {
		fmt.Println("Error: Please provide an image path using the -i flag.")
		flag.Usage()
		os.Exit(2)
	}

}
