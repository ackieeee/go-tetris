package tetris

import (
	"math/rand"
	"time"
)

func MinoCreate(x, y int) Mino {
	minos := make([]Mino, 0, 6)
	minos = append(minos, NewTMino(x, y, "type1"))
	minos = append(minos, NewSMino(x, y, "type1"))
	minos = append(minos, NewZMino(x, y, "type1"))
	minos = append(minos, NewLMino(x, y, "type1"))
	minos = append(minos, NewJMino(x, y, "type1"))
	minos = append(minos, NewIMino(x, y, "type1"))

	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(6)
	return minos[n]
}
