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
	frameCount = 0
)

const (
	BLOCK_SIZE = 35
	NEXT_POS   = 650
)

type Game struct {
	field *tetris.Field
	mino  tetris.Mino
	hold  tetris.Mino
	next  tetris.Mino
}

func (g *Game) drawField(screen *ebiten.Image) {
	g.drawHoldField(screen)
	g.drawNextField(screen)
	for i := 0; i < len(g.field.Tile); i++ {
		line := g.field.Tile[i]
		for j := 0; j < len(line); j++ {
			block := line[j]
			if block != 9 {
				x := j + tetris.FIELD_POS
				y := i
				ebitenutil.DrawRect(screen, float64(BLOCK_SIZE*x), float64(BLOCK_SIZE*y), BLOCK_SIZE, BLOCK_SIZE, tetris.Colors[block])
			}
		}
	}
}

func (g *Game) drawHoldField(screen *ebiten.Image) {
	ebitenutil.DrawRect(screen, float64(BLOCK_SIZE), float64(BLOCK_SIZE), 5, BLOCK_SIZE*5, color.White)
	ebitenutil.DrawRect(screen, float64(BLOCK_SIZE), float64(BLOCK_SIZE), BLOCK_SIZE*5, 5, color.White)
	ebitenutil.DrawRect(screen, float64(BLOCK_SIZE), float64(BLOCK_SIZE*6), BLOCK_SIZE*5, 5, color.White)
	ebitenutil.DrawRect(screen, float64(BLOCK_SIZE*6), float64(BLOCK_SIZE), 5, BLOCK_SIZE*5+5, color.White)
	if g.hold != nil {
		g.hold.Draw(screen, BLOCK_SIZE)
	}
}

func (g *Game) drawNextField(screen *ebiten.Image) {
	ebitenutil.DrawRect(screen, float64(BLOCK_SIZE+NEXT_POS), float64(BLOCK_SIZE), 5, BLOCK_SIZE*5, color.White)
	ebitenutil.DrawRect(screen, float64(BLOCK_SIZE+NEXT_POS), float64(BLOCK_SIZE), BLOCK_SIZE*5, 5, color.White)
	ebitenutil.DrawRect(screen, float64(BLOCK_SIZE+NEXT_POS), float64(BLOCK_SIZE*6), BLOCK_SIZE*5, 5, color.White)
	ebitenutil.DrawRect(screen, float64(BLOCK_SIZE*6+NEXT_POS), float64(BLOCK_SIZE), 5, BLOCK_SIZE*5+5, color.White)
	g.next.Draw(screen, BLOCK_SIZE)
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
				x := g.mino.GetX() + j - tetris.FIELD_POS
				y := g.mino.GetY() + i
				g.field.Tile[y][x] = block
			}
		}
		g.field.DeleteLine()
		g.mino = tetris.MinoCreate(tetris.FIELD_POS+4, 1)
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
	// ホールドと切り替え
	if inpututil.IsKeyJustPressed(ebiten.KeyH) {
		if g.hold == nil {
			g.hold = g.mino.Copy()
			g.hold.SetType("type1")
			g.hold.SetX(2)
			g.hold.SetY(2)
			g.mino = tetris.MinoCreate(tetris.FIELD_POS+4, 1)
		} else {
			copyHoldMino := g.hold.Copy()
			copyMino := g.mino.Copy()

			x := copyMino.GetX()
			y := copyMino.GetY()

			g.hold = copyMino
			g.hold.SetType("type1")
			g.hold.SetX(2)
			g.hold.SetY(2)

			g.mino = copyHoldMino
			g.mino.SetX(x)
			g.mino.SetY(y)
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.drawField(screen)
	g.mino.Draw(screen, BLOCK_SIZE)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 960, 1280
}

func main() {
	ebiten.SetWindowSize(960, 1280)
	ebiten.SetWindowTitle("テトリス")

	game := &Game{}
	game.field = tetris.NewField()
	//game.mino = tetris.NewMino(6, 1, 0)
	//game.mino = tetris.NewIMino(4, 1, "type6")
	game.mino = tetris.MinoCreate(tetris.FIELD_POS+4, 1)
	game.next = tetris.MinoCreate(NEXT_POS+2, 1)

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
