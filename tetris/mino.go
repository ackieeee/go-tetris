package tetris

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image/color"
	"math/rand"
	"time"
)

var minoList = [][][]int{
	// I mino
	{
		{-1, 0},
		{0, 0},
		{1, 0},
		{2, 0},
	},
	// O mino
	{
		{-1, -1},
		{-1, 0},
		{0, 0},
		{0, -1},
	},
	// S mino
	{
		{1, -1},
		{-1, 0},
		{0, 0},
		{0, -1},
	},
	// Z mino
	{
		{1, 0},
		{-1, -1},
		{0, 0},
		{0, -1},
	},
	// J mino
	{
		{1, 0},
		{0, 0},
		{-1, 0},
		{-1, -1},
	},
	// L mino
	{
		{1, -1},
		{1, 0},
		{0, 0},
		{-1, 0},
	},
	// T mino
	{
		{1, 0},
		{0, 0},
		{0, -1},
		{-1, 0},
	},
}

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

type Mino struct {
	X     int
	Y     int
	Type  [][]int
	Color int
	Shape int
}

func NewMino(x, y, shape int) *Mino {
	rand.Seed(time.Now().Unix())
	minoNum := rand.Intn(7)
	t := minoList[minoNum]
	m := Mino{
		X:     x,
		Y:     y,
		Type:  t,
		Color: minoNum,
		Shape: shape,
	}
	return &m
}

func (m *Mino) Move(x, y int) {
	m.X += x
	m.Y += y
}

func (m Mino) Draw(screen *ebiten.Image, blockSize float64) {
	// ミノの各ブロックごとに描画 ミノの座標＋ブロックの座標
	for _, block := range m.Type {
		x, y := m.Calc(block[0], block[1])
		ebitenutil.DrawRect(screen, float64(x)*blockSize, float64(y)*blockSize, blockSize, blockSize, Colors[m.Color])
	}
}

func (m *Mino) Calc(x, y int) (int, int) {
	// 座標初期値 回転なしの場合はここで定義したパラメータが使用される
	bx := x
	by := y
	// 回転処理
	for i := 0; i < m.Shape; i++ {
		xtmp := bx
		bx = -1 * by
		by = xtmp
	}
	// 座標+ブロックの座標を返却
	return m.X + bx, m.Y + by
}

func (m *Mino) Copy() Mino {
	return Mino{m.X, m.Y, m.Type, m.Color, m.Shape}
}
