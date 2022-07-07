package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGame(t *testing.T) {

	t.Run("Insert X", func(t *testing.T) {
		want := Game{Board{}, Counter{}}
		want.Board[2][2] = "X"
		want.Counter[0] = 1

		got := Game{Board{}, Counter{}}
		got.Mark_X(2, 2)

		assert.Equal(t, want, got)
	})

	t.Run("Insert Y", func(t *testing.T) {
		want := Game{Board{}, Counter{}}
		want.Mark_X(1, 1)
		want.Board[2][2] = "Y"
		want.Counter[1] = 1

		got := Game{Board{}, Counter{}}
		got.Mark_X(1, 1)

		got.Mark_Y(2, 2)

		assert.Equal(t, want, got)
	})

	t.Run("Count times played", func(t *testing.T) {
		game := Game{Board{}, Counter{}}
		game.Mark_X(2, 1)

		want := Counter{1, 0}

		got := game.Counter

		assert.Equal(t, want, got)
	})

	t.Run("Can't insert in the same place -->", func(t *testing.T) {

		t.Run("Inserting X above X", func(t *testing.T) {
			game := Game{Board{}, Counter{}}
			game.Mark_X(0, 0)
			game.Mark_Y(1, 1)

			want := "can't rewrite an existing value"

			err := game.Mark_X(0, 0)

			assert.EqualError(t, err, want)
		})

		t.Run("Inserting Y above Y", func(t *testing.T) {
			game := Game{Board{}, Counter{}}
			game.Mark_X(0, 0)
			game.Mark_Y(1, 1)
			game.Mark_X(2, 2)

			want := "can't rewrite an existing value"

			err := game.Mark_Y(1, 1)

			assert.EqualError(t, err, want)
		})

		t.Run("Inserting Y above X", func(t *testing.T) {
			game := Game{Board{}, Counter{}}
			game.Mark_X(0, 0)
			game.Mark_Y(1, 1)

			want := "can't rewrite an existing value"

			err := game.Mark_Y(0, 0)

			assert.EqualError(t, err, want)
		})

		t.Run("Inserting X above Y", func(t *testing.T) {
			game := Game{Board{}, Counter{}}
			game.Mark_X(0, 0)
			game.Mark_Y(1, 1)

			want := "can't rewrite an existing value"

			err := game.Mark_X(1, 1)

			assert.EqualError(t, err, want)
		})

	})

	t.Run("trying to go insert Y--> ", func(t *testing.T) {

		t.Run("As first player", func(t *testing.T) {
			game := Game{Board{}, Counter{}}

			want := "the x count must be either equals y or above"

			err := game.Mark_Y(1, 1)

			assert.EqualError(t, err, want)
		})

		t.Run("Again", func(t *testing.T) {
			game := Game{Board{}, Counter{}}
			game.Mark_X(1, 2)
			game.Mark_Y(0, 2)

			want := "the x count must be either equals y or above"

			err := game.Mark_Y(1, 1)

			assert.EqualError(t, err, want)
		})

	})

}

func TestBoard(t *testing.T) {
	t.Run("get main diagonal", func(t *testing.T) {
		b := Board{{"Y", "X", "X"}, {"X", "Y", "X"}, {"X", "X", "Y"}}

		want := [3]string{"Y", "Y", "Y"}

		got := b.getMainDiagonal()

		assert.Equal(t, want, got)
	})

	t.Run("get Inverse diagonal", func(t *testing.T) {
		b := Board{{"X", "X", "Y"}, {"X", "Y", "X"}, {"Y", "X", "X"}}

		want := [3]string{"Y", "Y", "Y"}

		got := b.getInverseDiagonal()

		assert.Equal(t, want, got)
	})

	t.Run("check lines for winner-->", func(t *testing.T) {
		t.Run("It must be either X or Y", func(t *testing.T) {
			boards := []Board{
				{{"X", "X", "X"}, {"Y", "Y", ""}, {"", "", ""}},
				{{"Y", "Y", ""}, {"X", "X", "X"}, {"X", "", ""}},
				{{"Y", "X", "Y"}, {"X", "Y", ""}, {"X", "X", "X"}},
				{{"Y", "Y", "Y"}, {"Y", "X", ""}, {"", "", "X"}},
				{{"", "Y", "X"}, {"Y", "Y", "Y"}, {"X", "", ""}},
				{{"Y", "X", "Y"}, {"X", "Y", ""}, {"Y", "Y", "Y"}},
			}

			for _, board := range boards {
				if value := board.checkLinesAndReturnWinner(); value != "X" && value != "Y" {
					t.Errorf("it was expection x or y but got %q", value)
				}
			}
		})

		t.Run("It must return nothing", func(t *testing.T) {
			b := Board{{"X", "X", "Y"}, {"X", "Y", "X"}, {"Y", "X", "X"}}
			if value := b.checkLinesAndReturnWinner(); value != "" {
				t.Errorf("it was expection x or y but got %s", value)
			}
		})

	})

	t.Run("check diagonals and return a winner", func(t *testing.T) {
		boards := []Board{
			{{"X", "X", "Y"}, {"X", "Y", "X"}, {"Y", "X", "X"}},
			{{"Y", "X", "X"}, {"X", "Y", "X"}, {"X", "X", "Y"}},
		}

		for _, board := range boards {
			if value := board.checkDiagonalsAndReturnWinner(); value != "Y" {
				t.Errorf("it was expection y but got %q", value)
			}
		}
	})

	t.Run("win or Draw? -->", func(t *testing.T) {

		t.Run("its a win", func(t *testing.T) {
			boards := []Board{
				{{"X", "X", "X"}, {"Y", "Y", ""}, {"", "", ""}},
				{{"Y", "Y", "Y"}, {"X", "X", ""}, {"X", "", ""}},
				{{"Y", "X", "Y"}, {"X", "Y", ""}, {"", "X", "Y"}},
				{{"X", "Y", ""}, {"Y", "X", ""}, {"", "", "X"}},
				{{"", "Y", "X"}, {"Y", "X", ""}, {"X", "", ""}},
				{{"Y", "X", "Y"}, {"X", "Y", ""}, {"Y", "X", ""}},
			}

			for _, board := range boards {
				if value := board.winOrDraw(); value != "Y" && value != "X" {
					t.Errorf("it was expection x or y but got %q", value)
				}
			}
		})

		t.Run("it is a draw", func(t *testing.T) {
			b := Board{{"Y", "Y", "X"}, {"X", "X", "Y"}, {"Y", "X", "Y"}}

			want := "its a Draw"

			got := b.winOrDraw()

			assert.Equal(t, want, got)
		})

	})
}
