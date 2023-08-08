package utils

import (
	"log"
	"math"

	"github.com/brianr01/go-blockus-serverless/constants"
	"github.com/brianr01/go-blockus-serverless/types"
)

// func IsPiecePlayableAtCoordinate(x, y, types.RidgidPiece)
// func GetPlayableCoordinatesForColor(clr types.ColorNumber, g types.Grid) []types.Coordinate {
// 	return GetCoordinatesWithMinAvailabilityNumber(constants.AvailabilityNumberPlayable, clr, g)
// }

func placePiece(clr types.ColorNumber, rp types.RidgidPiece, c types.Coordinate, g types.Grid) types.Grid {
	for _, cCoord := range GetColorCoordinatesForRidgidPiece(rp) {
		updateCoord := AddCoordinates(cCoord, c)

		g = ColorLocation(clr, updateCoord, g)
	}

	return g
}

func ColorLocation(clr types.ColorNumber, c types.Coordinate, g types.Grid) types.Grid {
	g[c.X][c.Y] = clr

	return g
}

func IsPlayableCoordinate(c types.Coordinate, clr types.ColorNumber, g types.Grid) bool {
	// Location is populated already.
	if !IsCoordinateOnGridEmpty(c, g) {
		return false
	}

	// Edges to coordinate has same color.
	if DoesEdgesToCoordnateHaveColor(c, clr, g) {
		return false
	}

	// None of the corners match the color
	if !DoesCornersToCoordnateHaveColor(c, clr, g) {
		return false
	}

	// The location is playable.
	return true
}

func DoesCornersToCoordnateHaveColor(c types.Coordinate, clr types.ColorNumber, g types.Grid) bool {
	return CoordinatesContainColor(GetCornerCoordinatesFromCoordinate(c, g), clr, g)
}

func DoesEdgesToCoordnateHaveColor(c types.Coordinate, clr types.ColorNumber, g types.Grid) bool {
	return CoordinatesContainColor(GetEdgeCoordinatesFromCoordinate(c, g), clr, g)
}

func CoordinatesContainColor(cs []types.Coordinate, clr types.ColorNumber, g types.Grid) bool {
	return ColorNumberInColorNumbers(clr, GetColorNumbersFromCoordinates(cs, g))
}

func IsCoordinateOnGridEmpty(c types.Coordinate, g types.Grid) bool {
	return DoesCoordinateHaveColorNumber(constants.ColorNumberEmpty, c, g)
}

func DoesCoordinateHaveColorNumber(clr types.ColorNumber, c types.Coordinate, g types.Grid) bool {
	return GetColorNumberFromCoordinate(c, g) == clr
}

func GetColorNumbersFromCoordinates(cs []types.Coordinate, g types.Grid) []types.ColorNumber {
	var colorNumbers []types.ColorNumber
	for _, c := range cs {
		colorNumbers = append(colorNumbers, GetColorNumberFromCoordinate(c, g))
	}

	return colorNumbers
}

func GetColorNumberFromCoordinate(c types.Coordinate, g types.Grid) types.ColorNumber {
	x, y := c.X, c.Y
	return g[x][y]
}

func GetCornerCoordinatesForCoordinates(cs []types.Coordinate, g types.Grid) []types.Coordinate {
	csCorners := make([]types.Coordinate, 0)

	for _, c := range cs {
		csCorners = append(csCorners, GetCornerCoordinatesFromCoordinate(c, g)...)
	}

	return csCorners
}

func GetEdgeCoordinatesFromCoordinate(coord types.Coordinate, g types.Grid) []types.Coordinate {
	var edgeCoords []types.Coordinate

	// Get edges in x and y direction.
	for _, i := range []int{-1, 1} {
		edgeCoords = append(edgeCoords, types.Coordinate{
			X: coord.X + i,
			Y: coord.Y + 0,
		})

		edgeCoords = append(edgeCoords, types.Coordinate{
			X: coord.X + 0,
			Y: coord.Y + i,
		})
	}

	return FilterCoordinatesByValidGridLocation(edgeCoords, g)
}

func FilterCoordinatesByValidGridLocation(cs []types.Coordinate, g types.Grid) []types.Coordinate {
	var r []types.Coordinate
	for _, c := range cs {
		if IsCoordinateOnGrid(c, g) {
			r = append(r, c)
		}
	}

	return r
}

func GetCornerCoordinatesFromCoordinate(coord types.Coordinate, g types.Grid) []types.Coordinate {
	var coords []types.Coordinate
	for _, x := range []int{-1, 1} {
		for _, y := range []int{-1, 1} {
			coords = append(coords, types.Coordinate{
				X: coord.X + x,
				Y: coord.Y + y,
			})
		}
	}

	return FilterCoordinatesByValidGridLocation(coords, g)
}

func AreCoordinatesOnGrid(cs []types.Coordinate, g types.Grid) bool {
	for _, c := range cs {
		if !IsCoordinateOnGrid(c, g) {
			return false
		}
	}

	return true
}

func IsCoordinateOnGrid(c types.Coordinate, g types.Grid) bool {
	x, y := c.X, c.Y

	gridSize := GetGridSize(g)
	for i, dir := range [2]int{x, y} {
		// Off the grid in the negative direction.
		if dir < 0 {
			return false
		}
		// Off the grid in the positive direction.
		if gridSize[i] <= dir {
			return false
		}
	}

	return true
}

func SwitchColorNumbers(m map[types.ColorNumber]types.ColorNumber, g types.Grid) types.Grid {
	size := GetGridSize(g)
	w, h := size[0], size[1]
	gNew := CreateEmptyGrid(w, h)

	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {

			mappedClr, found := m[g[i][j]]

			if found {
				gNew[i][j] = mappedClr
				continue
			}
			gNew[i][j] = g[i][j]
		}
	}

	return gNew
}

func RotateGrid(r int, g types.Grid) types.Grid {
	if !IsRotationValidFor90Degrees(r) {
		log.Fatalf("Cannot grid.  Invalid rotation %v", r)
	}

	rotateClockWise := r >= 0

	rotations := int(math.Abs(float64(r / 90)))

	ints2d := convertGridTo2dInts(g)

	for i := 0; i < rotations; i++ {
		if rotateClockWise {
			ints2d = ReverseColumns2d(Transpose2d(ints2d))
		} else {
			ints2d = ReverseRows2d(Transpose2d(ints2d))
		}
	}

	g = convert2dIntsToGrid(ints2d)

	return g
}

// func convertGridToHtmlString(g types.Grid) {

// }

func convert2dIntsToGrid(ints2d [][]int) types.Grid {
	w, h := len(ints2d), len(ints2d[0])
	g := CreateEmptyGrid(w, h)

	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			g[i][j] = types.ColorNumber(ints2d[i][j])
		}
	}

	return g
}

func convertGridTo2dInts(g types.Grid) [][]int {
	ints2d := make([][]int, len(g))

	for i, _ := range ints2d {
		row := make([]int, len(g[0]))
		ints2d[i] = row

		for j, _ := range row {
			ints2d[i][j] = int(g[i][j])
		}
	}

	return ints2d
}

func GetGridSize(g types.Grid) [2]int {
	return [2]int{len(g), len(g[0])}
}

func CreateEmptyGrid(w int, h int) types.Grid {
	var g types.Grid

	for i := 0; i < w; i++ {
		var row []types.ColorNumber
		for j := 0; j < h; j++ {
			row = append(row, constants.ColorNumberEmpty)
		}
		g = append(g, row)
	}

	return g
}
