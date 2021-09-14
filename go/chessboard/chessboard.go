package chessboard

// Rank stores if a square is occupied by a piece
type Rank []bool

// Chessboard contains eight Ranks, accessed with values from 'A' to 'H'
type Chessboard map[byte]Rank

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank
func (cb Chessboard) CountInRank(rank byte) (ret int) {
	for _, val := range cb[rank] {
		if val {
			ret++
		}
	}
	return ret
}

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file
func (cb Chessboard) CountInFile(file int) (ret int) {
	if file > 8 {
		return 0
	}
	for _, b := range cb {
		if b[file-1] {
			ret++
		}
	}
	return ret
}

// CountAll should count how many squares are present in the chessboard
func (cb Chessboard) CountAll() (ret int) {
	for _, r := range cb {
		for range r {
			ret++
		}
	}
	return ret
}

// CountOccupied returns how many squares are occupied in the chessboard
func (cb Chessboard) CountOccupied() (ret int) {
	for _, r := range cb {
		for _, v := range r {
			if v {
				ret++
			}
		}
	}
	return ret
}
