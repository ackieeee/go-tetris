package tetris

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image/color"
)

var Colors = []color.RGBA{
	{0xff, 0xff, 0xff, 0xff},
	{0xff, 0xff, 0x0, 0xff},
	{0xff, 0x0, 0xff, 0xff},
	{0xff, 0x0, 0x0, 0xff},
	{0x0, 0xff, 0xff, 0xff},
	{0x0, 0xff, 0x0, 0xff},
	{0x0, 0x0, 0xff, 0xff},
	{0x0, 0x0, 0x0, 0xff},
	{0xcc, 0xcc, 0xcc, 0xcc},
}

type Mino interface {
	LeftRotate(f *Field)
	RightRotate(f *Field)
	GetX() int
	SetX(x int)
	AddX(x int)
	GetY() int
	SetY(y int)
	AddY(y int)
	GetType() string
	SetType(t string)
	GetTypes(t string) [][]int
	Copy() Mino
	Move(x, y int)
	Draw(screen *ebiten.Image, blockSize float64)
}

type BaseMino struct {
	X     int
	Y     int
	Type  string
	Types map[string][][]int
	//Type [][]int
}

func (m *BaseMino) GetX() int {
	return m.X
}

func (m *BaseMino) SetX(x int) {
	m.X = x
}

func (m *BaseMino) AddX(x int) {
	m.X += x
}

func (m *BaseMino) GetY() int {
	return m.Y
}

func (m *BaseMino) SetY(y int) {
	m.Y = y
}

func (m *BaseMino) AddY(y int) {
	m.Y += y
}

func (m *BaseMino) GetType() string {
	return m.Type
}

func (m *BaseMino) GetTypes(t string) [][]int {
	return m.Types[t]
}

func (m *BaseMino) SetType(t string) {
	m.Type = t
}

func (m *BaseMino) Move(x, y int) {
	m.X += x
	m.Y += y
}

func CanMoveMinoByType(m Mino, field *Field, t string, nextX int, nextY int) bool {
	for i, line := range m.GetTypes(t) {
		for j, block := range line {
			if block == 0 {
				continue
			}
			x := j + nextX
			y := i + nextY
			if !field.CanMoveMino(m, x, y) {
				return false
			}
		}
	}
	return true
}

func (m *BaseMino) Draw(screen *ebiten.Image, blockSize float64) {
	// ミノの各ブロックごとに描画 ミノの座標＋ブロックの座標
	for i, line := range m.Types[m.Type] {
		for j, block := range line {
			if block == 0 {
				continue
			}
			x := m.X + j
			y := m.Y + i
			ebitenutil.DrawRect(screen, float64(x)*blockSize, float64(y)*blockSize, blockSize, blockSize, Colors[0])
		}
	}
}
