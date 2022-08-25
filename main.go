package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"go-tetris/tetris"
	_ "image/png"
	"log"
)

var (
	frameCount = 0
)

const (
	BLOCK_SIZE = 35
)

type Game struct {
	field *tetris.Field
	mino  tetris.Mino
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

func (g *Game) CanMoveMino(x, y int) bool {
	return g.field.CanMoveMino(g.mino, x, y)
}

func (g *Game) DownMino() {
	if g.CanMoveMino(0, 1) {
		g.mino.Move(0, 1)
	} else {
		// 動けないので現在位置でミノを固定
		for i, line := range g.mino.GetTypes(g.mino.GetType()) {
			for j, block := range line {
				if block == 0 {
					continue
				}
				x := g.mino.GetX() + j
				y := g.mino.GetY() + i
				g.field.Tile[y][x] = block
			}
		}
		g.field.DeleteLine()
		g.mino = tetris.MinoCreate(4, 1)
	}
}

func (g *Game) Update() error {
	frameCount++
	//48フレームで1つミノを落下させる
	if frameCount == 48 {
		g.DownMino()

		frameCount = 0
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
		g.DownMino()
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) {
		if g.CanMoveMino(1, 0) {
			g.mino.Move(1, 0)
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) {
		if g.CanMoveMino(-1, 0) {
			g.mino.Move(-1, 0)
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
		// 回転させた時に壁にぶつからなければ
		//if g.CanMoveMino(0, 0, g.mino.Shape+1) {
		//	g.mino.Shape++
		//}
		g.mino.LeftRotate(g.field)
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
	//game.mino = tetris.NewMino(6, 1, 0)
	//game.mino = tetris.NewIMino(4, 1, "type6")
	game.mino = tetris.MinoCreate(4, 1)

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
