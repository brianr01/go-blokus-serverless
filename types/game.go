package types

type Coordinate struct {
	X int
	Y int
}

type Move struct {
	Coordinate  Coordinate
	ColorNumber ColorNumber
	RidgidPiece RidgidPiece
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
