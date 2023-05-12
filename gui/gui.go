package gui

import (
	"bytes"
	"image"
	"image/color"
	"image/draw"
	_ "image/png"
	"log"
	"time"
	"math/rand"

	d "golang.org/x/image/draw"

	"github.com/hajimehoshi/ebiten"
	//"github.com/hajimehoshi/ebiten/v2"
	"example/pairProgramming/pieces"
)

const (
	size = 640
	squareSize = size/8
)

var (
	whitePawn *ebiten.Image
	whiteRook *ebiten.Image
	whiteKnight *ebiten.Image
	whiteBishop *ebiten.Image
	whiteKing *ebiten.Image
	whiteQueen *ebiten.Image
	blackPawn *ebiten.Image
	blackRook *ebiten.Image
	blackKnight *ebiten.Image
	blackBishop *ebiten.Image
	blackKing *ebiten.Image
	blackQueen *ebiten.Image
	pSlice []Piece
	moves map[string][]string
)

func init() {
	scaler := d.NearestNeighbor
	img, _ , err := image.Decode(bytes.NewReader(pieces.Whitepawn_png))
	if err != nil {
		log.Fatal(err)
	}
	img = scaleTo(img, image.Rect(0,0,squareSize,squareSize), scaler)
	whitePawn, err = ebiten.NewImageFromImage(img, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	img, _ , err = image.Decode(bytes.NewReader(pieces.Whiterook_png))
	if err != nil {
		log.Fatal(err)
	}
	img = scaleTo(img, image.Rect(0,0,squareSize,squareSize), scaler)
	whiteRook, err = ebiten.NewImageFromImage(img, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	img, _ , err = image.Decode(bytes.NewReader(pieces.Whitebishop_png))
	if err != nil {
		log.Fatal(err)
	}
	img = scaleTo(img, image.Rect(0,0,squareSize,squareSize), scaler)
	whiteBishop, err = ebiten.NewImageFromImage(img, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	img, _ , err = image.Decode(bytes.NewReader(pieces.Whiteknight_png))
	if err != nil {
		log.Fatal(err)
	}
	img = scaleTo(img, image.Rect(0,0,squareSize,squareSize), scaler)
	whiteKnight, err = ebiten.NewImageFromImage(img, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	img, _ , err = image.Decode(bytes.NewReader(pieces.Whiteking_png))
	if err != nil {
		log.Fatal(err)
	}
	img = scaleTo(img, image.Rect(0,0,squareSize,squareSize), scaler)
	whiteKing, err = ebiten.NewImageFromImage(img, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	img, _ , err = image.Decode(bytes.NewReader(pieces.Whitequeen_png))
	if err != nil {
		log.Fatal(err)
	}
	img = scaleTo(img, image.Rect(0,0,squareSize,squareSize), scaler)
	whiteQueen, err = ebiten.NewImageFromImage(img, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	img, _ , err = image.Decode(bytes.NewReader(pieces.Blackpawn_png))
	if err != nil {
		log.Fatal(err)
	}
	img = scaleTo(img, image.Rect(0,0,squareSize,squareSize), scaler)
	blackPawn, err = ebiten.NewImageFromImage(img, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	img, _ , err = image.Decode(bytes.NewReader(pieces.Blackrook_png))
	if err != nil {
		log.Fatal(err)
	}
	img = scaleTo(img, image.Rect(0,0,squareSize,squareSize), scaler)
	blackRook, err = ebiten.NewImageFromImage(img, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	img, _ , err = image.Decode(bytes.NewReader(pieces.Blackbishop_png))
	if err != nil {
		log.Fatal(err)
	}
	img = scaleTo(img, image.Rect(0,0,squareSize,squareSize), scaler)
	blackBishop, err = ebiten.NewImageFromImage(img, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	img, _ , err = image.Decode(bytes.NewReader(pieces.Blackknight_png))
	if err != nil {
		log.Fatal(err)
	}
	img = scaleTo(img, image.Rect(0,0,squareSize,squareSize), scaler)
	blackKnight, err = ebiten.NewImageFromImage(img, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	img, _ , err = image.Decode(bytes.NewReader(pieces.Blackking_png))
	if err != nil {
		log.Fatal(err)
	}
	img = scaleTo(img, image.Rect(0,0,squareSize,squareSize), scaler)
	blackKing, err = ebiten.NewImageFromImage(img, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	img, _ , err = image.Decode(bytes.NewReader(pieces.Blackqueen_png))
	if err != nil {
		log.Fatal(err)
	}
	img = scaleTo(img, image.Rect(0,0,squareSize,squareSize), scaler)
	blackQueen, err = ebiten.NewImageFromImage(img, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	pSlice = []Piece{
		{img: whitePawn, kind: "whitePawn", x: 0, y: 6},
		{img: whitePawn, kind: "whitePawn", x: 1, y: 6},
		{img: whitePawn, kind: "whitePawn", x: 2, y: 6},
		{img: whitePawn, kind: "whitePawn", x: 3, y: 6},
		{img: whitePawn, kind: "whitePawn", x: 4, y: 6},
		{img: whitePawn, kind: "whitePawn", x: 5, y: 6},
		{img: whitePawn, kind: "whitePawn", x: 6, y: 6},
		{img: whitePawn, kind: "whitePawn", x: 7, y: 6},
		{img: whiteRook, kind: "rook", x: 0, y: 7},
		{img: whiteRook, kind: "rook", x: 7, y: 7},
		{img: whiteKnight, kind: "knight", x: 1, y: 7},
		{img: whiteKnight, kind: "knight", x: 6, y: 7},
		{img: whiteBishop, kind: "bishop", x: 2, y: 7},
		{img: whiteBishop, kind: "bishop", x: 5, y: 7},
		{img: whiteKing, kind: "king", x: 4, y: 7},
		{img: whiteQueen, kind: "queen", x: 3, y: 7},
		{img: blackPawn, kind: "blackPawn", x: 0, y: 1},
		{img: blackPawn, kind: "blackPawn", x: 1, y: 1},
		{img: blackPawn, kind: "blackPawn", x: 2, y: 1},
		{img: blackPawn, kind: "blackPawn", x: 3, y: 1},
		{img: blackPawn, kind: "blackPawn", x: 4, y: 1},
		{img: blackPawn, kind: "blackPawn", x: 5, y: 1},
		{img: blackPawn, kind: "blackPawn", x: 6, y: 1},
		{img: blackPawn, kind: "blackPawn", x: 7, y: 1},
		{img: blackRook, kind: "rook", x: 0, y: 0},
		{img: blackRook, kind: "rook", x: 7, y: 0},
		{img: blackKnight, kind: "knight", x: 1, y: 0},
		{img: blackKnight, kind: "knight", x: 6, y: 0},
		{img: blackBishop, kind: "bishop", x: 2, y: 0},
		{img: blackBishop, kind: "bishop", x: 5, y: 0},
		{img: blackKing, kind: "king", x: 4, y: 0},
		{img: blackQueen, kind: "queen", x: 3, y: 0},
	}	

	moves = make(map[string][]string)
	moves["whitePawn"] = []string{"n"}
	moves["blackPawn"] = []string{"s"}
	moves["rook"] = []string{"s", "n", "w", "e"}
	moves["bishop"] = []string{"sw", "nw", "nw", "se"}
	moves["king"] = []string{"n","ne","e","se","s","sw","w","nw"}
	moves["queen"] = []string{"n","ne","e","se","s","sw","w","nw"}

}

func (g *Game) updateGroundImage(ground *ebiten.Image) {
//	op := &ebiten.DrawImageOptions{}
//	ground.DrawImage(whitePawn, op)
//	op.GeoM.Translate(float64(0), float64(size-squareSize))
//	ground.DrawImage(whitePawn, op)
//	op.GeoM.Reset()
}

func drawSquare(xStart, yStart int, img *ebiten.Image, color color.Color) {
	for y := yStart; y < yStart+squareSize; y++ {
		for x := xStart; x < xStart+squareSize; x++ {
			img.Set(x, y, color)
		}
	}
}

func (g *Game) drawGroundImage(screen, ground *ebiten.Image) {
	toggle := true
	col := color.White
	for y := 0; y < size; y += squareSize {
		toggle = !toggle
		if toggle {
			col = color.Black
		} else {
			col = color.White
		}
		for x := 0; x < size; x += squareSize {
			drawSquare(x,y,ground, col)	
			toggle = !toggle
			if toggle {
				col = color.Black
			} else {
				col = color.White
			}
		}
	}
	op := &ebiten.DrawImageOptions{}
	for _, piece := range pSlice {
		op.GeoM.Translate(float64(piece.x*squareSize), float64(piece.y*squareSize))
		ground.DrawImage(piece.img, op)
		op.GeoM.Reset()
	}
	op.GeoM.Reset()
	screen.DrawImage(ground, &ebiten.DrawImageOptions{})
}

func NewGame() *Game {
	img, err := ebiten.NewImage(size,size,ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	g := &Game{groundImage: img}
	return g
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.updateGroundImage(g.groundImage)
	g.drawGroundImage(screen, g.groundImage)
}

type Piece struct {
	img *ebiten.Image
	kind string
	x int
	y int
}

type Game struct {
	groundImage *ebiten.Image
}

func (g *Game) Update(img *ebiten.Image) error {
	// Manipulate the player by the input.
	time.Sleep(time.Second*2)
	piece := rand.Intn(32)
	x := rand.Intn(8)
	y := rand.Intn(8)
	pSlice[piece].x = x
	pSlice[piece].y = y
	return nil
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return size, size
}

func GuiRun() {
	ebiten.SetWindowSize(size,size)
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

