package types

type Coordinate struct {
	X int
	Y int
}

type Move struct {
	Coordinate  Coordinate
	ColorNumber ColorNumber
	RidgidPiece RidgidPiece
	PieceName   string
}

type ColorNumber int

type RidgidPiece [][]int
type Grid [][]ColorNumber

//	type Availability struct {
//		Empty                bool
//		HasPlayableCorner    bool
//		HasEdgeWithSameColor bool
//	}
type AvailabilityNumber int

type AvailabilityGrid [][]AvailabilityNumber

type ProbabilityGrid [][]float64

type ProbabilityPiece float64

type Symmetry struct {
	Mirror              bool
	Rotation            int
	RidgidPiece         RidgidPiece
	PlayableCoordinates []Coordinate
}

type PieceDetail struct {
	RidgidPiece RidgidPiece
	Name        string
	Dimensions  [2]int
	Symmetries  []Symmetry
}
