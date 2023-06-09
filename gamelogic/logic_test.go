package gamelogic

import (
	"testing"
)

func TestSetUpPlayer(t *testing.T) {
	SetUpPlayer()
	got := white
	want := Player("white")
	if got != want {
		t.Errorf("got %s, but wanted %s", got, want)
	}
	got = black
	want = Player("black")
	if got != want {
		t.Errorf("got %s, but wanted %s", got, want)
	}
}

func TestCreatePiece(t *testing.T) {
	SetUpBoard()
	createPiece("white", "rook", &Board[0])
	if Board[0].Occupied == nil || Board[0].Occupied.Kind != "rook" || Board[0].Occupied.Colour != "white" {
		t.Errorf("The piece was not created correctly")
	}

}

func TestSetUpBoard(t *testing.T) {
	SetUpBoard()
	if len(Board) != 64 {
		t.Errorf("the board is %d squares, should be 64", len(Board))
	}
	if Board[30].Occupied != nil {
		t.Errorf("On square 30 there should be no pieces but there is %+v ", Board[30].Occupied)
	}

	if Board[4].Occupied == nil {
		t.Errorf("On square 4 there should be a black king but square 4 is %v", Board[4].Occupied)
	}
}

func TestGetSquare(t *testing.T) {
	SetUpBoard()
	l, n, k, c := GetSquare(0)
	if l != 0 || n != 0 || k != "rook" || c != "black" {
		t.Errorf("Incorrect black square was retrieved")
	}

	l, n, k, c = GetSquare(62)
	if l != 6 || n != 7 || k != "knight" || c != "white" {
		t.Errorf("Incorrect white square was retrieved")
	}
}

func TestRandPickSquare(t *testing.T) {
	SetUpBoard()
	SetUpPlayer()
	if len(randPickSquare()) != 16 {
		t.Errorf("Number of pieces is %d and should be 16 at start", len(randPickSquare()))
	}
	for _, s := range randPickSquare() {
		if s.Occupied == nil || s.Occupied.Colour != "white" {
			t.Errorf("randPickSquare returns a square that is incorrect")
		}
	}
}

func TestSwap(t *testing.T) {
	SetUpPlayer()
	if *CurrentPlayer != "white" {
		t.Errorf("current player should be white but is %s", *CurrentPlayer)
	}
	swap()
	if *CurrentPlayer != "black" {
		t.Errorf("current player should be black but is %s", *CurrentPlayer)
	}
}

func TestStep(t *testing.T) {
	SetUpBoard()
	if step(Board[30], "sw") != 37 {
		t.Errorf("step should have returned 37 but gave %d", step(Board[30], "sw"))
	}

	if step(Board[30], "") != 30 {
		t.Errorf("step should have returned 30 but gave %d", step(Board[30], ""))
	}

	if step(Board[30], "ne", 2) != 16 {
		t.Errorf("step should have returned 16 but gave %d", step(Board[30], "ne", 2))
	}
}

func TestLegalMoves(t *testing.T) {
	SetUpBoard()
	SetUpPlayer()
	m := legalMoves([]Square{Board[62]}) // 45 47
	if squares, ok := m[Board[62]]; ok {
		for _, s := range squares {
			if s != Board[45] && s != Board[47] {
				t.Errorf("did not return a legal move 45 or 47 for knight on square 62 it returned %v", s)
			}
		}
	} else {
		t.Errorf("legalMoves did not return a map with moves starting from 62")
	}

	m = legalMoves([]Square{Board[48]}) // 40 32
	if squares, ok := m[Board[48]]; ok {
		for _, s := range squares {
			if s != Board[40] && s != Board[32] {
				t.Errorf("did not return a legal move 40 or 32 for pawn on square 48 it returned %v", s)
			}
		}
	} else {
		t.Errorf("legalMoves did not return a map with moves starting from 48")
	}

}

func TestRandMove(t *testing.T) {
	SetUpBoard()
	SetUpPlayer()

	m := legalMoves([]Square{Board[62]}) // 45 && 47
	from, to := randMove(m)
	if from != Board[62] {
		t.Errorf("randMove returned wrong from square should have been 62, was %v", from)
	}
	if to != Board[45] && to != Board[47] {
		t.Errorf("randMove did not return the legal moves 53 or 55 for knight on square 62 it returned %v", to)
	}

	m = legalMoves([]Square{Board[48]}) // 40 32
	from, to = randMove(m)
	if from != Board[48] {
		t.Errorf("randMove returned wrong from square should have been 48, was %v", from)
	}
	if to != Board[40] && to != Board[32] {
		t.Errorf("randMove returned did not return the legal moves 48 or 32 for pawn on square 48 it returned %v", to)
	}
}

func TestLegalStep(t *testing.T) {
	SetUpBoard()
	SetUpPlayer()
	squares := legalStep(Board[48])
	for _, s := range squares {
		if s != Board[40] && s != Board[32] {
			t.Errorf("did not return a legal move 40 or 32 for pawn on square 48 it returned %v", s)
		}
	}
	createPiece("white", "whitePawn", &Board[32])
	squares = legalStep(Board[48])
	for _, s := range squares {
		if s != Board[40] && s != Board[32] {
			t.Errorf("did not return a legal move 40 or 32 for pawn on square 48 it returned %v", s)
		}
	}
	createPiece("black", "blackPawn", &Board[40])
	squares = legalStep(Board[48])
	for _, s := range squares {
		if s != Board[40] && s != Board[32] {
			t.Errorf("did not return a legal move 40 or 32 for pawn on square 48 it returned %v", s)
		}
	}
}

func TestLegalKnightStep(t *testing.T) {
	SetUpBoard()
	SetUpPlayer()
	squares := legalKnightStep(Board[62]) // 45 && 47
	for _, s := range squares {
		if s != Board[45] && s != Board[47] {
			t.Errorf("did not return a legal move 45 or 47 for knight on square 62 it returned %v", s)
		}
	}
}

func TestMove(t *testing.T) {
	Board = make([]Square, 64)
	for i := 0; i < 64; i++ {
		temp := Square{Letter: i % 8, Number: int(i / 8), Occupied: nil}
		Board[i] = temp
	}
	SetUpPlayer()

	createPiece("white", "whitePawn", &Board[15])

	Move()
	// if Board[8].Occupied != nil {
	// 	t.Errorf("Pawn didnt move")
	// }
	// if Board[0].Occupied == nil {
	// 	t.Errorf("Pawn didnt move correctly")
	// }
	swap()
	Move()
}
