package utils

import (
	"fmt"
	"log"
	"strings"
)

func IsRotationValidFor90Degrees(d int) bool {
	return d%90 == 0
}

func ReverseRows2d(original [][]int) [][]int {
	xMax, yMax := getSize2d(original)

	// Make the result matrix the same size as the original.
	result := make([][]int, xMax)
	for i := range result {
		result[i] = make([]int, yMax)
	}

	// Put [i, j] in the result using [iMax - i, j] from the original matrix.
	for i := 0; i < xMax; i++ {
		for j := 0; j < yMax; j++ {
			result[i][j] = original[xMax-1-i][j]
		}
	}

	return result
}

func ReverseColumns2d(original [][]int) [][]int {
	xMax, yMax := getSize2d(original)

	// Make the result matrix the same size as the original.
	result := make([][]int, xMax)
	for i := range result {
		result[i] = make([]int, yMax)
	}

	// Put [i,j] in the result using [i, jMax -j] from the original matrix.
	for i := 0; i < xMax; i++ {
		for j := 0; j < yMax; j++ {
			result[i][j] = original[i][yMax-1-j]
		}
	}

	return result
}

func Transpose2d(original [][]int) [][]int {
	xMax, yMax := getSize2d(original)

	// Make the output matrix size.
	result := make([][]int, yMax)
	for i := range result {
		result[i] = make([]int, xMax)
	}

	// Put [j,i] in the result using [i, j] from the original matrix.
	for i := 0; i < xMax; i++ {
		for j := 0; j < yMax; j++ {
			result[j][i] = original[i][j]
		}
	}

	return result
}

func getSize2d(array2d [][]int) (int, int) {
	xMax := len(array2d)

	if xMax == 0 {
		log.Fatal("Cannot get size of 2d array with no rows.")
	}

	yMax := len(array2d[0])

	return xMax, yMax
}

func GetStringFrom2d(slice2d [][]int) string {

	rows := make([]string, len(slice2d))
	for i, slice1d := range slice2d {
		rows[i] = fmt.Sprint(slice1d)
	}

	result := fmt.Sprintf("[\n    %s\n]\n", strings.Join(rows, "\n    "))

	return result
}
