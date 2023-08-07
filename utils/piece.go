package utils

import (
	"fmt"
	"image"
	"log"
	"math"
	"reflect"

	"github.com/brianr01/go-blockus-serverless/types"
)

func CreateAllPieceDetails(folderLocation string) []types.PieceDetails {
	var details []types.PieceDetails
	for _, fileName := range ListDirectory(folderLocation) {
		details = append(details, CreatePieceDetails(fileName, folderLocation))
	}

	return details
}

func CreatePieceDetails(fileName string, folderLocation string) types.PieceDetails {
	name := GetNameFromFile(fileName)

	image := GetPngImageFromFile(fmt.Sprintf("%s/%s", folderLocation, fileName))

	ridgidPiece := CreateRidigPieceFromImage(image)
	dimensions := GetDimensionsFromRidigPiece(ridgidPiece)
	symmetries := GetValidSymmetriesFromRidgidPiece(ridgidPiece)

	pieceDetails := types.PieceDetails{
		RidgidPiece: ridgidPiece,
		Name:        name,
		Dimensions:  dimensions,
		Symmetries:  symmetries,
	}

	return pieceDetails
}

func GetValidSymmetriesFromRidgidPiece(p types.RidgidPiece) []types.Symmetry {
	var symmetries []types.Symmetry
	var validRidigPieces []types.RidgidPiece

	// Make all 8 possible symmetries. Two at a time.
	for r := 0; r < 4; r++ { // 4 rotations
		for m := 0; m < 2; m++ { // 2 mirrors
			mirrored := m == 1

			rp := RotateRidgidPiece(p, r*90)
			if mirrored {
				rp = ReverseRows2d(rp)
			}

			if !RidgidPieceInRidgidPieces(rp, validRidigPieces) {
				validRidigPieces = append(validRidigPieces, rp)
				symmetries = append(symmetries, types.Symmetry{
					Mirror:      mirrored,
					Rotation:    r,
					RidgidPiece: rp,
				})
			}
		}
	}

	return symmetries
}

func CreateRidigPieceFromImage(img image.Image) types.RidgidPiece {
	bX := img.Bounds().Max.X
	bY := img.Bounds().Max.Y

	p := make(types.RidgidPiece, bX)

	for x := 0; x < bX; x++ {
		p[x] = make([]int, bY)

		for y := 0; y < bY; y++ {
			if IsPixelWhite(x, y, img) {
				p[x][y] = 0
				continue
			}

			p[x][y] = 1
		}
	}

	return p
}

func GetDimensionsFromRidigPiece(p types.RidgidPiece) [2]int {
	w := len(p)

	if w == 0 {
		log.Fatalf("Cannot get dimensions fom ridig piece that has no width.")
	}

	h := len(p[0])

	return [2]int{w, h}
}

func RotateRidgidPiece(p types.RidgidPiece, r int) types.RidgidPiece {
	if !IsRotationValidFor90Degrees(r) {
		log.Fatalf("Cannot rotate ridgid piece.  Invalid rotation %v", r)
	}

	rotateClockWise := r >= 0

	rotations := int(math.Abs(float64(r / 90)))

	if rotateClockWise {
		for i := 0; i < rotations; i++ {
			p = RotateRidigidPieceClockwise(p)
		}

		return p
	}

	for i := 0; i < rotations; i++ {
		p = RotateRidigidPieceCounterClockwise(p)
	}

	return p
}

func RotateRidigidPieceCounterClockwise(p types.RidgidPiece) types.RidgidPiece {
	return ReverseRows2d(Transpose2d(p))
}

func RotateRidigidPieceClockwise(p types.RidgidPiece) types.RidgidPiece {
	return ReverseColumns2d(Transpose2d(p))
}

func RidgidPieceInRidgidPieces(p1 types.RidgidPiece, ps []types.RidgidPiece) bool {
	for _, p2 := range ps {
		if reflect.DeepEqual(p1, p2) {
			return true
		}
	}

	return false
}
