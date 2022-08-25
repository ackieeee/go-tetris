package tetris

var sminoTypes = map[string][][]int{
	"type1": {
		{0, 3, 3},
		{3, 3, 0},
		{0, 0, 0},
	},
	"type2": {
		{0, 3, 0},
		{0, 3, 3},
		{0, 0, 3},
	},
	"type3": {
		{0, 0, 0},
		{0, 3, 3},
		{3, 3, 0},
	},
	"type4": {
		{3, 0, 0},
		{3, 3, 0},
		{0, 3, 0},
	},
}

type SMino struct {
	BaseMino
}

func NewSMino(x, y int, t string) Mino {
	return &SMino{
		BaseMino{x, y, t, sminoTypes},
	}
}

func (m *SMino) LeftRotate(field *Field) {
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

func (m *SMino) RightRotate(f *Field) {
}

func (m *SMino) Copy() Mino {
	return &SMino{
		BaseMino{m.X, m.Y, m.Type, m.Types},
	}
}
