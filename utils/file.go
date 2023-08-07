package utils

import (
	"image"
	"image/png"
	"log"
	"os"
	"strings"
)

// Separates the name from the extension.
func GetNameFromFile(n string) string {
	return strings.Split(n, ".")[0]
}

// Lists a directory's files given a path.
func ListDirectory(dirName string) []string {
	file, err := os.Open(dirName)

	if err != nil {
		log.Fatalf("Unable to open directory '%s' due to error %s", dirName, err)
	}

	defer file.Close()

	list, err := file.Readdirnames(0)
	if err != nil {
		log.Fatalf("Unable to read directory names due to error: %s", err)
	}

	return list
}

func GetPngImageFromFile(filePath string) image.Image {
	file, err := os.Open(filePath)
	defer file.Close()

	if err != nil {
		log.Fatalf("Unable to open directory '%s' due to error %s", filePath, err)
	}

	image, err := png.Decode(file)

	if err != nil {
		log.Fatalf("Unable to decode file into image.  Error: %s", err)
	}

	return image
}
