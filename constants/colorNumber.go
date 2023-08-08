package constants

import "github.com/brianr01/go-blockus-serverless/types"

const ColorNumberEmpty types.ColorNumber = 0
const ColorNumberGeneric types.ColorNumber = 1
const ColorNumberRed types.ColorNumber = 1
const ColorNumberGreen types.ColorNumber = 2
const ColorNumberBlue types.ColorNumber = 3
const ColorNumberYellow types.ColorNumber = 4

const AvailabilityNumberOccupied types.AvailabilityNumber = 0 // The location already has a piece there.
const AvailabilityNumberEmpty types.AvailabilityNumber = 1    // The location is empty.
const AvailabilityNumberValid types.AvailabilityNumber = 2    // The location is a valid place to put a color, but not start a piece.
const AvailabilityNumberPlayable types.AvailabilityNumber = 3 // The location is playble (can start a piece).
