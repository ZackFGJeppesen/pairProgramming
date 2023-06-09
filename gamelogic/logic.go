package gamelogic

import (
	"fmt"
	"log"
	"math/rand"
	"os"
)

type Player string

type Piece struct {
	Colour string
	Kind   string
}

type Square struct {
	Letter   int
	Number   int
	Occupied *Piece
}

var whitePawnDir = []string{"n"}
var blackPawnDir = []string{"s"}
var kingDir = []string{"nw", "n", "ne", "w", "e", "sw", "s", "se"}
var queenDir = []string{"nw", "n", "ne", "w", "e", "sw", "s", "se"}
var rookDir = []string{"n", "w", "e", "s"}
var bishopDir = []string{"nw", "ne", "sw", "se"}

//var knightDir = []string{}

var white Player
var black Player

var Board []Square //= createBoard()
var CurrentPlayer *Player = &white

func GetSquare(i int) (int, int, string, string) {
	var kind string
	var colour string
	if Board[i].Occupied == nil {
		kind = ""
		colour = ""
	} else {
		kind = Board[i].Occupied.Kind
		colour = Board[i].Occupied.Colour
	}
	return Board[i].Letter, Board[i].Number, kind, colour
}

func SetUpPlayer() {
	white = Player("white")
	black = Player("black")

}

func Move() {
	var lm map[Square][]Square
	var squaresWithPieces []Square
	for {
		squaresWithPieces = randPickSquare()
		lm = legalMoves(squaresWithPieces)
		if len(lm) != 0 {
			break
		}
	}
	s, point := randMove(lm)
	fmt.Printf("Piece %s is at Square (%d, %d) \n", s.Occupied.Kind, s.Letter, s.Number)
	fmt.Println(lm)
	Board[step(s, "")].Occupied = nil
	Board[step(point, "")].Occupied = s.Occupied
	swap()
	point = Board[step(point, "")]
	fmt.Printf("Piece %s is now at Square (%d, %d) \n", point.Occupied.Kind, point.Letter, point.Number)
}

func SetUpBoard() {
	Board = make([]Square, 64)
	for i := 0; i < 64; i++ {
		temp := Square{Letter: i % 8, Number: int(i / 8), Occupied: nil}
		Board[i] = temp
	}

	for i := 0; i < 8; i++ {
		createPiece("white", "whitePawn", &Board[i+48])
		createPiece("black", "blackPawn", &Board[i+8])
	}
	createPiece("black", "rook", &Board[7])
	createPiece("black", "rook", &Board[0])
	createPiece("white", "rook", &Board[56])
	createPiece("white", "rook", &Board[63])
	createPiece("black", "knight", &Board[1])
	createPiece("black", "knight", &Board[6])
	createPiece("white", "knight", &Board[57])
	createPiece("white", "knight", &Board[62])
	createPiece("black", "bishop", &Board[2])
	createPiece("black", "bishop", &Board[5])
	createPiece("white", "bishop", &Board[58])
	createPiece("white", "bishop", &Board[61])
	createPiece("black", "queen", &Board[3])
	createPiece("black", "king", &Board[4])
	createPiece("white", "queen", &Board[59])
	createPiece("white", "king", &Board[60])
}

func randPickSquare() []Square {
	var squares []Square = []Square{}
	for n := range Board {
		if Board[n].Occupied != nil && Board[n].Occupied.Colour == string(*CurrentPlayer) {
			squares = append(squares, Board[n])
		}
	}
	return squares
}

func randMove(returnMoves map[Square][]Square) (Square, Square) {
	if len(returnMoves) == 0 {
		fmt.Println("Draw")
		os.Exit(0)
	}
	fmt.Println("--------------", len(returnMoves))
	r := rand.Intn(len(returnMoves))
	counter := 0
	for from, to := range returnMoves {
		if r == counter {
			return from, to[rand.Intn(len(to))]
		} else {
			counter++
		}
	}
	return Square{}, Square{}
}

func legalMoves(squares []Square) map[Square][]Square {
	returnMoves := make(map[Square][]Square)
	for _, s := range squares {
		if s.Occupied.Kind != "knight" {
			tempMoves := legalStep(s)
			if len(tempMoves) > 0 {
				returnMoves[s] = tempMoves
			}
		} else {
			tempMoves := legalKnightStep(s)
			if len(tempMoves) > 0 {
				returnMoves[s] = tempMoves
			}
		}
	}
	return returnMoves
}

func swap() {
	if *CurrentPlayer == "white" {
		CurrentPlayer = &black
	} else {
		CurrentPlayer = &white
	}
}

func legalKnightStep(s Square) []Square {
	squareMoves := []Square{}
	var nextStep Square
	nextStep = Square{Letter: s.Letter - 2, Number: s.Number + 1}
	str := check(nextStep)
	if str == "empty" || str == "capture" {
		squareMoves = append(squareMoves, nextStep)
	}
	nextStep = Square{Letter: s.Letter - 1, Number: s.Number + 2}
	str = check(nextStep)
	if str == "empty" || str == "capture" {
		squareMoves = append(squareMoves, nextStep)
	}
	nextStep = Square{Letter: s.Letter + 1, Number: s.Number + 2}
	str = check(nextStep)
	if str == "empty" || str == "capture" {
		squareMoves = append(squareMoves, nextStep)
	}
	nextStep = Square{Letter: s.Letter + 2, Number: s.Number + 1}
	str = check(nextStep)
	if str == "empty" || str == "capture" {
		squareMoves = append(squareMoves, nextStep)
	}
	nextStep = Square{Letter: s.Letter + 2, Number: s.Number - 1}
	str = check(nextStep)
	if str == "empty" || str == "capture" {
		squareMoves = append(squareMoves, nextStep)
	}
	nextStep = Square{Letter: s.Letter + 1, Number: s.Number - 2}
	str = check(nextStep)
	if str == "empty" || str == "capture" {
		squareMoves = append(squareMoves, nextStep)
	}
	nextStep = Square{Letter: s.Letter - 1, Number: s.Number - 2}
	str = check(nextStep)
	if str == "empty" || str == "capture" {
		squareMoves = append(squareMoves, nextStep)
	}
	nextStep = Square{Letter: s.Letter - 2, Number: s.Number - 1}
	str = check(nextStep)
	if str == "empty" || str == "capture" {
		squareMoves = append(squareMoves, nextStep)
	}
	return squareMoves
}

func step(s Square, dir string, dist ...int) int {
	var distance int
	if len(dist) == 0 {
		distance = 1
	} else {
		distance = dist[0]
	}
	switch dir {
	case "nw":
		return s.Letter - distance + (s.Number-distance)*8
	case "n":
		return s.Letter + (s.Number-distance)*8
	case "ne":
		return s.Letter + distance + (s.Number-distance)*8
	case "w":
		return s.Letter - distance + (s.Number)*8
	case "e":
		return s.Letter + distance + (s.Number)*8
	case "sw":
		return s.Letter - distance + (s.Number+distance)*8
	case "s":
		return s.Letter + (s.Number+distance)*8
	case "se":
		return s.Letter + distance + (s.Number+distance)*8
	default:
		return s.Letter + (s.Number)*8
	}
}

func direction(p Square) map[string]int {
	moves := make(map[string]int)
	switch p.Occupied.Kind {
	case "whitePawn":
		//n capture
		if Board[step(p, "n")].Occupied == nil {
			if p.Number == 6 && Board[step(p, "n", 2)].Occupied == nil {
				moves["n"] = 2
			} else {
				moves["n"] = 1
			}
		} else {
			moves["n"] = 0
		}
		//nw capture
		if Board[step(p, "nw")].Occupied != nil && Board[step(p, "nw")].Occupied.Colour != p.Occupied.Colour {
			moves["nw"] = 1
		}

		//ne capture
		if Board[step(p, "ne")].Occupied != nil && Board[step(p, "ne")].Occupied.Colour != p.Occupied.Colour {
			moves["ne"] = 1
		}
	case "blackPawn":
		//s capture
		if Board[step(p, "s")].Occupied == nil {
			if p.Number == 1 && Board[step(p, "s", 2)].Occupied == nil {
				moves["s"] = 2
			} else {
				moves["s"] = 1
			}
		} else {
			moves["s"] = 0
		}
		//sw capture
		if Board[step(p, "sw")].Occupied != nil && Board[step(p, "sw")].Occupied.Colour != p.Occupied.Colour {
			moves["sw"] = 1
		}

		//se capture
		if Board[step(p, "se")].Occupied != nil && Board[step(p, "se")].Occupied.Colour != p.Occupied.Colour {
			moves["se"] = 1
		}
	case "king":
		for _, dir := range kingDir {
			moves[dir] = 1
		}
	case "queen":
		for _, dir := range queenDir {
			moves[dir] = 8
		}
	case "rook":
		for _, dir := range rookDir {
			moves[dir] = 8
		}
	case "bishop":
		for _, dir := range bishopDir {
			moves[dir] = 8
		}
	default:
		log.Fatal("Error, no piece selected")
	}
	return moves
}

func check(nextStep Square) string {
	if 0 <= nextStep.Letter && nextStep.Letter < 8 && 0 <= nextStep.Number && nextStep.Number < 8 {
		if Board[step(nextStep, "")].Occupied == nil {
			return "empty"
		} else if Board[step(nextStep, "")].Occupied.Colour == string(*CurrentPlayer) {
			return "friend"
		} else {
			if Board[step(nextStep, "")].Occupied.Kind == "king" {
				fmt.Printf("%v Player has Won the Game! \n", CurrentPlayer)
				os.Exit(0)
			}
			return "capture"
		}
	}
	return "outOfBounds"
}

func createPiece(colour, kind string, Square *Square) {
	temp := &Piece{Colour: colour, Kind: kind}
	Square.Occupied = temp
}

func legalStep(s Square) []Square {
	squareMoves := []Square{}
	for dir, dist := range direction(s) {
		var nextStep Square
	loop:
		for step := 1; step <= dist; step++ {
			switch dir {
			case "nw":
				nextStep = Square{Letter: s.Letter - step, Number: s.Number - step}
			case "n":
				nextStep = Square{Letter: s.Letter, Number: s.Number - step}
			case "ne":
				nextStep = Square{Letter: s.Letter + step, Number: s.Number - step}
			case "w":
				nextStep = Square{Letter: s.Letter - step, Number: s.Number}
			case "e":
				nextStep = Square{Letter: s.Letter + step, Number: s.Number}
			case "sw":
				nextStep = Square{Letter: s.Letter - step, Number: s.Number + step}
			case "s":
				nextStep = Square{Letter: s.Letter, Number: s.Number + step}
			case "se":
				nextStep = Square{Letter: s.Letter + step, Number: s.Number + step}
			default:
			}
			switch check(nextStep) {
			case "empty":
				squareMoves = append(squareMoves, nextStep)
			case "capture":
				squareMoves = append(squareMoves, nextStep)
				break loop
			case "friend":
				break loop
			case "outOfBounds":
				break loop
			default:
				break loop
			}
		}
	}
	return squareMoves
}
