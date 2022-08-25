package tetris

var zminoTypes = map[string][][]int{
	"type1": {
		{5, 5, 0},
		{0, 5, 5},
		{0, 0, 0},
	},
	"type2": {
		{0, 0, 5},
		{0, 5, 5},
		{0, 5, 0},
	},
	"type3": {
		{0, 0, 0},
		{5, 5, 0},
		{0, 5, 5},
	},
	"type4": {
		{0, 5, 0},
		{5, 5, 0},
		{5, 0, 0},
	},
}

type ZMino struct {
	BaseMino
}

func NewZMino(x, y int, t string) Mino {
	return &ZMino{
		BaseMino{x, y, t, zminoTypes},
	}
}

func (m *ZMino) LeftRotate(field *Field) {
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

func (m *ZMino) RightRotate(f *Field) {
}

func (m *ZMino) Copy() Mino {
	return &ZMino{
		BaseMino{m.X, m.Y, m.Type, m.Types},
	}
}
