package types

type Coordinate struct {
	X int
	Y int
}

type Location struct {
	Coordinate  Coordinate
	ColorNumber int
}

type RidgidPiece [][]int
type Grid [200][200]int

type Symmetry struct {
	Mirror      bool
	Rotation    int
	RidgidPiece RidgidPiece
}

type PieceDetails struct {
	RidgidPiece RidgidPiece
	Name        string
	Dimensions  [2]int
	Symmetries  []Symmetry
}
