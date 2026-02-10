package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools

type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"

type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	f, ok := cb[file]
    if !ok {
        return 0
    }
    count := 0;
    for _, occupied := range f {
        if occupied {
            count++
        }
    }
    return count
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
    if rank <1 || rank >8 {
        return 0
    }
    files := []string{"A", "B", "C", "D", "E", "F", "G", "H"}

    count := 0
    for _, f := range files {
        if cb[f][rank-1]{
            count++
        }
    }
    return count
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
    count := 0
    for _, _ = range cb {
        count++
    }
    return count*8
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
    countOcc := 0
    files := []string{"A", "B", "C", "D", "E", "F", "G", "H"}
    for _, f := range files {
        countOcc+=CountInFile(cb,f)
    }
    return countOcc
}
