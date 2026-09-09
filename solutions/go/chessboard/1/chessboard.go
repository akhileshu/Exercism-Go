package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
/*
Each File ( "A" through "H" ) can hold 8 booleans, where for example:

false → _
true  → #
*/
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

func Count[T comparable](items []T, target T) int {
	count := 0

	for _, item := range items {
		if item == target {
			count++
		}
	}

	return count
}

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	return Count(cb[file], true)
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
// CountInRank checks the given rank (1–8) across all 8 files (A–H).
//
// Think of the board as 8 vertical files, each containing 8 squares:
//
//	A  B  C  D  E  F  G  H
//
// Rank 8  _  #  _  _  _  #  _  _
// Rank 7  #  _  _  _  #  _  _  _
// Rank 6  _  _  #  _  _  _  _  _
// ...
// Rank 2  #  _  #  _  _  #  _  _
// Rank 1  _  _  _  #  _  _  _  _
//
// `rank` is 1-based, while a Go slice is 0-based,
// so rank 1 → index 0, rank 2 → index 1, ..., rank 8 → index 7.
//
// We loop through every file, look at the square at that rank,
// and count it if it is occupied (`true`).
func CountInRank(cb Chessboard, rank int) int {
	if rank < 1 || rank > 8 {
		return 0
	}
	count := 0
	for _, file := range cb {
		if file[rank-1] {
			count++
		}

	}
	return count
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	return len(cb) * len(cb)
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	count := 0
	for _, file := range cb {
		count += Count(file, true)

	}
	return count
}
