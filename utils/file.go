package utils

import (
	"image"
	"image/png"
	"log"
	"os"
	"strings"
)

func GetNameFromFile(n string) string {
	return strings.Split(n, ".")[0]
}

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

func GetImageFromFile(filePath string) image.Image {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Unable to open directory '%s' due to error %s", filePath, err)
	}
	defer file.Close()

	image, err := png.Decode(file)

	if err != nil {
		log.Fatalf("Unable to decode file into image.  Error: %s", err)
	}

	return image
}
