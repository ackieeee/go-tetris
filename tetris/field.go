package tetris

const (
	FIELD_HEIGHT = 22
	FIELD_WIDTH  = 12
)

type Field struct {
	Tile [][]int
}

func NewField() *Field {
	field := Field{}
	field.Tile = [][]int{
		{8, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 8},
		{8, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 8},
		{8, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 8},
		{8, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 8},
		{8, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 8},
		{8, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 8},
		{8, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 8},
		{8, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 8},
		{8, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 8},
		{8, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 8},
		{8, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 8},
		{8, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 8},
		{8, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 8},
		{8, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 8},
		{8, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 8},
		{8, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 8},
		{8, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 8},
		{8, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 8},
		{8, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 8},
		{8, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 8},
		{8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8},
	}
	return &field
}

func (f *Field) DeleteLine() {
	lineNums := []int{}
	// すべて1のものの行番号を配列で保持
	for i, line := range f.Tile[:len(f.Tile)-1] {
		delFlg := true
		for _, block := range line {
			if block == 9 {
				delFlg = false
				break
			}
		}
		if delFlg {
			lineNums = append(lineNums, i)
		}
	}
	// 保持した行番号の行を削除して先頭に新しい行を追加
	for _, lineNum := range lineNums {
		f.deleteLine(lineNum)
		f.frontInsertLine()
	}
}

func (f *Field) deleteLine(lineNum int) {
	f.Tile = append(f.Tile[0:lineNum], f.Tile[lineNum+1:]...)
}

func (f *Field) frontInsertLine() {
	tile := [][]int{
		{8, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 8},
	}
	f.Tile = append(tile, f.Tile...)
}