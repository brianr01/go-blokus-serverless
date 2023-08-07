package utils

import (
	"fmt"

	"github.com/brianr01/go-blockus-serverless/helpers"
	"github.com/brianr01/go-blockus-serverless/types"
)

func CreateAllPieceDetails(folderLocation string) []types.PieceDetails {
	var details []types.PieceDetails
	for _, fileName := range helpers.ListDirectory(folderLocation) {
		details = append(details, CreatePieceDetails(fileName, folderLocation))
	}

	return details
}

func CreatePieceDetails(fileName string, folderLocation string) types.PieceDetails {
	name := helpers.GetNameFromFile(fileName)

	image := helpers.GetImageFromFile(fmt.Sprintf("%s/%s", folderLocation, fileName))

	ridgidPiece := helpers.CreateRidigPieceFromImage(image)
	dimensions := helpers.GetDimensionsFromRidigPiece(ridgidPiece)
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

			rp := helpers.RotateRidgidPiece(p, r*90)
			if mirrored {
				rp = helpers.ReverseRows2d(rp)
			}

			if !helpers.RidgidPieceInRidgidPieces(rp, validRidigPieces) {
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
