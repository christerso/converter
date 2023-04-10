package main

import (
	"flag"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
)

func main() {
	targetDir := flag.String("target", ".", "target directory to convert")
	outputFormat := flag.String("format", "jpg", "output image format (jpg or png)")
	flag.Parse()

	// Loop through all image files in the target directory
	count := 0
	err := filepath.Walk(*targetDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Only convert files with .jpg, .jpeg, and .png extensions
		ext := filepath.Ext(path)
		if ext == ".jpg" || ext == ".jpeg" || ext == ".png" {
			// Check if output format is the same as input format
			if (ext == ".jpg" || ext == ".jpeg") && *outputFormat == "jpg" || ext == ".png" && *outputFormat == "png" {
				return nil
			}

			// Open the input image file
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()

			// Decode the input image to an image.Image object
			img, _, err := image.Decode(file)
			if err != nil {
				return err
			}

			// Determine the output file name and format
			outFile := filepath.Join(*targetDir, info.Name())
			if *outputFormat == "png" {
				outFile = outFile[:len(outFile)-len(ext)] + ".png"
			} else {
				outFile = outFile[:len(outFile)-len(ext)] + ".jpg"
			}

			// Create the output file
			out, err := os.Create(outFile)
			if err != nil {
				return err
			}
			defer out.Close()

			// Encode the image in the desired output format
			if *outputFormat == "png" {
				// Set the PNG encoder to use maximum compression level
				pngEnc := png.Encoder{CompressionLevel: png.BestCompression}
				err = pngEnc.Encode(out, img)
			} else {
				// Set the JPEG encoder to use maximum quality
				err = jpeg.Encode(out, img, &jpeg.Options{Quality: 100})
			}
			if err != nil {
				return err
			}

			// Print out the filenames of the original and converted files
			fmt.Printf("Converted %s to %s\n", path, outFile)

			// Increment the file count
			count++
		}

		return nil
	})

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Print out the total number of files converted
	fmt.Printf("Converted %d files\n", count)
}
