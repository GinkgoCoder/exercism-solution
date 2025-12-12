package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
    res := 0
	if files, ok := cb[file]; ok {
        for _, v := range files {
            if v == true {
                res++
            }
        }
    }
    return res
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	res := 0
    for _, file := range cb {
        if rank > 0 && rank <= len(file) && file[rank-1] {
            res++
        }
    }
    return res
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	for _, v := range cb {
        return len(v) * len(cb)
    }
    return 0
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	res := 0
    for _, v := range cb {
        for _, o := range v {
            if o {
                res++
            }
        }
    }
    return res
}
