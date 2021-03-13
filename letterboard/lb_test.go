package letterboard

import (
	"reflect"
	"testing"
)

func assertMoves(t *testing.T, expected Moves, actual Moves) {
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("\nExpected\n  %v,\nbut received\n  %v", expected, actual)
	}
}

func TestSolveLetterboard_With_Solve(t *testing.T) {
	actual := SolveLetterboard([]rune{'k', 's', 'u', 'o', 'c', 'l', 'v', 'd', 'e'}, "solve")
	expected := Moves{
		/* s */ Move{Left, -1}, Move{Left, 's'},
		/* o */ Move{Left, -1}, Move{Left, 'o'},
		/* l */ Move{Left, -1}, Move{Left, 'l'},
		/* v */ Move{Left, 'v'},
		/* e */ Move{Left, -1}, Move{Left, 'e'},
	}

	assertMoves(t, expected, actual)
}

func TestSolveLetterboard_With_Duck(t *testing.T) {
	actual := SolveLetterboard([]rune{'k', 's', 'u', 'o', 'c', 'l', 'v', 'd', 'e'}, "duck")
	expected := Moves{
		/* d */ Move{Right, -1}, Move{Right, -1}, Move{Right, -1}, Move{Right, -1}, Move{Right, -1}, Move{Right, -1}, Move{Right, 'd'},
		/* u */ Move{Left, 'u'},
		/* c */ Move{Left, 'c'}, Move{Right, -1}, Move{Right, -1}, Move{Right, -1},
		/* k */ Move{Right, 'k'},
	}

	assertMoves(t, expected, actual)
}
