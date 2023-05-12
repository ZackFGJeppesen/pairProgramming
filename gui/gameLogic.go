package gui

import (
	"fmt"
	"log"
	"math/rand"
)

type Player struct {
	Name   string
	Pieces []*Piece
}
type Piece struct {
	Colour  string
	Kind    string
	Img interface{}
}

type Point struct {
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
var knightDir = []string{}

var white Player
var black Player

var Board []*Point //= createBoard()
var CurrentPlayer *Player = &white


func RandPickPiece() *Point {
	for {
		n := rand.Intn(len(Board))
		if Board[n].Occupied != nil && Board[n].Occupied.Colour == CurrentPlayer.Name {
			return Board[n]
		}
	}
}

func RandMove(returnMoves []Point) Point {
	fmt.Println("--------------", len(returnMoves))
	return returnMoves[rand.Intn(len(returnMoves))]
}

func Swap() {
	if CurrentPlayer.Name == "white" {
		CurrentPlayer = &black
	} else {
		CurrentPlayer = &white
	}
}

func LegalMoves(p Point) []Point {
	var returnMoves []Point
	if p.Occupied.Kind != "knight" {
		for dir, dist := range direction(*p.Occupied) {

			var nextStep Point
		loop:
			for step := 1; step <= dist; step++ {
				switch dir {
				case "nw":
					nextStep = Point{Letter: p.Letter - step, Number: p.Number + step}
				case "n":
					nextStep = Point{Letter: p.Letter, Number: p.Number + step}
				case "ne":
					nextStep = Point{Letter: p.Letter + step, Number: p.Number + step}
				case "w":
					nextStep = Point{Letter: p.Letter - step, Number: p.Number}
				case "e":
					nextStep = Point{Letter: p.Letter + step, Number: p.Number}
				case "sw":
					nextStep = Point{Letter: p.Letter - step, Number: p.Number - step}
				case "s":
					nextStep = Point{Letter: p.Letter, Number: p.Number - step}
				case "se":
					nextStep = Point{Letter: p.Letter + step, Number: p.Number - step}
				default:
				}
				switch check(nextStep) {
				case "empty":
					returnMoves = append(returnMoves, nextStep)
				case "capture":
					returnMoves = append(returnMoves, nextStep)
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
	} else {
		tempP := Point{Letter: p.Letter - 2, Number: p.Number + 1}
		s := check(tempP)
		if s == "empty" || s == "capture" {
			returnMoves = append(returnMoves, tempP)
		}

		tempP = Point{Letter: p.Letter - 1, Number: p.Number + 2}
		s = check(tempP)
		if s == "empty" || s == "capture" {
			returnMoves = append(returnMoves, tempP)
		}

		tempP = Point{Letter: p.Letter + 1, Number: p.Number + 2}
		s = check(tempP)
		if s == "empty" || s == "capture" {
			returnMoves = append(returnMoves, tempP)
		}

		tempP = Point{Letter: p.Letter + 2, Number: p.Number + 1}
		s = check(tempP)
		if s == "empty" || s == "capture" {
			returnMoves = append(returnMoves, tempP)
		}

		tempP = Point{Letter: p.Letter + 2, Number: p.Number - 1}
		s = check(tempP)
		if s == "empty" || s == "capture" {
			returnMoves = append(returnMoves, tempP)
		}

		tempP = Point{Letter: p.Letter + 1, Number: p.Number - 2}
		s = check(tempP)
		if s == "empty" || s == "capture" {
			returnMoves = append(returnMoves, tempP)
		}

		tempP = Point{Letter: p.Letter - 1, Number: p.Number - 2}
		s = check(tempP)
		if s == "empty" || s == "capture" {
			returnMoves = append(returnMoves, tempP)
		}

		tempP = Point{Letter: p.Letter - 2, Number: p.Number - 1}
		s = check(tempP)
		if s == "empty" || s == "capture" {
			returnMoves = append(returnMoves, tempP)
		}
	}
	return returnMoves
}

func direction(p Piece) map[string]int {
	moves := make(map[string]int)
	switch p.Kind {
	case "whitePawn":
		for _, dir := range whitePawnDir {
			moves[dir] = 1
		}

	case "blackPawn":
		for _, dir := range blackPawnDir {
			moves[dir] = 1
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

func check(nextStep Point) string {
	if 0 <= nextStep.Letter && nextStep.Letter < 8 && 0 <= nextStep.Number && nextStep.Number < 8 {
		if Board[nextStep.Letter*8+nextStep.Number].Occupied == nil {
			return "empty"
		} else if Board[nextStep.Letter*8+nextStep.Number].Occupied.Colour == CurrentPlayer.Name {
			return "friend"
		} else {
			return "capture"
		}
	}
	return "outOfBounds"
}

func SetUpBoard() {
	Board = make([]*Point, 64)
	for i := 0; i < 64; i++ {
		temp := Point{Letter: i%8, Number: int(i/8), Occupied: nil}
		Board[i] = &temp
	}
}

func createPiece(colour, kind string, point *Point, n int) {
	fmt.Println(colour, kind, point, n)
	temp := &Piece{Colour: colour, Kind: kind, Img: nil}
	point.Occupied = temp
	CurrentPlayer.Pieces[n] = temp
}

func SetUpPlayer() {
	white = Player{Name: "white", Pieces: make([]*Piece,16)}
	black = Player{Name: "black", Pieces: make([]*Piece,16)}
	for i := 0; i < 8; i++ {
		createPiece("white", "whitePawn", Board[i+8], i)
		createPiece("black", "blackPawn", Board[i+47], i)
	}
	createPiece("black", "rook", Board[56],8)
	createPiece("black", "rook", Board[63],9)
	createPiece("white", "rook", Board[7],8)
	createPiece("white", "rook", Board[0],9)
	createPiece("black", "knight", Board[57],10)
	createPiece("black", "knight", Board[62],11)
	createPiece("white", "knight", Board[1],10)
	createPiece("white", "knight", Board[6],11)
	createPiece("black", "bishop", Board[58],12)
	createPiece("black", "bishop", Board[61],13)
	createPiece("white", "bishop", Board[2],12)
	createPiece("white", "bishop", Board[5],13)
	createPiece("black", "queen", Board[59],14)
	createPiece("black", "king", Board[60],15)
	createPiece("white", "queen", Board[3],14)
	createPiece("white", "king", Board[4],15)
}
