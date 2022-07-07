package game

import "errors"

type Game struct {
	Board
	Counter
}

type Board [3][3]string

type Counter [2]uint8

func (g *Game) Mark_X(positionX, positionY uint8) error {
	if err := g.ValidadeMovement(positionX, positionY); err != nil {
		return err
	}
	g.Board[positionX][positionY] = "X"
	g.Counter.TimesEachPlayed("X")
	return nil
}

func (g *Game) Mark_Y(positionX, positionY uint8) error {

	if err := g.ValidadeMovement(positionX, positionY); err != nil {
		return err
	}

	if g.Counter[0] > g.Counter[1] {
		g.Board[positionX][positionY] = "Y"
		g.Counter.TimesEachPlayed("Y")
		return nil
	}

	return errors.New("the x count must be either equals y or above")

}

func (g Game) ValidadeMovement(positionX, positionY uint8) error {
	if g.Board[positionX][positionY] != "" {
		return errors.New("can't rewrite an existing value")
	} else {
		return nil
	}
}

func (c *Counter) TimesEachPlayed(letter string) {
	if letter == "X" {
		c[0]++
	} else if letter == "Y" {
		c[1]++
	}
}

func (b Board) getMainDiagonal() [3]string {
	var mainDiagonal [3]string
	for i := 0; i < 3; i++ {
		mainDiagonal[i] = b[i][i]
	}
	return mainDiagonal
}

func (b Board) getInverseDiagonal() [3]string {
	var inverseDiagonal [3]string
	for i := 0; i < 3; i++ {
		inverseDiagonal[i] = b[i][2-i]
	}
	return inverseDiagonal
}

func (b Board) checkColumnsAndReturnWinner() string {

	for i := 0; i < 3; i++ {
		if b[0][i] == "X" && b[1][i] == "X" && b[2][i] == "X" {
			return "X"
		} else if b[0][i] == "Y" && b[1][i] == "Y" && b[2][i] == "Y" {
			return "Y"
		}
	}
	return ""
}

func (b Board) checkLinesAndReturnWinner() string {

	for i := 0; i < 3; i++ {
		if b[i][0] == "X" && b[i][1] == "X" && b[i][2] == "X" {
			return "X"
		} else if b[i][0] == "Y" && b[i][1] == "Y" && b[i][2] == "Y" {
			return "Y"
		}
	}
	return ""
}

func (b Board) checkDiagonalsAndReturnWinner() string {
	if main := b.getMainDiagonal(); main == [3]string{"Y", "Y", "Y"} || main == [3]string{"X", "X", "X"} {
		return main[0]
	} else if main := b.getInverseDiagonal(); main == [3]string{"Y", "Y", "Y"} || main == [3]string{"X", "X", "X"} {
		return main[0]
	}
	return ""
}

func (b Board) isFull() bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if b[i][j] == "" {
				return false
			}
		}
	}
	return true
}

func (b Board) winOrDraw() string {
	if value := b.checkDiagonalsAndReturnWinner(); value != "" {
		return value
	} else if value := b.checkColumnsAndReturnWinner(); value != "" {
		return value
	} else if value := b.checkLinesAndReturnWinner(); value != "" {
		return value
	} else if b.isFull() {
		return "its a Draw"
	}
	return ""
}
