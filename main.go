package main

import (
	"bytes"
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
	for i := 0; i < 64; i++ {
		letter, number, kind, colour := gamelogic.GetSquare(i)
		if (letter+number)%2 == 0 {
			col = color.RGBA{R: 103, G: 51, B: 20, A: 127}
		} else {
			col = color.RGBA{R: 249, G: 172, B: 113, A: 127}
		}
		drawSquare(letter*squareSize, number*squareSize, ground, col)
		if kind != "" {
			op.GeoM.Translate(float64(letter*squareSize), float64(number*squareSize))
			ground.DrawImage(getImage(kind, colour), op)
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
	gamelogic.Move()
	time.Sleep(time.Second * 2)
	return nil
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return size, size
}

func main() {
	gamelogic.SetUpBoard()
	gamelogic.SetUpPlayer()
	ebiten.SetWindowSize(size, size)
	ebiten.SetWindowTitle("Chess")
	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}

func scaleTo(src image.Image, rect image.Rectangle, sc d.Scaler) image.Image {
	dst := image.NewRGBA(rect)
	sc.Scale(dst, rect, src, src.Bounds(), draw.Over, nil)
	return dst
}

func getImage(kind, colour string) *ebiten.Image {
	switch kind {
	case "whitePawn":
		return WhitePawn
	case "blackPawn":
		return BlackPawn
	case "rook":
		if colour == "white" {
			return WhiteRook
		} else {
			return BlackRook
		}
	case "bishop":
		if colour == "white" {
			return WhiteBishop
		} else {
			return BlackBishop
		}
	case "knight":
		if colour == "white" {
			return WhiteKnight
		} else {
			return BlackKnight
		}
	case "queen":
		if colour == "white" {
			return WhiteQueen
		} else {
			return BlackQueen
		}
	case "king":
		if colour == "white" {
			return WhiteKing
		} else {
			return BlackKing
		}
	default:
		log.Fatal("Error in getImage")
		return nil
	}
}
