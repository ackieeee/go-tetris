package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"go-tetris/tetris"
	"image/color"
	_ "image/png"
	"log"
)

var (
	keys     []string
	runImage *ebiten.Image
	colors   = []color.RGBA{
		{0xff, 0xff, 0xff, 0xff},
		{0xff, 0xff, 0x0, 0xff},
		{0xff, 0x0, 0xff, 0xff},
		{0xff, 0x0, 0x0, 0xff},
		{0x0, 0xff, 0xff, 0xff},
		{0x0, 0xff, 0x0, 0xff},
		{0x0, 0x0, 0xff, 0xff},
		{0x0, 0x0, 0x0, 0xff},
		{0xcc, 0xcc, 0xcc, 0xcc},
		{0x0, 0x0, 0x0, 0x0},
	}
	gray       = color.RGBA{0xcc, 0xcc, 0xcc, 0xcc}
	frameCount = 0
)

const (
	BLOCK_SIZE = 35
)

type Game struct {
	field *tetris.Field
	mino  *tetris.Mino
}

func (g *Game) drawField(screen *ebiten.Image) {
	for i := 0; i < len(g.field.Tile); i++ {
		line := g.field.Tile[i]
		for j := 0; j < len(line); j++ {
			block := line[j]
			if block != 9 {
				ebitenutil.DrawRect(screen, float64(BLOCK_SIZE*j), float64(BLOCK_SIZE*i), BLOCK_SIZE, BLOCK_SIZE, tetris.Colors[block])
			}
		}
	}
}

func (g *Game) CanMoveMino(x, y, shape int) bool {
	featureMino := g.mino.Copy()
	featureMino.X += x
	featureMino.Y += y
	featureMino.Shape = shape

	// 衝突検知
	for _, t := range featureMino.Type {
		x, y := featureMino.Calc(t[0], t[1])
		if g.field.Tile[y][x] != 9 {
			return false
		}
	}
	return true
}

func (g *Game) DownMino() {
	if g.CanMoveMino(0, 1, g.mino.Shape) {
		g.mino.Move(0, 1)
	} else {
		// 動けないので現在位置でミノを固定
		for _, t := range g.mino.Type {
			// 回転を含む現在のブロックの座標を取得
			x, y := g.mino.Calc(t[0], t[1])
			// 各ブロックの色に空白を部分を更新(ブロック固定処理)
			g.field.Tile[y][x] = g.mino.Color
		}
		g.field.DeleteLine()
		g.mino = tetris.NewMino(5, 1, 0)
	}
}

func (g *Game) Update() error {
	frameCount++
	// 48フレームで1つミノを落下させる
	if frameCount == 48 {
		g.DownMino()
		frameCount = 0
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
		if g.CanMoveMino(0, 1, g.mino.Shape) {
			g.mino.Move(0, 1)
		} else {
			// 動けないので現在位置でミノを固定
			g.DownMino()
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) {
		if g.CanMoveMino(1, 0, g.mino.Shape) {
			g.mino.Move(1, 0)
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) {
		if g.CanMoveMino(-1, 0, g.mino.Shape) {
			g.mino.Move(-1, 0)
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
		// 回転させた時に壁にぶつからなければ
		if g.CanMoveMino(0, 0, g.mino.Shape+1) {
			g.mino.Shape++
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.drawField(screen)
	g.mino.Draw(screen, BLOCK_SIZE)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 960
}

func main() {
	ebiten.SetWindowSize(640, 960)
	ebiten.SetWindowTitle("テトリス")

	game := &Game{}
	game.field = tetris.NewField()
	game.mino = tetris.NewMino(5, 1, 0)

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
