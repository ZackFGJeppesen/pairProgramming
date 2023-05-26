package gamelogic

import (
	"fmt"
	"log"
	"math/rand"
	"os"

	"github.com/hajimehoshi/ebiten"
)

type Player struct {
	Name string
}
type Piece struct {
	Colour string
	Kind   string
	Img    *ebiten.Image
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
var knightDir = []string{}

var white Player
var black Player

var Board []Square //= createBoard()
var CurrentPlayer *Player = &white

func RandPickSquare() []Square {
	var squares []Square = []Square{}
	for n := range Board {
		if Board[n].Occupied != nil && Board[n].Occupied.Colour == CurrentPlayer.Name {
			squares = append(squares, Board[n])
		}
	}
	return squares
}

func RandMove(returnMoves map[Square]Square) (Square, Square) {
	fmt.Println("--------------", len(returnMoves))
	r := rand.Intn(len(returnMoves))
	counter := 0
	for from, to := range returnMoves {
		if r == counter {
			return from, to
		} else {
			counter++
		}
	}
	return Square{}, Square{}
}

func Swap() {
	if CurrentPlayer.Name == "white" {
		CurrentPlayer = &black
	} else {
		CurrentPlayer = &white
	}
}

func LegalMoves(pieces []Square) map[Square]Square {
	returnMoves := make(map[Square]Square)
	for _, p := range pieces {
		if p.Occupied.Kind != "knight" {
			for dir, dist := range direction(p) {

				var nextStep Square
			loop:
				for step := 1; step <= dist; step++ {
					switch dir {
					case "nw":
						nextStep = Square{Letter: p.Letter - step, Number: p.Number - step}
					case "n":
						nextStep = Square{Letter: p.Letter, Number: p.Number - step}
					case "ne":
						nextStep = Square{Letter: p.Letter + step, Number: p.Number - step}
					case "w":
						nextStep = Square{Letter: p.Letter - step, Number: p.Number}
					case "e":
						nextStep = Square{Letter: p.Letter + step, Number: p.Number}
					case "sw":
						nextStep = Square{Letter: p.Letter - step, Number: p.Number + step}
					case "s":
						nextStep = Square{Letter: p.Letter, Number: p.Number + step}
					case "se":
						nextStep = Square{Letter: p.Letter + step, Number: p.Number + step}
					default:
					}
					switch check(nextStep) {
					case "empty":
						returnMoves[p] = nextStep
					case "capture":
						returnMoves[p] = nextStep
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
			var nextStep Square
			nextStep = Square{Letter: p.Letter - 2, Number: p.Number + 1}
			s := check(nextStep)
			if s == "empty" || s == "capture" {
				returnMoves[p] = nextStep
			}

			nextStep = Square{Letter: p.Letter - 1, Number: p.Number + 2}
			s = check(nextStep)
			if s == "empty" || s == "capture" {
				returnMoves[p] = nextStep
			}

			nextStep = Square{Letter: p.Letter + 1, Number: p.Number + 2}
			s = check(nextStep)
			if s == "empty" || s == "capture" {
				returnMoves[p] = nextStep
			}

			nextStep = Square{Letter: p.Letter + 2, Number: p.Number + 1}
			s = check(nextStep)
			if s == "empty" || s == "capture" {
				returnMoves[p] = nextStep
			}

			nextStep = Square{Letter: p.Letter + 2, Number: p.Number - 1}
			s = check(nextStep)
			if s == "empty" || s == "capture" {
				returnMoves[p] = nextStep
			}

			nextStep = Square{Letter: p.Letter + 1, Number: p.Number - 2}
			s = check(nextStep)
			if s == "empty" || s == "capture" {
				returnMoves[p] = nextStep
			}

			nextStep = Square{Letter: p.Letter - 1, Number: p.Number - 2}
			s = check(nextStep)
			if s == "empty" || s == "capture" {
				returnMoves[p] = nextStep
			}

			nextStep = Square{Letter: p.Letter - 2, Number: p.Number - 1}
			s = check(nextStep)
			if s == "empty" || s == "capture" {
				returnMoves[p] = nextStep
			}
		}
	}
	return returnMoves
}

func direction(p Square) map[string]int {
	moves := make(map[string]int)
	switch p.Occupied.Kind {
	case "whitePawn":
		for _, dir := range whitePawnDir {
			if p.Number == 6 {
				moves[dir] = 2
			} else {
				moves[dir] = 1
			}
		}
		//nw capture
		if Board[p.Letter-1+(p.Number-1)*8].Occupied != nil && Board[p.Letter-1+(p.Number-1)*8].Occupied.Colour != p.Occupied.Colour {
			moves["nw"] = 1
		}

		//ne capture
		if Board[p.Letter+1+(p.Number-1)*8].Occupied != nil && Board[p.Letter+1+(p.Number-1)*8].Occupied.Colour != p.Occupied.Colour {
			moves["ne"] = 1
		}

	case "blackPawn":
		for _, dir := range blackPawnDir {
			if p.Number == 1 {
				moves[dir] = 2
			} else {
				moves[dir] = 1
			}
		}
		//sw capture
		if Board[p.Letter-1+(p.Number+1)*8].Occupied != nil && Board[p.Letter-1+(p.Number+1)*8].Occupied.Colour != p.Occupied.Colour {
			moves["sw"] = 1
		}

		//se capture
		if Board[p.Letter+1+(p.Number+1)*8].Occupied != nil && Board[p.Letter+1+(p.Number+1)*8].Occupied.Colour != p.Occupied.Colour {
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
		if Board[nextStep.Letter+nextStep.Number*8].Occupied == nil {
			return "empty"
		} else if Board[nextStep.Letter+nextStep.Number*8].Occupied.Colour == CurrentPlayer.Name {
			return "friend"
		} else {
			if Board[nextStep.Letter+nextStep.Number*8].Occupied.Kind == "king" {
				fmt.Printf("%s Player has Won the Game! \n", CurrentPlayer.Name)
				os.Exit(0)
			}
			return "capture"
		}
	}
	return "outOfBounds"
}

func SetUpBoard() {
	Board = make([]Square, 64)
	for i := 0; i < 64; i++ {
		temp := Square{Letter: i % 8, Number: int(i / 8), Occupied: nil}
		Board[i] = temp
	}
}

func createPiece(colour, kind string, Square *Square, n int) {
	temp := &Piece{Colour: colour, Kind: kind, Img: nil}
	Square.Occupied = temp
}

func SetUpPlayer() {
	white = Player{Name: "white"}
	black = Player{Name: "black"}
	for i := 0; i < 8; i++ {
		createPiece("white", "whitePawn", &Board[i+48], i)
		createPiece("black", "blackPawn", &Board[i+8], i)
	}
	createPiece("black", "rook", &Board[7], 8)
	createPiece("black", "rook", &Board[0], 9)
	createPiece("white", "rook", &Board[56], 8)
	createPiece("white", "rook", &Board[63], 9)
	createPiece("black", "knight", &Board[1], 10)
	createPiece("black", "knight", &Board[6], 11)
	createPiece("white", "knight", &Board[57], 10)
	createPiece("white", "knight", &Board[62], 11)
	createPiece("black", "bishop", &Board[2], 12)
	createPiece("black", "bishop", &Board[5], 13)
	createPiece("white", "bishop", &Board[58], 12)
	createPiece("white", "bishop", &Board[61], 13)
	createPiece("black", "queen", &Board[3], 14)
	createPiece("black", "king", &Board[4], 15)
	createPiece("white", "queen", &Board[59], 14)
	createPiece("white", "king", &Board[60], 15)
}
