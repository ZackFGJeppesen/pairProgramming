package main

import (
	"log"
	"math/rand"
	"example/pairProgramming/gui"
)

type Player struct {
	name   string
	pieces []Piece
}

type Piece struct {
	colour  string
	name    string
	kind    string
	postion *Point
	moves   []Point
}

type Point struct {
	letter   int
	number   int
	occupied *Piece
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

var Board []Point //= createBoard()
var currentPlayer *Player = &white

func main() {
	gui.GuiRun()
	for {
		p := randPickPiece(currentPlayer)
		point := randMove(legalMoves(p))
		Board[p.postion.letter*8+p.postion.number].occupied = nil
		Board[point.letter*8+point.number].occupied = &p
		swap(currentPlayer)
	}
}

func randPickPiece(currentPlayer *Player) Piece {
	return currentPlayer.pieces[rand.Intn(len(currentPlayer.pieces))]
}

func randMove(returnMoves []Point) Point {
	return returnMoves[rand.Intn(len(returnMoves))]
}

func swap(currentPlayer *Player) {
	if currentPlayer.name == "white" {
		currentPlayer = &black
	} else {
		currentPlayer = &white
	}
}

func legalMoves(p Piece) []Point {
	var returnMoves []Point
	if p.kind != "knight" {
		for dir, dist := range direction(p) {

			var nextStep Point
		loop:
			for step := 1; step <= dist; step++ {
				switch dir {
				case "nw":
					nextStep = Point{letter: p.postion.letter - step, number: p.postion.number + step}
				case "n":
					nextStep = Point{letter: p.postion.letter, number: p.postion.number + step}
				case "ne":
					nextStep = Point{letter: p.postion.letter + step, number: p.postion.number + step}
				case "w":
					nextStep = Point{letter: p.postion.letter - step, number: p.postion.number}
				case "e":
					nextStep = Point{letter: p.postion.letter + step, number: p.postion.number}
				case "sw":
					nextStep = Point{letter: p.postion.letter - step, number: p.postion.number - step}
				case "s":
					nextStep = Point{letter: p.postion.letter, number: p.postion.number - step}
				case "se":
					nextStep = Point{letter: p.postion.letter + step, number: p.postion.number - step}
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
		tempP := Point{letter: p.postion.letter - 2, number: p.postion.number + 1}
		s := check(tempP)
		if s == "empty" || s == "capture" {
			returnMoves = append(returnMoves, tempP)
		}

		tempP = Point{letter: p.postion.letter - 1, number: p.postion.number + 2}
		s = check(tempP)
		if s == "empty" || s == "capture" {
			returnMoves = append(returnMoves, tempP)
		}

		tempP = Point{letter: p.postion.letter + 1, number: p.postion.number + 2}
		s = check(tempP)
		if s == "empty" || s == "capture" {
			returnMoves = append(returnMoves, tempP)
		}

		tempP = Point{letter: p.postion.letter + 2, number: p.postion.number + 1}
		s = check(tempP)
		if s == "empty" || s == "capture" {
			returnMoves = append(returnMoves, tempP)
		}

		tempP = Point{letter: p.postion.letter + 2, number: p.postion.number - 1}
		s = check(tempP)
		if s == "empty" || s == "capture" {
			returnMoves = append(returnMoves, tempP)
		}

		tempP = Point{letter: p.postion.letter + 1, number: p.postion.number - 2}
		s = check(tempP)
		if s == "empty" || s == "capture" {
			returnMoves = append(returnMoves, tempP)
		}

		tempP = Point{letter: p.postion.letter - 1, number: p.postion.number - 2}
		s = check(tempP)
		if s == "empty" || s == "capture" {
			returnMoves = append(returnMoves, tempP)
		}

		tempP = Point{letter: p.postion.letter - 2, number: p.postion.number - 1}
		s = check(tempP)
		if s == "empty" || s == "capture" {
			returnMoves = append(returnMoves, tempP)
		}
	}
	return returnMoves
}

func direction(p Piece) map[string]int {
	moves := make(map[string]int)
	switch p.kind {
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
	if 0 <= nextStep.letter && nextStep.letter <= 8 && 0 <= nextStep.number && nextStep.number <= 8 {
		if Board[nextStep.letter*8+nextStep.number].occupied == nil {
			return "empty"
		} else if Board[nextStep.letter*8+nextStep.number].occupied.colour == currentPlayer.name {
			return "friend"
		} else {
			return "capture"
		}
	}
	return "outOfBounds"
}
