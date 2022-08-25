package tetris

var jminoTypes = map[string][][]int{
	"type1": {
		{4, 0, 0},
		{4, 4, 4},
		{0, 0, 0},
	},
	"type2": {
		{0, 4, 4},
		{0, 4, 0},
		{0, 4, 0},
	},
	"type3": {
		{0, 0, 0},
		{4, 4, 4},
		{0, 0, 4},
	},
	"type4": {
		{0, 4, 0},
		{0, 4, 0},
		{4, 4, 0},
	},
}

type JMino struct {
	BaseMino
}

func NewJMino(x, y int, t string) Mino {
	return &JMino{
		BaseMino{x, y, t, jminoTypes},
	}
}

func (m *JMino) LeftRotate(field *Field) {
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

func (m *JMino) RightRotate(f *Field) {
	//TODO implement me
	panic("implement me")
}

func (m *JMino) Copy() Mino {
	return &JMino{
		BaseMino{m.X, m.Y, m.Type, m.Types},
	}
}
