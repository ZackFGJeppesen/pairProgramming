package main

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	_ "image/png"
	"log"
	"time"

	d "golang.org/x/image/draw"

	"example/pairProgramming/gamelogic"
	"example/pairProgramming/pieces"

	"github.com/hajimehoshi/ebiten"
)

const (
	size       = 640
	squareSize = size / 8
)

var (
	WhitePawn    *ebiten.Image
	WhiteRook    *ebiten.Image
	WhiteKnight  *ebiten.Image
	WhiteBishop  *ebiten.Image
	WhiteKing    *ebiten.Image
	WhiteQueen   *ebiten.Image
	BlackPawn    *ebiten.Image
	BlackRook    *ebiten.Image
	BlackKnight  *ebiten.Image
	BlackBishop  *ebiten.Image
	BlackKing    *ebiten.Image
	BlackQueen   *ebiten.Image
	moves        map[string][]string
	notFirstMove bool
)

func init() {
	scaler := d.NearestNeighbor
	img, _, err := image.Decode(bytes.NewReader(pieces.Whitepawn_png))
	if err != nil {
		log.Fatal(err)
	}
	img = scaleTo(img, image.Rect(0, 0, squareSize, squareSize), scaler)
	WhitePawn, err = ebiten.NewImageFromImage(img, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	img, _, err = image.Decode(bytes.NewReader(pieces.Whiterook_png))
	if err != nil {
		log.Fatal(err)
	}
	img = scaleTo(img, image.Rect(0, 0, squareSize, squareSize), scaler)
	WhiteRook, err = ebiten.NewImageFromImage(img, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	img, _, err = image.Decode(bytes.NewReader(pieces.Whitebishop_png))
	if err != nil {
		log.Fatal(err)
	}
	img = scaleTo(img, image.Rect(0, 0, squareSize, squareSize), scaler)
	WhiteBishop, err = ebiten.NewImageFromImage(img, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	img, _, err = image.Decode(bytes.NewReader(pieces.Whiteknight_png))
	if err != nil {
		log.Fatal(err)
	}
	img = scaleTo(img, image.Rect(0, 0, squareSize, squareSize), scaler)
	WhiteKnight, err = ebiten.NewImageFromImage(img, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	img, _, err = image.Decode(bytes.NewReader(pieces.Whiteking_png))
	if err != nil {
		log.Fatal(err)
	}
	img = scaleTo(img, image.Rect(0, 0, squareSize, squareSize), scaler)
	WhiteKing, err = ebiten.NewImageFromImage(img, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	img, _, err = image.Decode(bytes.NewReader(pieces.Whitequeen_png))
	if err != nil {
		log.Fatal(err)
	}
	img = scaleTo(img, image.Rect(0, 0, squareSize, squareSize), scaler)
	WhiteQueen, err = ebiten.NewImageFromImage(img, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	img, _, err = image.Decode(bytes.NewReader(pieces.Blackpawn_png))
	if err != nil {
		log.Fatal(err)
	}
	img = scaleTo(img, image.Rect(0, 0, squareSize, squareSize), scaler)
	BlackPawn, err = ebiten.NewImageFromImage(img, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	img, _, err = image.Decode(bytes.NewReader(pieces.Blackrook_png))
	if err != nil {
		log.Fatal(err)
	}
	img = scaleTo(img, image.Rect(0, 0, squareSize, squareSize), scaler)
	BlackRook, err = ebiten.NewImageFromImage(img, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	img, _, err = image.Decode(bytes.NewReader(pieces.Blackbishop_png))
	if err != nil {
		log.Fatal(err)
	}
	img = scaleTo(img, image.Rect(0, 0, squareSize, squareSize), scaler)
	BlackBishop, err = ebiten.NewImageFromImage(img, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	img, _, err = image.Decode(bytes.NewReader(pieces.Blackknight_png))
	if err != nil {
		log.Fatal(err)
	}
	img = scaleTo(img, image.Rect(0, 0, squareSize, squareSize), scaler)
	BlackKnight, err = ebiten.NewImageFromImage(img, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	img, _, err = image.Decode(bytes.NewReader(pieces.Blackking_png))
	if err != nil {
		log.Fatal(err)
	}
	img = scaleTo(img, image.Rect(0, 0, squareSize, squareSize), scaler)
	BlackKing, err = ebiten.NewImageFromImage(img, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	img, _, err = image.Decode(bytes.NewReader(pieces.Blackqueen_png))
	if err != nil {
		log.Fatal(err)
	}
	img = scaleTo(img, image.Rect(0, 0, squareSize, squareSize), scaler)
	BlackQueen, err = ebiten.NewImageFromImage(img, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	// PSlice = []piece{
	// 	{img: WhitePawn, kind: "whitePawn", x: 0, y: 1},
	// 	{img: WhitePawn, kind: "whitePawn", x: 1, y: 1},
	// 	{img: WhitePawn, kind: "whitePawn", x: 2, y: 1},
	// 	{img: WhitePawn, kind: "whitePawn", x: 3, y: 1},
	// 	{img: WhitePawn, kind: "whitePawn", x: 4, y: 1},
	// 	{img: WhitePawn, kind: "whitePawn", x: 5, y: 1},
	// 	{img: WhitePawn, kind: "whitePawn", x: 6, y: 1},
	// 	{img: WhitePawn, kind: "whitePawn", x: 7, y: 1},
	// 	{img: WhiteRook, kind: "rook", x: 0, y: 0},
	// 	{img: WhiteRook, kind: "rook", x: 7, y: 0},
	// 	{img: WhiteKnight, kind: "knight", x: 1, y: 0},
	// 	{img: WhiteKnight, kind: "knight", x: 6, y: 0},
	// 	{img: WhiteBishop, kind: "bishop", x: 2, y: 0},
	// 	{img: WhiteBishop, kind: "bishop", x: 5, y: 0},
	// 	{img: WhiteKing, kind: "king", x: 4, y: 0},
	// 	{img: WhiteQueen, kind: "queen", x: 3, y: 0},
	// 	{img: BlackPawn, kind: "blackPawn", x: 0, y: 6},
	// 	{img: BlackPawn, kind: "blackPawn", x: 1, y: 6},
	// 	{img: BlackPawn, kind: "blackPawn", x: 2, y: 6},
	// 	{img: BlackPawn, kind: "blackPawn", x: 3, y: 6},
	// 	{img: BlackPawn, kind: "blackPawn", x: 4, y: 6},
	// 	{img: BlackPawn, kind: "blackPawn", x: 5, y: 6},
	// 	{img: BlackPawn, kind: "blackPawn", x: 6, y: 6},
	// 	{img: BlackPawn, kind: "blackPawn", x: 7, y: 6},
	// 	{img: BlackRook, kind: "rook", x: 0, y: 7},
	// 	{img: BlackRook, kind: "rook", x: 7, y: 7},
	// 	{img: BlackKnight, kind: "knight", x: 1, y: 7},
	// 	{img: BlackKnight, kind: "knight", x: 6, y: 7},
	// 	{img: BlackBishop, kind: "bishop", x: 2, y: 7},
	// 	{img: BlackBishop, kind: "bishop", x: 5, y: 7},
	// 	{img: BlackKing, kind: "king", x: 4, y: 7},
	// 	{img: BlackQueen, kind: "queen", x: 3, y: 7},
	// }

	moves = make(map[string][]string)
	moves["whitePawn"] = []string{"n"}
	moves["blackPawn"] = []string{"s"}
	moves["rook"] = []string{"s", "n", "w", "e"}
	moves["bishop"] = []string{"sw", "nw", "nw", "se"}
	moves["king"] = []string{"n", "ne", "e", "se", "s", "sw", "w", "nw"}
	moves["queen"] = []string{"n", "ne", "e", "se", "s", "sw", "w", "nw"}

}

func drawSquare(xStart, yStart int, img *ebiten.Image, color color.Color) {
	for y := yStart; y < yStart+squareSize; y++ {
		for x := xStart; x < xStart+squareSize; x++ {
			img.Set(x, y, color)
		}
	}
}

func (g *Game) drawGroundImage(screen, ground *ebiten.Image) {
	var col color.RGBA
	op := &ebiten.DrawImageOptions{}
	for _, point := range gamelogic.Board {
		if (point.Letter+point.Number)%2 == 0 {
			col = color.RGBA{R: 103, G: 51, B: 20, A: 127}
		} else {
			col = color.RGBA{R: 249, G: 172, B: 113, A: 127}
		}
		drawSquare(point.Letter*squareSize, point.Number*squareSize, ground, col)
		if point.Occupied != nil {
			op.GeoM.Translate(float64(point.Letter*squareSize), float64(point.Number*squareSize))
			ground.DrawImage(point.Occupied.Img, op)
			op.GeoM.Reset()
		}
	}
	op.GeoM.Reset()
	screen.DrawImage(ground, &ebiten.DrawImageOptions{})
}

func NewGame() *Game {
	img, err := ebiten.NewImage(size, size, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	g := &Game{groundImage: img}
	return g
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.drawGroundImage(screen, g.groundImage)
}

type piece struct {
	img  *ebiten.Image
	kind string
	x    int
	y    int
}

type Game struct {
	groundImage *ebiten.Image
}

func (g *Game) Update(img *ebiten.Image) error {
	// Manipulate the player by the input.
	// for i, piece := range gamelogic.Board {
	// 	fmt.Printf("Square %d is %v", i, piece.Occupied)
	// }
	var lm map[gamelogic.Square]gamelogic.Square
	var squaresWithPieces []gamelogic.Square
	for {
		squaresWithPieces = gamelogic.RandPickSquare()
		lm = gamelogic.LegalMoves(squaresWithPieces)
		if len(lm) != 0 {
			break
		}
	}

	s, point := gamelogic.RandMove(lm)
	fmt.Printf("Piece %s is at Square (%d, %d) \n", s.Occupied.Kind, s.Letter, s.Number)
	fmt.Println(lm)
	gamelogic.Board[s.Letter+s.Number*8].Occupied = nil
	gamelogic.Board[point.Letter+point.Number*8].Occupied = s.Occupied
	gamelogic.Swap()
	point = gamelogic.Board[point.Letter+point.Number*8]
	fmt.Printf("Piece %s is now at Square (%d, %d) \n", point.Occupied.Kind, point.Letter, point.Number)
	time.Sleep(time.Second * 2)
	return nil
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return size, size
}

func addImg() {
	for i := range gamelogic.Board {
		if gamelogic.Board[i].Occupied != nil {
			switch gamelogic.Board[i].Occupied.Kind {
			case "whitePawn":
				gamelogic.Board[i].Occupied.Img = WhitePawn
			case "blackPawn":
				gamelogic.Board[i].Occupied.Img = BlackPawn
			case "rook":
				if gamelogic.Board[i].Occupied.Colour == "white" {
					gamelogic.Board[i].Occupied.Img = WhiteRook
				} else {
					gamelogic.Board[i].Occupied.Img = BlackRook
				}
			case "bishop":
				if gamelogic.Board[i].Occupied.Colour == "white" {
					gamelogic.Board[i].Occupied.Img = WhiteBishop
				} else {
					gamelogic.Board[i].Occupied.Img = BlackBishop
				}
			case "knight":
				if gamelogic.Board[i].Occupied.Colour == "white" {
					gamelogic.Board[i].Occupied.Img = WhiteKnight
				} else {
					gamelogic.Board[i].Occupied.Img = BlackKnight
				}
			case "queen":
				if gamelogic.Board[i].Occupied.Colour == "white" {
					gamelogic.Board[i].Occupied.Img = WhiteQueen
				} else {
					gamelogic.Board[i].Occupied.Img = BlackQueen
				}
			case "king":
				if gamelogic.Board[i].Occupied.Colour == "white" {
					gamelogic.Board[i].Occupied.Img = WhiteKing
				} else {
					gamelogic.Board[i].Occupied.Img = BlackKing
				}
			}
		}
	}
}

func main() {
	gamelogic.SetUpBoard()
	gamelogic.SetUpPlayer()
	ebiten.SetWindowSize(size, size)
	ebiten.SetWindowTitle("Chess")
	addImg()
	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}

func scaleTo(src image.Image, rect image.Rectangle, sc d.Scaler) image.Image {
	dst := image.NewRGBA(rect)
	sc.Scale(dst, rect, src, src.Bounds(), draw.Over, nil)
	return dst
}
