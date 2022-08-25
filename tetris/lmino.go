package tetris

var lminoTypes = map[string][][]int{
	"type1": {
		{0, 0, 6},
		{6, 6, 6},
		{0, 0, 0},
	},
	"type2": {
		{0, 6, 0},
		{0, 6, 0},
		{0, 6, 6},
	},
	"type3": {
		{0, 0, 0},
		{6, 6, 6},
		{6, 0, 0},
	},
	"type4": {
		{6, 6, 0},
		{0, 6, 0},
		{0, 6, 0},
	},
}

type LMino struct {
	BaseMino
}

func NewLMino(x, y int, t string) Mino {
	return &LMino{
		BaseMino{x, y, t, lminoTypes},
	}
}

func (m *LMino) LeftRotate(field *Field) {
	switch m.Type {
	case "type1":
		if CanMoveMinoByType(m, field, "type4", 0, 0) {
			m.Type = "type4"
			return
		}
	case "type2":
		if CanMoveMinoByType(m, field, "type1", 0, 0) {
			m.Type = "type1"
			return
		}
	case "type3":
		if CanMoveMinoByType(m, field, "type2", 0, 0) {
			m.Type = "type2"
			return
		}
	case "type4":
		if CanMoveMinoByType(m, field, "type3", 0, 0) {
			m.Type = "type3"
			return
		}
	}
}

func (m *LMino) RightRotate(f *Field) {
}

func (m *LMino) Copy() Mino {
	return &LMino{
		BaseMino{m.X, m.Y, m.Type, m.Types},
	}
}
