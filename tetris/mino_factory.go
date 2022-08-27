package tetris

import (
	"math/rand"
	"time"
)

var stacks = []int{}

func init() {
	rand.Seed(time.Now().UnixNano())
	stacks = append(stacks, rand.Perm(6)...)
	stacks = append(stacks, rand.Perm(6)...)
}

func MinoCreate(x, y int) Mino {
	minos := make([]Mino, 0, 6)
	minos = append(minos, NewTMino(x, y, "type1"))
	minos = append(minos, NewSMino(x, y, "type1"))
	minos = append(minos, NewZMino(x, y, "type1"))
	minos = append(minos, NewLMino(x, y, "type1"))
	minos = append(minos, NewJMino(x, y, "type1"))
	minos = append(minos, NewIMino(x, y, "type1"))

	next := stacks[0]
	stacks = stacks[1:]

	if len(stacks) < 7 {
		stacks = append(stacks, rand.Perm(6)...)
	}

	return minos[next]
}
