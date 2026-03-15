package engine

// scanDiagonal - returns:
//
//	true if there is an attacking piece in the diagonal
//	false if there is no piece attacking the diagonal
func (e *Engine) scanDiagonal(sqx, sqy, dx, dy int32, attackingColor PieceColor) bool {
	startingX := sqx + dx
	startingY := sqy + dy

	for e.isInsideBoard(startingX, startingY) {
		piece := e.board[startingY][startingX]
		if piece != nil {
			if piece.Color == attackingColor && (piece.Type == Bishop || piece.Type == Queen) {
				return true
			} else {
				// Anything else in the path not an attacking bishop or queen is automatically an obstacle
				return false
			}
		}

		startingX = startingX + dx
		startingY = startingY + dy
	}

	return false
}

// scanStraightPath - returns:
//
//	true if there is an attacking piece in the straight paths
//	false if there is no piece attacking the straight paths
func (e *Engine) scanStraightPath(sqx, sqy, dx, dy int32, attackingColor PieceColor) bool {
	startingX := sqx + dx
	startingY := sqy + dy

	for e.isInsideBoard(startingX, startingY) {
		piece := e.board[startingY][startingX]
		if piece != nil {
			if piece.Color == attackingColor && (piece.Type == Rook || piece.Type == Queen) {
				return true
			} else {
				// Anything else in the path not an attacking rook or queen is automatically an obstacle
				return false
			}
		}

		startingX = startingX + dx
		startingY = startingY + dy
	}

	return false
}
